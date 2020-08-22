package gir

import (
	"encoding/xml"

	"github.com/dave/jennifer/jen"
)

type Constructor struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Name        string   `xml:"name,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	CallableAttrs
}

func (c Constructor) GenFunc(class Class, classes []Class) *jen.Statement {
	return jen.Func().Id(class.Name + "New").Params().Op("*").Id(class.Name).BlockFunc(
		func(g *jen.Group) {
			g.Id("v").Op(":=").Qual("C", c.CIdentifier).Call()
			g.Id("obj").Op(":=").Qual("github.com/gotk3/gotk3/glib", "Take").Call(
				jen.Qual("unsafe", "Pointer").Call(jen.Id("v")),
			)
			g.Return(jen.Op("&").Id(class.GoType()).Values(
				resolveWrapValues(classes, class.ParentInstance().GoType())),
			)
		},
	)
}

func resolveWrapValues(classes []Class, currentType string) *jen.Statement {
	switch currentType {
	case "*glib.Object":
		return jen.Id("obj")
	case "":
		return nil
	}

	return jen.Id(currentType).Values(
		resolveWrapValues(classes, ParentField(classes, currentType)),
	)
}
