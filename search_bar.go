package handy

// #include <handy.h>
// #include <gtk/gtk.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type SearchBar struct {
	gtk.Bin
}

func (b *SearchBar) native() *C.HdySearchBar {
	return C.HDY_SEARCH_BAR(gwidget(b))
}

func SearchBarNew() *SearchBar {
	v := C.hdy_search_bar_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &SearchBar{gtk.Bin{container(obj)}}
}

func (b *SearchBar) ConnectEntry(entry *gtk.Entry) {
	C.hdy_search_bar_connect_entry(b.native(), (*C.GtkEntry)(nwidget(entry)))
}

func (b *SearchBar) GetSearchMode() bool {
	v := C.hdy_search_bar_get_search_mode(b.native())
	return v == C.TRUE
}

func (b *SearchBar) SetSearchMode(searchMode bool) {
	C.hdy_search_bar_set_search_mode(b.native(), cbool(searchMode))
}

func (b *SearchBar) GetShowCloseButton() bool {
	v := C.hdy_search_bar_get_show_close_button(b.native())
	return v == C.TRUE
}

func (b *SearchBar) SetShowCloseButton(visible bool) {
	C.hdy_search_bar_set_show_close_button(b.native(), cbool(visible))
}

func (b *SearchBar) HandleEvent(ev *gdk.Event) bool {
	v := C.hdy_search_bar_handle_event(b.native(), (*C.GdkEvent)(ev.GdkEvent))
	return v == C.TRUE
}
