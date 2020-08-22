package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Type struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 type"`
	Name           string   `xml:"name,attr"`
	CType          string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Introspectable *bool    `xml:"introspectable,attr"`
}

func (t Type) IsPtr() bool {
	return len(t.CType) > 0 && t.CType[len(t.CType)-1] == '*'
}

// CGoType returns the C type in CGo.
func (t Type) CGoType() string {
	return CGoType(t.CType)
}

func (t Type) GenCGoType() *jen.Statement {
	var stmt = new(jen.Statement)
	if t.IsPtr() {
		stmt.Op("*")
	}

	stmt.Qual("C", strings.TrimSuffix(t.CType, "*"))

	return stmt
}

func CGoType(ctype string) (gotype string) {
	var ptr = len(ctype) > 0 && ctype[len(ctype)-1] == '*'
	var typ = fmt.Sprintf("C.%s", strings.TrimSuffix(ctype, "*"))
	if ptr {
		return "*" + typ
	}
	return typ
}

// GoType returns the type in Go.
func (t Type) GoType() string {
	return t.Type().GoString()
}

// Type returns the generated Go type in Go code.
func (t Type) Type() *jen.Statement {
	return t.Map(false)
}

// Map maps the type to a Go type in Go code.
func (t Type) Map(isObj bool) *jen.Statement {
	switch t.Name {
	case "void", "none":
		return nil
	case "gboolean":
		return jen.Bool()
	case "gfloat":
		return jen.Float32()
	case "gdouble":
		return jen.Float64()
	case "gint":
		return jen.Int()
	case "guint":
		return jen.Uint()
	case "utf8":
		return jen.String()
	case "gpointer":
		// TODO: ignore field
		// TODO: aaaaaaaaaaaaaaaaaaaaaaa
		return jen.Qual("unsafe", "Pointer")

	case "GLib.DestroyNotify":
		return jen.Id("DestroyNotify")
	}

	if parts := strings.Split(t.Name, "."); len(parts) == 2 {
		var stmt = new(jen.Statement)
		if typeMapInterface(parts) && t.IsPtr() {
			stmt = jen.Op("*")
		}

		switch parts[0] {
		case "Gtk":
			return stmt.Qual("github.com/gotk3/gotk3/gtk", parts[1])
		case "Gdk":
			return stmt.Qual("github.com/gotk3/gotk3/gdk", parts[1])
		case "GObject", "Gio", "GLib":
			return stmt.Qual("github.com/gotk3/gotk3/glib", parts[1])
		case "Pango":
			return stmt.Qual("github.com/gotk3/gotk3/pango", parts[1])
		case "Cairo":
			return stmt.Qual("github.com/gotk3/gotk3/cairo", parts[1])
		}
	}

	if isObj {
		return jen.Qual("github.com/gotk3/gotk3/glib", "InitiallyUnowned")
	}

	return jen.Id(t.Name)
}

func (t Type) ZeroValue() *jen.Statement {
	if t.IsPtr() {
		return jen.Nil()
	}

	switch t.Name {
	case "gboolean":
		return jen.False()
	case "gfloat", "gdouble":
		return jen.Lit(0.0)
	case "gint", "guint", "gpointer":
		return jen.Lit(0)
	case "utf8":
		return jen.Lit("")
	case "GLib.DestroyNotify":
		return jen.Nil()
	}

	return t.Type().Values()
}

func typeMapInterface(parts []string) (ptr bool) {
	switch parts[0] {
	case "Gtk":
		switch parts[1] {
		case "Widget":
			parts[1] = "IWidget"
			return false
		}
	}

	return true
}

// GenCaster generates the type or function to be used to cast or convert C to
// Go types.
func (t Type) GenCaster(tmpVar, value *jen.Statement) *jen.Statement {
	var stmt = tmpVar.Clone().Op(":=")

	switch t.GoType() {
	case "bool":
		stmt.Id("gobool")
	case "string":
		stmt.Qual("C", "GoString")

	// Handle IWidget separately.
	case "gtk.IWidget":
		stmt = jen.List(tmpVar, jen.Err()).Op(":=").Id("castWidget").Call(value)
		stmt.Line()
		stmt.If(jen.Err().Op("!=").Nil()).Block(
			jen.Panic(
				jen.Lit(fmt.Sprintf("cast widget %s failed: ", t.CGoType())).
					Op("+").
					Err().Dot("Error").Call(),
			),
		)

		return stmt

	// Handle glib.Object separately.
	case "glib.Object":
		// TODO: see if this leaks.
		stmt.Qual("github.com/gotk3/gotk3/glib", "Take")

	default:
		if t.IsPtr() {
			stmt.Parens(t.Type())
		} else {
			stmt.Add(t.Type())
		}
	}

	stmt.Call(value)

	return stmt
}
