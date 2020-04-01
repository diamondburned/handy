package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PreferencesPage struct {
	gtk.ScrolledWindow
}

func (p *PreferencesPage) native() *C.HdyPreferencesPage {
	return C.HDY_PREFERENCES_PAGE(gwidget(p))
}

func PreferencesPageNew() *PreferencesPage {
	v := C.hdy_preferences_page_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &PreferencesPage{gtk.ScrolledWindow{gtk.Bin{container(obj)}}}
}

func (p *PreferencesPage) GetIconName() string {
	v := C.hdy_preferences_page_get_icon_name(p.native())
	return C.GoString(v)
}

func (p *PreferencesPage) SetIconName(iconName string) {
	var _iconName *C.gchar
	if iconName != "" {
		_iconName = C.CString(iconName)
		defer C.free(unsafe.Pointer(_iconName))
	}

	C.hdy_preferences_page_set_icon_name(p.native(), _iconName)
}

func (p *PreferencesPage) GetTitle() string {
	v := C.hdy_preferences_page_get_title(p.native())
	return C.GoString(v)
}

func (p *PreferencesPage) SetTitle(title string) {
	var _title *C.gchar
	if title != "" {
		_title = C.CString(title)
		defer C.free(unsafe.Pointer(_title))
	}

	C.hdy_preferences_page_set_title(p.native(), _title)
}
