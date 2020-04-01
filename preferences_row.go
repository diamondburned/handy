package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PreferencesRow struct {
	gtk.Box
}

func (p *PreferencesRow) native() *C.HdyPreferencesRow {
	return C.HDY_PREFERENCES_ROW(gwidget(p))
}

func PreferencesRowNew() *PreferencesRow {
	v := C.hdy_preferences_row_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &PreferencesRow{gtk.Box{container(obj)}}
}

func (p *PreferencesRow) GetTitle() string {
	v := C.hdy_preferences_row_get_title(p.native())
	return C.GoString(v)
}

func (p *PreferencesRow) SetTitle(title string) {
	cs := C.CString(title)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_preferences_row_set_title(p.native(), cs)
}

func (p *PreferencesRow) GetUseUnderline() bool {
	v := C.hdy_preferences_row_get_use_underline(p.native())
	return v == C.TRUE
}

func (p *PreferencesRow) SetUseUnderline(useUnderline bool) {
	C.hdy_preferences_row_set_use_underline(p.native(), cbool(useUnderline))
}
