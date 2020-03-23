package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type LeafletTransitionType int

const (
	LEAFLET_TRANSITION_TYPE_NONE  LeafletTransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_NONE
	LEAFLET_TRANSITION_TYPE_SLIDE LeafletTransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_SLIDE
	LEAFLET_TRANSITION_TYPE_OVER  LeafletTransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_OVER
	LEAFLET_TRANSITION_TYPE_UNDER LeafletTransitionType = C.HDY_LEAFLET_TRANSITION_TYPE_UNDER
)

// Leaflet is an adaptive container acting like a box or a stack.
//
// Description
//
// The HdyLeaflet widget can display its children like a GtkBox does or like a
// GtkStack does, adapting to size changes by switching between the two modes.
//
// When there is enough space the children are displayed side by side, otherwise
// only one is displayed. The threshold is dictated by the preferred minimum
// sizes of the children.
type Leaflet struct {
	gtk.Container
}

func (l *Leaflet) native() *C.HdyLeaflet {
	return C.HDY_LEAFLET(gwidget(l))
}

func LeafletNew() *Leaflet {
	c := C.hdy_leaflet_new()
	obj := glib.Take(unsafe.Pointer(c))
	return &Leaflet{container(obj)}
}

// GetFold gets whether self is folded.
func (l *Leaflet) GetFold() Fold {
	c := C.hdy_leaflet_get_fold(l.native())
	return Fold(c)
}

// GetVisibleChild gets the visible child widget.
func (l *Leaflet) GetVisibleChild() *gtk.Widget {
	c := C.hdy_leaflet_get_visible_child(l.native())
	obj := glib.Take(unsafe.Pointer(c))
	return &gtk.Widget{glib.InitiallyUnowned{obj}}
}

// SetVisibleChild makes visible_child visible using a transition determined by
// HdyLeaflet:transition-type and HdyLeaflet:child-transition-duration. The
// transition can be cancelled by the user, in which case visible child will
// change back to the previously visible child.
func (l *Leaflet) SetVisibleChild(visibleChild gtk.IWidget) {
	C.hdy_leaflet_set_visible_child(l.native(), cwidget(visibleChild))
}

// GetVisibleChildName gets the name of the currently visible child widget.
func (l *Leaflet) GetVisibleChildName() string {
	c := C.hdy_leaflet_get_visible_child_name(l.native())
	return C.GoString(c)
}

// SetVisibleChildName makes the child with the name name visible.
// See (*Leaflet).SetVisibleChild() for more details.
func (l *Leaflet) SetVisibleChildName(name string) {
	C.hdy_leaflet_set_visible_child_name(l.native(), C.CString(name))
}

// GetHomogeneous gets whether self is homogeneous for the given fold and
// orientation. See (*Leaflet).SetHomogeneous().
func (l *Leaflet) GetHomogeneous(fold Fold, orientation gtk.Orientation) bool {
	c := C.hdy_leaflet_get_homogeneous(l.native(), C.HdyFold(fold), C.GtkOrientation(orientation))
	return gobool(c)
}

// SetHomogeneous ets the HdyLeaflet to be homogeneous or not for the given fold
// and orientation. If it is homogeneous, the HdyLeaflet will request the same
// width or height for all its children depending on the orientation. If it
// isn't and it is folded, the leaflet may change width or height when a
// different child becomes visible.
func (l *Leaflet) SetHomogeneous(fold Fold, orientation gtk.Orientation, homogeneous bool) {
	C.hdy_leaflet_set_homogeneous(
		l.native(), C.HdyFold(fold), C.GtkOrientation(orientation), cbool(homogeneous))
}

// GetTransitionType gets the type of animation that will be used for
// transitions between modes and children in self.
func (l *Leaflet) GetTransitionType() LeafletTransitionType {
	c := C.hdy_leaflet_get_transition_type(l.native())
	return LeafletTransitionType(c)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between modes and children in self.
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the mode or child that is about to
// become current.
func (l *Leaflet) SetTransitionType(transition LeafletTransitionType) {
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
	return gobool(c)
}

// GetInterpolateSize returns wether the HdyLeaflet is set up to interpolate
// between the sizes of children on page switch.
func (l *Leaflet) GetInterpolateSize() bool {
	c := C.hdy_leaflet_get_interpolate_size(l.native())
	return gobool(c)
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
	return gobool(c)
}

// SetCanSwipeBack sets whether or not self allows switching to the previous
// child that has 'allow-visible' child property set to TRUE via a swipe
// gesture.
func (l *Leaflet) SetCanSwipeBack(canSwipeBack bool) {
	C.hdy_leaflet_set_can_swipe_back(l.native(), cbool(canSwipeBack))
}

// GetCanSwipeForward returns whether the HdyLeaflet allows swiping to the next
// child.
func (l *Leaflet) GetCanSwipeForward() bool {
	c := C.hdy_leaflet_get_can_swipe_forward(l.native())
	return gobool(c)
}

// SetCanSwipeForward sets whether or not self allows switching to the next
// child that has 'allow-visible' child property set to TRUE via a swipe
// gesture.
func (l *Leaflet) SetCanSwipeForward(canSwipeForward bool) {
	C.hdy_leaflet_set_can_swipe_forward(l.native(), cbool(canSwipeForward))
}
