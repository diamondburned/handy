package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PaginatorIndicatorStyle int

const (
	PAGINATOR_INDICATOR_STYLE_NONE  PaginatorIndicatorStyle = C.HDY_PAGINATOR_INDICATOR_STYLE_NONE
	PAGINATOR_INDICATOR_STYLE_DOTS  PaginatorIndicatorStyle = C.HDY_PAGINATOR_INDICATOR_STYLE_DOTS
	PAGINATOR_INDICATOR_STYLE_LINES PaginatorIndicatorStyle = C.HDY_PAGINATOR_INDICATOR_STYLE_LINES
)

type Paginator struct {
	gtk.EventBox
}

func (p *Paginator) native() *C.HdyPaginator {
	return C.HDY_PAGINATOR(gwidget(p))
}

func PaginatorNew() *Paginator {
	v := C.hdy_paginator_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &Paginator{gtk.EventBox{gtk.Bin{container(obj)}}}
}

func (p *Paginator) Prepend(child gtk.IWidget) {
	C.hdy_paginator_prepend(p.native(), cwidget(child))
}

func (p *Paginator) Insert(child gtk.IWidget, position int) {
	C.hdy_paginator_insert(p.native(), cwidget(child), C.gint(position))
}

func (p *Paginator) Reorder(child gtk.IWidget, position int) {
	C.hdy_paginator_reorder(p.native(), cwidget(child), C.gint(position))
}

func (p *Paginator) ScrollTo(widget gtk.IWidget) {
	C.hdy_paginator_scroll_to(p.native(), cwidget(widget))
}

func (p *Paginator) ScrollToFull(widget gtk.IWidget, duration int64) {
	C.hdy_paginator_scroll_to_full(p.native(), cwidget(widget), C.gint64(duration))
}

func (p *Paginator) GetNPages() uint {
	v := C.hdy_paginator_get_n_pages(p.native())
	return uint(v)
}

// GetPosition gets current scroll position in self. It's unitless, 1 matches 1
// page.
func (p *Paginator) GetPosition() float64 {
	v := C.hdy_paginator_get_position(p.native())
	return float64(v)
}

func (p *Paginator) GetInteractive() bool {
	v := C.hdy_paginator_get_interactive(p.native())
	return gobool(v)
}

func (p *Paginator) SetInteractive(interactive bool) {
	C.hdy_paginator_set_interactive(p.native(), cbool(interactive))
}

func (p *Paginator) GetIndicatorStyle() PaginatorIndicatorStyle {
	v := C.hdy_paginator_get_indicator_style(p.native())
	return PaginatorIndicatorStyle(v)
}

func (p *Paginator) SetIndicatorStyle(v PaginatorIndicatorStyle) {
	C.hdy_paginator_set_indicator_style(p.native(), C.HdyPaginatorIndicatorStyle(v))
}

func (p *Paginator) GetIndicatorSpacing() uint {
	v := C.hdy_paginator_get_indicator_spacing(p.native())
	return uint(v)
}

func (p *Paginator) SetIndicatorSpacing(spacing uint) {
	C.hdy_paginator_set_indicator_spacing(p.native(), C.guint(spacing))
}

func (p *Paginator) GetCenterContent() bool {
	v := C.hdy_paginator_get_center_content(p.native())
	return gobool(v)
}

func (p *Paginator) SetCenterContent(centerContent bool) {
	C.hdy_paginator_set_center_content(p.native(), cbool(centerContent))
}

func (p *Paginator) GetSpacing() uint {
	v := C.hdy_paginator_get_spacing(p.native())
	return uint(v)
}

func (p *Paginator) SetSpacing(spacing uint) {
	C.hdy_paginator_set_spacing(p.native(), C.guint(spacing))
}

func (p *Paginator) GetAnimationDuration() uint {
	v := C.hdy_paginator_get_animation_duration(p.native())
	return uint(v)
}

func (p *Paginator) SetAnimationDuration(duration uint) {
	C.hdy_paginator_set_animation_duration(p.native(), C.guint(duration))
}

func (p *Paginator) GetAllowMouseDrag() bool {
	v := C.hdy_paginator_get_allow_mouse_drag(p.native())
	return gobool(v)
}

func (p *Paginator) SetAllowMouseDrag(allow bool) {
	C.hdy_paginator_set_allow_mouse_drag(p.native(), cbool(allow))
}
