package gir

import (
	"encoding/xml"

	"github.com/dave/jennifer/jen"
)

type Class struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	Parent string `xml:"parent,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Callbacks    []Callback    `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
}

func (c Class) GenType() *jen.Statement {
	return jen.Type().Id(c.Name).Struct(jen.Id(c.ParentInstance().GoType()))
}

func (c Class) GenNative() *jen.Statement {
	i := firstChar(c.Name)
	p := jen.Id(i).Op("*").Id(c.Name)

	f := jen.Func().Params(p).Id("native").Params().Id("*" + c.CGoType()).Block(
		jen.Return(
			jen.Parens(jen.Op("*").Qual("C", c.CType)).Call(
				jen.Id("gwidget").Call(jen.Id(i)),
			),
		),
	)
	f.Line()

	return f
}

func (c Class) GenConstructors(classes []Class) *jen.Statement {
	var stmt = make(jen.Statement, 0, len(c.Constructors)*2)
	for _, ctor := range c.Constructors {
		stmt.Add(ctor.GenFunc(c, classes))
		stmt.Line()
	}

	return &stmt
}

func (c Class) GenMethods() *jen.Statement {
	var stmt = make(jen.Statement, 0, len(c.Methods)*3)
	for _, method := range c.Methods {
		stmt.Add(method.GenFunc(c))
		stmt.Line()
		stmt.Line()
	}

	return &stmt
}

func (c Class) ParentInstance() *Field {
	for _, field := range c.Fields {
		if field.Name == "parent_instance" {
			return &field
		}
	}

	return nil
}

func (c Class) CGoType() string {
	return CGoType(c.CType)
}

func (c Class) GoType() string {
	return snakeToGo(true, c.Name)
}
