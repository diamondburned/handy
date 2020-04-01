package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PreferencesWindow struct {
	gtk.Window
}

func (w *PreferencesWindow) native() *C.HdyPreferencesWindow {
	return C.HDY_PREFERENCES_WINDOW(gwidget(w))
}

func PreferencesWindowNew() *PreferencesWindow {
	v := C.hdy_preferences_window_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &PreferencesWindow{gtk.Window{gtk.Bin{container(obj)}}}
}

// func (l *Leaflet) GetSearchEnabled() bool {
// 	v := C.hdy_preferences_window_get_search_enabled(l.native())
// 	return v == C.TRUE
// }

// func (l *Leaflet) SetSearchEnabled(searchEnabled bool) {
// 	C.hdy_preferences_window_set_search_enabled(l.native, cbool(searchEnabled))
// }
