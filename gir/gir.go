package gir

import (
	"encoding/xml"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"
)

func firstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

var (
	snakeRegex = regexp.MustCompile(`_\w`)
	snakeRepl  = strings.NewReplacer(
		"Xalign", "XAlign",
		"Yalign", "YAlign",
		"Id", "ID",
	)
)

func snakeToGo(pascal bool, snakeString string) string {
	if pascal {
		snakeString = "_" + snakeString
	}

	snakeString = snakeRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			return string(unicode.ToUpper(rune(orig[1])))
		},
	)

	return snakeRepl.Replace(snakeString)
}

func NewGotk3Generator(name string) *jen.File {
	f := jen.NewFile(name)
	f.ImportName("github.com/gotk3/gotk3/gtk", "gtk")
	f.ImportName("github.com/gotk3/gotk3/gdk", "gdk")
	f.ImportName("github.com/gotk3/gotk3/glib", "glib")
	f.ImportName("github.com/gotk3/gotk3/pango", "pango")
	f.ImportName("github.com/gotk3/gotk3/cairo", "cairo")
	f.ImportName("github.com/diamondburned/handy/handy/callback", "callback")
	f.CgoPreamble("#include <handy.h>")
	f.CgoPreamble("extern void callbackDelete(gpointer)")

	f.Comment("//export callbackDelete")
	f.Func().Id("callbackDelete").Params(jen.Id("ptr").Uintptr()).Block(
		jen.Qual("github.com/diamondburned/handy/handy/callback", "Delete").Call(
			jen.Id("ptr"),
		),
	)

	return f
}

type Annotation struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type CInclude struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Name    string   `xml:"name,attr"`
}

type Include struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Name    string   `xml:"name,attr"`
	Version *string  `xml:"version,attr"`
}

type Package struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Name    string   `xml:"name,attr"`
}

type Function struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Name    string   `xml:"name,attr"`
	CallableAttrs
}

type Enum struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Name    string   `xml:"name,attr"`
	Version *string  `xml:"version,attr"`

	GLibTypeName *string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType  *string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Members []Member `xml:"http://www.gtk.org/introspection/core/1.0 member"`
}

type Member struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Name        string   `xml:"name,attr"`
	Value       string   `xml:"value,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifer,attr"`
}

type TransferOwnership struct {
	TransferOwnership *string `xml:"transfer-ownership,attr"`
}
