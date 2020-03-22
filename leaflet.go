package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type TransitionType int

const (
	LEAFLET_TRANSITION_TYPE_NONE  TransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_NONE
	LEAFLET_TRANSITION_TYPE_SLIDE TransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_SLIDE
	LEAFLET_TRANSITION_TYPE_OVER  TransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_OVER
	LEAFLET_TRANSITION_TYPE_UNDER TransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_UNDER
)

type Leaflet struct {
	gtk.Container
}

func (l *Leaflet) native() *C.HdyLeaflet {
	if l == nil || l.GObject == nil {
		return nil
	}
	return C.HDY_LEAFLET(gpointer(unsafe.Pointer(l.GObject)))
}

func LeafletNew() *Leaflet {
	c := C.hdy_leaflet_new()
	obj := glib.Take(unsafe.Pointer(c))
	return &Leaflet{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}
}

// func (l *Leaflet) GetFolded() bool {
// 	c := C.hdy_leaflet_get_folded(l.native())
// 	return c == C.TRUE
// }

func (l *Leaflet) GetVisibleChild() *gtk.Widget {
	c := C.hdy_leaflet_get_visible_child(l.native())
	obj := glib.Take(unsafe.Pointer(c))
	return &gtk.Widget{glib.InitiallyUnowned{obj}}
}

func (l *Leaflet) SetVisibleChild(visibleChild gtk.IWidget) {
	C.hdy_leaflet_set_visible_child(l.native(), cwidget(visibleChild))
}

func (l *Leaflet) GetVisibleChildName() string {
	c := C.hdy_leaflet_get_visible_child_name(l.native())
	return C.GoString(c)
}

func (l *Leaflet) SetVisibleChildName(name string) {
	C.hdy_leaflet_set_visible_child_name(l.native(), C.CString(name))
}

func (l *Leaflet) GetHomogeneous(fold Fold, orientation gtk.Orientation) bool {
	c := C.hdy_leaflet_get_homogeneous(l.native(), C.HdyFold(fold), C.GtkOrientation(orientation))
	return c == C.TRUE
}

func (l *Leaflet) SetHomogeneous(fold Fold, orientation gtk.Orientation, homogeneous bool) {
	C.hdy_leaflet_set_homogeneous(
		l.native(), C.HdyFold(fold), C.GtkOrientation(orientation), cbool(homogeneous))
}

func (l *Leaflet) GetTransitionType() TransitionType {
	c := C.hdy_leaflet_get_transition_type(l.native())
	return TransitionType(c)
}

func (l *Leaflet) SetTransitionType(transition TransitionType) {
	C.hdy_leaflet_set_transition_type(l.native(), C.HdyLeafletTransitionType(transition))
}

// GetModeTransitionDuration returns the amount of time (in milliseconds) that
// transitions between modes in self will take.
func (l *Leaflet) GetModeTransitionDuration() uint {
	c := C.hdy_leaflet_get_mode_transition_duration(l.native())
	return uint(c)
}

// SetModeTransitionDuration sets the duration that transitions between modes
// in self will take.
func (l *Leaflet) SetModeTransitionDuration(duration uint) {
	C.hdy_leaflet_set_mode_transition_duration(l.native(), C.guint(duration))
}

// GetChildTransitionDuration returns the amount of time (in milliseconds) that
// transitions between children in self will take.
func (l *Leaflet) GetChildTransitionDuration() uint {
	c := C.hdy_leaflet_get_child_transition_duration(l.native())
	return uint(c)
}

// SetChildTransitionDuration sets the duration that transitions between
// children in self will take.
func (l *Leaflet) SetChildTransitionDuration(duration uint) {
	C.hdy_leaflet_set_child_transition_duration(l.native(), C.guint(duration))
}

// GetChildTransitionRunning returns whether self is currently in a transition
// from one page to another.
func (l *Leaflet) GetChildTransitionRunning() bool {
	c := C.hdy_leaflet_get_child_transition_running(l.native())
	return c == C.TRUE
}

// GetInterpolateSize returns wether the HdyLeaflet is set up to interpolate
// between the sizes of children on page switch.
func (l *Leaflet) GetInterpolateSize() bool {
	c := C.hdy_leaflet_get_interpolate_size(l.native())
	return c == C.TRUE
}

// SetInterpolateSize sets whether or not self will interpolate its size when
// changing the visible child. If the “interpolate-size” property is set to
// TRUE, stack will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set transition
// duration.
func (l *Leaflet) SetInterpolateSize(interpolateSize bool) {
	C.hdy_leaflet_set_interpolate_size(l.native(), cbool(interpolateSize))
}

// GetCanSwipeBack returns whether the HdyLeaflet allows swiping to the previous
// child.
func (l *Leaflet) GetCanSwipeBack() bool {
	c := C.hdy_leaflet_get_can_swipe_back(l.native())
	return c == C.TRUE
}

// SetCanSwipeBack sets whether or not self allows switching to the previous
// child that has 'allow-visible' child property set to TRUE via a swipe
// gesture.
func (l *Leaflet) SetCanSwipeBack(canSwipeBack bool) {
	C.hdy_leaflet_set_can_swipe_back(l.native(), cbool(canSwipeBack))
}
