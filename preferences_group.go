package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PreferencesGroup struct {
	gtk.Box
}

func (p *PreferencesGroup) native() *C.HdyPreferencesGroup {
	return C.HDY_PREFERENCES_GROUP(gwidget(p))
}

func PreferencesGroupNew() *PreferencesGroup {
	v := C.hdy_preferences_group_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &PreferencesGroup{gtk.Box{container(obj)}}
}

func (p *PreferencesGroup) GetTitle() string {
	v := C.hdy_preferences_group_get_title(p.native())
	return C.GoString(v)
}

func (p *PreferencesGroup) SetTitle(title string) {
	cs := C.CString(title)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_preferences_group_set_title(p.native(), cs)
}

func (p *PreferencesGroup) GetDescription() string {
	v := C.hdy_preferences_group_get_description(p.native())
	return C.GoString(v)
}

func (p *PreferencesGroup) SetDescription(description string) {
	cs := C.CString(description)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_preferences_group_set_description(p.native(), cs)
}
