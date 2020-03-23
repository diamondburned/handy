package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// TitleBar is a simple title bar container.
//
// Description
//
// HdyTitleBar is meant to be used as the top-level widget of your window's
// title bar. It will be drawn with the same style as a GtkHeaderBar but it
// won't force a widget layout on you: you can put whatever widget you want in
// it, including a GtkHeaderBar.
// HdyTitleBar becomes really useful when you want to animate header bars, like
// an adaptive application using HdyLeaflet would do.
//
// https://developer.puri.sm/projects/libhandy/unstable/HdyTitleBar.html
type TitleBar struct {
	gtk.Bin
}

func (t *TitleBar) native() *C.HdyTitleBar {
	return C.HDY_TITLE_BAR(gwidget(t))
}

// TitleBarNew creates a new HdyTitleBar.
func TitleBarNew() *TitleBar {
	v := C.hdy_title_bar_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &TitleBar{gtk.Bin{container(obj)}}
}

// GetSelectionMode returns wether whether self is in selection mode.
func (t *TitleBar) GetSelectionMode() bool {
	v := C.hdy_title_bar_get_selection_mode(t.native())
	return gobool(v)
}

// SetSelectionMode sets whether self is in selection mode.
func (t *TitleBar) SetSelectionMode(selectionMode bool) {
	C.hdy_title_bar_set_selection_mode(t.native(), cbool(selectionMode))
}
