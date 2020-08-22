package gir

import (
	"encoding/xml"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Namespace struct {
	XMLName            xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
	Version            string   `xml:"version,attr"`
	SharedLibrary      string   `xml:"shared-library,attr"`
	IdentifierPrefixes string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	SymbolPrefixes     string   `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`

	Classes     []Class      `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Enums       []Enum       `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Functions   []Function   `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Callbacks   []Callback   `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
}

func (n Namespace) GenerateToFile(f *jen.File) {
	f.CgoPreamble(n.GenCallbackPreamble())
	f.Add(n.GenCallbacks())
	f.Add(n.GenClasses())
}

func (n Namespace) GenCallbackPreamble() string {
	var preambles = make([]string, 0, len(n.Callbacks))
	for _, callback := range n.Callbacks {
		preambles = append(preambles, callback.GenExternC())
	}

	return strings.Join(preambles, "\n")
}

func (n Namespace) GenCallbacks() *jen.Statement {
	var f = new(jen.Statement)

	for _, callback := range n.Callbacks {
		f.Add(callback.GenGoType())
		f.Line()
		f.Add(callback.GenGlobalGoFunction())
		f.Line()
	}

	return f
}

func (n Namespace) GenClasses() *jen.Statement {
	var f = new(jen.Statement)

	for _, class := range n.Classes {
		if class.ParentInstance() == nil {
			continue
		}

		f.Add(class.GenType())
		f.Line()

		f.Add(class.GenConstructors(n.Classes))
		f.Line()

		f.Add(class.GenNative())
		f.Line()

		f.Add(class.GenMethods())
		f.Line()
	}

	return f
}

func (n Namespace) ParentField(gotype string) string {
	return ParentField(n.Classes, gotype)
}

func ParentField(classes []Class, gotype string) string {
	switch gotype {
	case "gtk.ApplicationWindow": // TODO: handle interfaces inside here
		return "gtk.Window"
	case "gtk.Window":
		return "gtk.Bin"
	case "gtk.ListBoxRow":
		return "gtk.Bin"
	case "gtk.Bin":
		return "gtk.Container"
	case "gtk.Container":
		return "gtk.Widget"
	case "gtk.Widget":
		return "glib.InitiallyUnowned"
	case "glib.InitiallyUnowned":
		return "*glib.Object"
	}

	for _, class := range classes {
		if class.GoType() == gotype {
			return class.ParentInstance().GoType()
		}
	}

	log.Panicln("Unknown type:", gotype)
	return ""
}
