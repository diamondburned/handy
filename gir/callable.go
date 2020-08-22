package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type CallableAttrs struct {
	Parameters  *Parameters
	ReturnValue *ReturnValue `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
}

func (a CallableAttrs) HasInstanceParameter() bool {
	return a.Parameters != nil && a.Parameters.HasInstanceParameter()
}

type Parameters struct {
	XMLName           xml.Name           `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	InstanceParameter *InstanceParameter `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	Parameters        []Parameter        `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
}

// HasInstanceParameter returns true if p and InstanceParameter are not nil.
func (p *Parameters) HasInstanceParameter() bool {
	return p != nil && p.InstanceParameter != nil && !p.InstanceParameter.IsIgnored()
}

// SearchUserData searches for the UserData parameter. It returns nil if p is
// nil or if userData is not in the list of parameters.
func (p *Parameters) SearchUserData() *Parameter {
	if p == nil {
		return nil
	}

	for _, param := range p.Parameters {
		if param.IsUserData() {
			return &param
		}
	}

	return nil
}

type InstanceParameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	ParameterAttrs
}

type Parameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
	ParameterAttrs
}

type ParameterAttrs struct {
	Name string `xml:"name,attr"`
	TransferOwnership
	Type Type
}

func (p ParameterAttrs) GoName() string {
	return snakeToGo(false, p.Name)
}

var ignoredParams = []func(ParameterAttrs) bool{
	ParameterAttrs.IsUserData,
	ParameterAttrs.IsUserDataFreeFunc,
}

func (p ParameterAttrs) IsIgnored() bool {
	for _, isIgnored := range ignoredParams {
		if isIgnored(p) {
			return true
		}
	}

	return false
}

func (p ParameterAttrs) IsUserData() bool {
	return p.Name == "user_data" && p.Type.Name == "gpointer"
}

func (p ParameterAttrs) IsUserDataFreeFunc() bool {
	return p.Name == "user_data_free_func" && p.Type.Name == "GLib.DestroyNotify"
}

// GenValueCall generates a value conversion call from the given names. If
// argName is nil, then the returned value will be a zero-value.
func (p ParameterAttrs) GenValueCall(argName, valueName *jen.Statement) *jen.Statement {
	var stmt = jen.Add(valueName).Op(":=")

	// Filter out ignored parameters.
	if p.IsIgnored() || argName == nil {
		return nil
	}

	// TODO: account for enums

	switch p.Type.GoType() {
	case "bool":
		stmt.Id("cbool").Call(argName)
	case "float32":
		stmt.Qual("C", "gfloat").Call(argName)
	case "float64":
		stmt.Qual("C", "gdouble").Call(argName)
	case "int":
		stmt.Qual("C", "gint").Call(argName)
	case "uint":
		stmt.Qual("C", "guint").Call(argName)
	case "unsafe.Pointer":
		stmt.Qual("C", "gpointer").Call(jen.Uintptr().Call(argName))
	case "string":
		stmt.Qual("C", "CString").Call(argName)
		stmt.Line()
		stmt.Defer().Qual("C", "free").Call(jen.Qual("unsafe", "Pointer").Call(valueName))
	case "gtk.IWidget":
		stmt.Id("cwidget").Call(argName)
	default:
		argName := argName.Clone()
		stmt.Parens(p.Type.GenCGoType()).Call(
			jen.Qual("unsafe", "Pointer").Call(argName.Op(".").Id("Native").Call()),
		)
	}

	return stmt
}

// TODO:
// 1. Create C function
// 2. Make C function access the map

// func (p ParameterAttrs)

type ReturnValue struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	TransferOwnership
	Type Type
}

// IsVoid returns true if the type name is "none" or if *ReturnValue is nil.
func (r *ReturnValue) IsVoid() bool {
	if r == nil {
		return true
	}

	return r.Type.Name == "none"
}

// GenReturn generates a statement with the return token.
func (r *ReturnValue) GenReturnFunc(call *jen.Statement) *jen.Statement {
	if r.IsVoid() {
		return call
	}

	v := jen.Id("r")
	return jen.Add(r.Type.GenCaster(v, call)).Line().Return(v)
}

type Callback struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Name    string   `xml:"name,attr"`
	CType   string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CallableAttrs
}

func (c Callback) GenGoType() *jen.Statement {
	s := jen.Type().Id(c.Name).Func()

	s.ParamsFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		for _, param := range c.Parameters.Parameters {
			if param.IsIgnored() {
				continue
			}

			g.Add(param.Type.Type())
		}
	})

	if !c.ReturnValue.IsVoid() {
		s.Add(c.ReturnValue.Type.Type())
	}

	return s
}

// GenGlobalGoFunction generates a Go function with the export comment. This
// function is used to be called from C. It triggers the callback inside the
// map.
func (c Callback) GenGlobalGoFunction() *jen.Statement {
	s := jen.Comment("//export callback" + c.Name)
	s.Line()
	s.Func().Id("callback" + c.Name)

	s.ParamsFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		for _, param := range c.Parameters.Parameters {
			g.Add(jen.Id(param.GoName()), param.Type.GenCGoType())
		}
	})

	if !c.ReturnValue.IsVoid() {
		s.Add(c.ReturnValue.Type.GenCGoType())
	}

	var goargs map[string]*jen.Statement
	if c.Parameters != nil {
		goargs = make(map[string]*jen.Statement, len(c.Parameters.Parameters))
	}

	s.BlockFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		// Get the callback closure from the global map, if there's a userData
		// argument.
		var userData = c.Parameters.SearchUserData()
		if userData == nil {
			return
		}

		g.Id("fn").Op(":=").Qual("github.com/diamondburned/handy/handy/callback", "Get").Call(
			jen.Uintptr().Call(jen.Id(userData.GoName())),
		)

		// TODO: is this panic worthy?
		g.If(jen.Id("fn").Op("==").Nil()).Block(
			jen.Panic(jen.Lit(fmt.Sprintf("callback for %s not found", c.Name))),
		)

		g.Line()

		// Convert C arguments to Go variables.
		for i, param := range c.Parameters.Parameters {
			if param.IsIgnored() {
				continue
			}

			v := jen.Id(fmt.Sprintf("arg%d", i))
			goargs[param.Name] = v

			g.Add(param.Type.GenCaster(v, jen.Id(param.GoName())))
		}

		g.Line()

		g.Id("fn").Op(".").Parens(jen.Id(c.Name)).CallFunc(func(g *jen.Group) {
			for _, param := range c.Parameters.Parameters {
				if goarg, ok := goargs[param.Name]; ok {
					g.Add(goarg)
				}
			}
		})
	})

	return s
}

func (c Callback) GenExternC() string {
	s := strings.Builder{}
	s.WriteString("extern ")
	s.WriteString(c.ReturnValue.Type.CType)
	s.WriteString(" ")
	s.WriteString("callback_")
	s.WriteString(c.Name)

	s.WriteString("(")

	if c.Parameters != nil {
		var params = make([]string, len(c.Parameters.Parameters))
		for i, param := range c.Parameters.Parameters {
			params[i] = param.Type.CType
		}
		s.WriteString(strings.Join(params, ", "))
	}

	s.WriteString(")")

	return s.String()
}

// TODO: gen assign fn

// func (c Callback) GenAssignCall() *jen.Statement {}
