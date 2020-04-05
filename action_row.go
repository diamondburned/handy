package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// ActionRow is a gtk.ListBoxRow used to present actions.
//
// Description
//
// The HdyActionRow widget can have a title, a subtitle and an icon. The row can
// receive action widgets at its end, prefix widgets at its start or widgets
// below it.
//
// Note that action widgets are packed starting from the end.
//
// It is convenient to present a list of preferences and their related actions.
type ActionRow struct {
	PreferencesRow
}

func (r *ActionRow) native() *C.HdyActionRow {
	return C.HDY_ACTION_ROW(gwidget(r))
}

func ActionRowNew() *ActionRow {
	v := C.hdy_action_row_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &ActionRow{PreferencesRow{gtk.ListBoxRow{gtk.Bin{container(obj)}}}}
}

func (r *ActionRow) GetTitle() string {
	v := C.hdy_action_row_get_title(r.native())
	return C.GoString(v)
}

func (r *ActionRow) SetTitle(title string) {
	cs := C.CString(title)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_action_row_set_title(r.native(), cs)
}

func (r *ActionRow) GetSubtitle() string {
	v := C.hdy_action_row_get_subtitle(r.native())
	return C.GoString(v)
}

func (r *ActionRow) SetSubtitle(subtitle string) {
	cs := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_action_row_set_subtitle(r.native(), cs)
}

func (r *ActionRow) GetIconName() string {
	v := C.hdy_action_row_get_icon_name(r.native())
	return C.GoString(v)
}

func (r *ActionRow) SetIconName(iconName string) {
	cs := C.CString(iconName)
	defer C.free(unsafe.Pointer(cs))
	C.hdy_action_row_set_icon_name(r.native(), cs)
}

func (r *ActionRow) GetActivatableWidget() *gtk.Widget {
	v := C.hdy_action_row_get_activatable_widget(r.native())
	obj := glib.Take(unsafe.Pointer(v))
	return widget(obj)
}

func (r *ActionRow) SetActivatableWidget(w gtk.IWidget) {
	C.hdy_action_row_set_activatable_widget(r.native(), cwidget(w))
}

func (r *ActionRow) GetUseUnderline() bool {
	v := C.hdy_action_row_get_use_underline(r.native())
	return v == C.TRUE
}

func (r *ActionRow) SetUseUnderline(useUnderline bool) {
	C.hdy_action_row_set_use_underline(r.native(), cbool(useUnderline))
}

func (r *ActionRow) AddAction(w gtk.IWidget) {
	C.hdy_action_row_add_action(r.native(), cwidget(w))
}

func (r *ActionRow) AddPrefix(w gtk.IWidget) {
	C.hdy_action_row_add_prefix(r.native(), cwidget(w))
}

func (r *ActionRow) Activate() {
	C.hdy_action_row_activate(r.native())
}
