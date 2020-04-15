package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type ExpanderRow struct {
	ActionRow
}

func (e *ExpanderRow) native() *C.HdyExpanderRow {
	return C.HDY_EXPANDER_ROW(gwidget(e))
}

func ExpanderRowNew() *ExpanderRow {
	v := C.hdy_expander_row_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &ExpanderRow{ActionRow{PreferencesRow{gtk.ListBoxRow{gtk.Bin{container(obj)}}}}}
}

// GetTitle sets the "title" property. As hdy_expander_row_set_title isn't
// available until libhandy v1.0, this is what we'll be using.
func (e *ExpanderRow) GetTitle() string {
	v, err := e.GetProperty("title")
	if err != nil || v == nil {
		return ""
	}
	return v.(string)
}

func (e *ExpanderRow) SetTitle(title string) {
	e.SetProperty("title", title)
}

// func (e *ExpanderRow) GetSubtitle() string {
// 	v := C.hdy_expander_row_get_subtitle(e.native())
// 	return C.GoString(v)
// }

// func (e *ExpanderRow) SetSubtitle(subtitle string) {
// 	cs := C.CString(subtitle)
// 	defer C.free(unsafe.Pointer(cs))
// 	C.hdy_expander_row_set_subtitle(e.native(), cs)
// }

// func (e *ExpanderRow) GetUseUnderline() bool {
// 	v := C.hdy_expander_row_get_use_underline(e.native())
// 	return gobool(v)
// }

// func (e *ExpanderRow) SetUseUnderline(useUnderline bool) {
// 	C.hdy_expander_row_set_use_underline(e.native(), cbool(useUnderline))
// }

// func (e *ExpanderRow) GetIconName() string {
// 	v := C.hdy_expander_row_get_icon_name(e.native())
// 	return C.GoString(v)
// }

// func (e *ExpanderRow) SetIconName(iconName string) {
// 	cs := C.CString(iconName)
// 	defer C.free(unsafe.Pointer(cs))
// 	C.hdy_expander_row_set_icon_name(e.native(), cs)
// }

func (e *ExpanderRow) GetExpanded() bool {
	v := C.hdy_expander_row_get_expanded(e.native())
	return gobool(v)
}

func (e *ExpanderRow) SetExpanded(expanded bool) {
	C.hdy_expander_row_set_expanded(e.native(), cbool(expanded))
}

func (e *ExpanderRow) GetEnableExpansion() bool {
	v := C.hdy_expander_row_get_enable_expansion(e.native())
	return gobool(v)
}

func (e *ExpanderRow) SetEnableExpansion(v bool) {
	C.hdy_expander_row_set_enable_expansion(e.native(), cbool(v))
}

func (e *ExpanderRow) GetShowEnableSwitch() bool {
	v := C.hdy_expander_row_get_show_enable_switch(e.native())
	return gobool(v)
}

func (e *ExpanderRow) SetShowEnableSwitch(v bool) {
	C.hdy_expander_row_set_show_enable_switch(e.native(), cbool(v))
}
