package gir

import (
	"encoding/xml"
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Method struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Name        string   `xml:"name,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	CallableAttrs
}

// TODO
// the method generation should use the full parameters, but only generate v%d for some

// // Parameters returns the list of filtered out parameters.
// func (m Method) Parameters() []Parameter {
// 	var params []Parameter
// 	if p := m.CallableAttrs.Parameters; p != nil {
// 		params = make([]Parameter, 0, len(p.Parameters))

// 		for _, param := range p.Parameters {
// 			if param.IsIgnored() {
// 				continue
// 			}

// 			params = append(params, param)
// 		}
// 	}

// 	return params
// }

func (m Method) GenFunc(c Class) *jen.Statement {
	i := jen.Id(firstChar(c.Name))
	p := jen.Add(i).Op("*").Id(c.Name)

	stmt := jen.Func().Params(p)
	stmt.Id(m.GoType())

	var parm = []Parameter{}
	if m.Parameters != nil {
		parm = m.Parameters.Parameters
	}
	var args = make(map[string]*jen.Statement, len(parm))

	// Generate the parameters in the function signature.
	stmt.ParamsFunc(func(g *jen.Group) {
		for _, param := range parm {
			if param.IsIgnored() {
				continue
			}

			n := jen.Id(param.GoName())
			args[param.Name] = n

			g.Add(n, param.Type.Type())
		}
	})

	if m.ReturnValue != nil {
		stmt.Add(m.ReturnValue.Type.Type())
	}

	// List of arguments to call the C function. Not to be confused with the
	// above list of arguments to call the current Go function.
	var cargs = make(map[string]*jen.Statement, len(parm))

	// Generate the value type converters in the function body.
	stmt.BlockFunc(func(g *jen.Group) {
		for i, param := range parm {
			// Only create needed values.
			if arg, ok := args[param.Name]; ok {
				valueVar := jen.Id(fmt.Sprintf("v%d", i+1))
				cargs[param.Name] = valueVar

				g.Add(param.GenValueCall(arg, valueVar))
			}
		}

		if len(parm) > 1 {
			g.Line()
		}

		g.Add(m.ReturnValue.GenReturnFunc(
			jen.Qual("C", m.CIdentifier).ParamsFunc(func(g *jen.Group) {
				if m.HasInstanceParameter() {
					g.Add(jen.Add(i).Op(".").Id("native").Call())
				}

				for _, param := range parm {
					a, ok := cargs[param.Name]
					if ok {
						g.Add(a)
					} else {
						// Add as a constant to allow implicit type casting.
						g.Add(param.Type.ZeroValue())
					}
				}
			}),
		))
	})

	return stmt
}

func (m Method) GoType() string {
	return snakeToGo(true, m.Name)
}

func (m Method) Type() *jen.Statement {
	return jen.Id(m.GoType())
}
