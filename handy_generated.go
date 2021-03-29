package handy

import (
	"github.com/diamondburned/handy/internal/callback"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
	"unsafe"
)

// #cgo pkg-config: libhandy-1 gtk+-3.0 glib-2.0 gio-2.0 glib-2.0 gobject-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <handy.h>
// #include <gtk/gtk.h>
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer ptr);
// extern GdkPixbuf* callbackAvatarImageLoadFunc(gint v0, gpointer v1);
// extern gchar* callbackComboRowGetEnumValueNameFunc(HdyEnumValueObject* v0, gpointer v1);
// extern gchar* callbackComboRowGetNameFunc(gpointer v0, gpointer v1);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) {
	callback.Delete(uintptr(ptr))
}

// objector is used internally for other interfaces.
type objector interface {
	glib.IObject
	Connect(string, interface{}) glib.SignalHandle
	ConnectAfter(string, interface{}) glib.SignalHandle
	GetProperty(name string) (interface{}, error)
	SetProperty(name string, value interface{}) error
	Native() uintptr
}

// asserting objector interface
var _ objector = (*glib.Object)(nil)

// Caster is the interface that allows casting objects to widgets.
type Caster interface {
	objector
	Cast() (gtk.IWidget, error)
}

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{
		// Enums
		{glib.Type(C.hdy_centering_policy_get_type()), marshalCenteringPolicy},
		{glib.Type(C.hdy_deck_transition_type_get_type()), marshalDeckTransitionType},
		{glib.Type(C.hdy_flap_fold_policy_get_type()), marshalFlapFoldPolicy},
		{glib.Type(C.hdy_flap_transition_type_get_type()), marshalFlapTransitionType},
		{glib.Type(C.hdy_header_group_child_type_get_type()), marshalHeaderGroupChildType},
		{glib.Type(C.hdy_leaflet_transition_type_get_type()), marshalLeafletTransitionType},
		{glib.Type(C.hdy_navigation_direction_get_type()), marshalNavigationDirection},
		{glib.Type(C.hdy_squeezer_transition_type_get_type()), marshalSqueezerTransitionType},
		{glib.Type(C.hdy_view_switcher_policy_get_type()), marshalViewSwitcherPolicy},

		// Objects/Classes
		{glib.Type(C.hdy_action_row_get_type()), marshalActionRow},
		{glib.Type(C.hdy_application_window_get_type()), marshalApplicationWindow},
		{glib.Type(C.hdy_avatar_get_type()), marshalAvatar},
		{glib.Type(C.hdy_carousel_get_type()), marshalCarousel},
		{glib.Type(C.hdy_carousel_indicator_dots_get_type()), marshalCarouselIndicatorDots},
		{glib.Type(C.hdy_carousel_indicator_lines_get_type()), marshalCarouselIndicatorLines},
		{glib.Type(C.hdy_clamp_get_type()), marshalClamp},
		{glib.Type(C.hdy_combo_row_get_type()), marshalComboRow},
		{glib.Type(C.hdy_deck_get_type()), marshalDeck},
		{glib.Type(C.hdy_enum_value_object_get_type()), marshalEnumValueObject},
		{glib.Type(C.hdy_expander_row_get_type()), marshalExpanderRow},
		{glib.Type(C.hdy_flap_get_type()), marshalFlap},
		{glib.Type(C.hdy_header_bar_get_type()), marshalHeaderBar},
		{glib.Type(C.hdy_header_group_get_type()), marshalHeaderGroup},
		{glib.Type(C.hdy_header_group_child_get_type()), marshalHeaderGroupChild},
		{glib.Type(C.hdy_keypad_get_type()), marshalKeypad},
		{glib.Type(C.hdy_leaflet_get_type()), marshalLeaflet},
		{glib.Type(C.hdy_preferences_group_get_type()), marshalPreferencesGroup},
		{glib.Type(C.hdy_preferences_page_get_type()), marshalPreferencesPage},
		{glib.Type(C.hdy_preferences_row_get_type()), marshalPreferencesRow},
		{glib.Type(C.hdy_preferences_window_get_type()), marshalPreferencesWindow},
		{glib.Type(C.hdy_search_bar_get_type()), marshalSearchBar},
		{glib.Type(C.hdy_squeezer_get_type()), marshalSqueezer},
		{glib.Type(C.hdy_status_page_get_type()), marshalStatusPage},
		{glib.Type(C.hdy_swipe_group_get_type()), marshalSwipeGroup},
		{glib.Type(C.hdy_swipe_tracker_get_type()), marshalSwipeTracker},
		{glib.Type(C.hdy_tab_bar_get_type()), marshalTabBar},
		{glib.Type(C.hdy_tab_page_get_type()), marshalTabPage},
		{glib.Type(C.hdy_tab_view_get_type()), marshalTabView},
		{glib.Type(C.hdy_title_bar_get_type()), marshalTitleBar},
		{glib.Type(C.hdy_value_object_get_type()), marshalValueObject},
		{glib.Type(C.hdy_view_switcher_get_type()), marshalViewSwitcher},
		{glib.Type(C.hdy_view_switcher_bar_get_type()), marshalViewSwitcherBar},
		{glib.Type(C.hdy_view_switcher_title_get_type()), marshalViewSwitcherTitle},
		{glib.Type(C.hdy_window_get_type()), marshalWindow},
		{glib.Type(C.hdy_window_handle_get_type()), marshalWindowHandle},
	})
}

type CenteringPolicy int

func marshalCenteringPolicy(p uintptr) (interface{}, error) {
	return CenteringPolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// CenteringPolicyLoose keep the title centered when possible
	CenteringPolicyLoose CenteringPolicy = 0
	// CenteringPolicyStrict keep the title centered at all cost
	CenteringPolicyStrict CenteringPolicy = 1
)

// DeckTransitionType enumeration value describes the possible transitions
// between children in a Deck widget.
//
// New values may be added to this enumeration over time.
type DeckTransitionType int

func marshalDeckTransitionType(p uintptr) (interface{}, error) {
	return DeckTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// DeckTransitionTypeOver cover the old page or uncover the new page,
	// sliding from or towards the end according to orientation, text direction
	// and children order
	DeckTransitionTypeOver DeckTransitionType = 0
	// DeckTransitionTypeUnder uncover the new page or cover the old page,
	// sliding from or towards the start according to orientation, text
	// direction and children order
	DeckTransitionTypeUnder DeckTransitionType = 1
	// DeckTransitionTypeSlide slide from left, right, up or down according to
	// the orientation, text direction and the children order
	DeckTransitionTypeSlide DeckTransitionType = 2
)

// FlapFoldPolicy these enumeration values describe the possible folding
// behavior in a Flap widget.
type FlapFoldPolicy int

func marshalFlapFoldPolicy(p uintptr) (interface{}, error) {
	return FlapFoldPolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// FlapFoldPolicyNever disable folding, the flap cannot reach narrow
	// sizes.
	FlapFoldPolicyNever FlapFoldPolicy = 0
	// FlapFoldPolicyAlways keep the flap always folded.
	FlapFoldPolicyAlways FlapFoldPolicy = 1
	// FlapFoldPolicyAuto fold and unfold the flap based on available   space.
	FlapFoldPolicyAuto FlapFoldPolicy = 2
)

// FlapTransitionType these enumeration values describe the possible transitions
// between children in a Flap widget, as well as which areas can be swiped via
// Flap:swipe-to-open and Flap:swipe-to-close.
//
// New values may be added to this enum over time.
type FlapTransitionType int

func marshalFlapTransitionType(p uintptr) (interface{}, error) {
	return FlapTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// FlapTransitionTypeOver the flap slides over the content, which is
	// dimmed. When folded, only the flap can be swiped.
	FlapTransitionTypeOver FlapTransitionType = 0
	// FlapTransitionTypeUnder the content slides over the flap. Only the
	// content can be swiped.
	FlapTransitionTypeUnder FlapTransitionType = 1
	// FlapTransitionTypeSlide the flap slides offscreen when hidden,   neither
	// the flap nor content overlap each other. Both widgets can be   swiped.
	FlapTransitionTypeSlide FlapTransitionType = 2
)

// HeaderGroupChildType enumeration value describes the child types handled by
// HeaderGroup.
//
// New values may be added to this enumeration over time.
type HeaderGroupChildType int

func marshalHeaderGroupChildType(p uintptr) (interface{}, error) {
	return HeaderGroupChildType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// HeaderGroupChildTypeHeaderBar the child is a HeaderBar
	HeaderGroupChildTypeHeaderBar HeaderGroupChildType = 0
	// HeaderGroupChildTypeGtkHeaderBar the child is a HeaderBar
	HeaderGroupChildTypeGtkHeaderBar HeaderGroupChildType = 1
	// HeaderGroupChildTypeHeaderGroup the child is a HeaderGroup
	HeaderGroupChildTypeHeaderGroup HeaderGroupChildType = 2
)

// LeafletTransitionType enumeration value describes the possible transitions
// between modes and children in a Leaflet widget.
//
// New values may be added to this enumeration over time.
type LeafletTransitionType int

func marshalLeafletTransitionType(p uintptr) (interface{}, error) {
	return LeafletTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// LeafletTransitionTypeOver cover the old page or uncover the new page,
	// sliding from or towards the end according to orientation, text direction
	// and children order
	LeafletTransitionTypeOver LeafletTransitionType = 0
	// LeafletTransitionTypeUnder uncover the new page or cover the old page,
	// sliding from or towards the start according to orientation, text
	// direction and children order
	LeafletTransitionTypeUnder LeafletTransitionType = 1
	// LeafletTransitionTypeSlide slide from left, right, up or down according
	// to the orientation, text direction and the children order
	LeafletTransitionTypeSlide LeafletTransitionType = 2
)

// NavigationDirection represents direction of a swipe navigation gesture in
// Deck and Leaflet.
type NavigationDirection int

func marshalNavigationDirection(p uintptr) (interface{}, error) {
	return NavigationDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// NavigationDirectionBack corresponds to start or top, depending on
	// orientation and text direction
	NavigationDirectionBack NavigationDirection = 0
	// NavigationDirectionForward corresponds to end or bottom, depending on
	// orientation and text direction
	NavigationDirectionForward NavigationDirection = 1
)

// SqueezerTransitionType these enumeration values describe the possible
// transitions between children in a Squeezer widget.
type SqueezerTransitionType int

func marshalSqueezerTransitionType(p uintptr) (interface{}, error) {
	return SqueezerTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// SqueezerTransitionTypeNone no transition
	SqueezerTransitionTypeNone SqueezerTransitionType = 0
	// SqueezerTransitionTypeCrossfade a cross-fade
	SqueezerTransitionTypeCrossfade SqueezerTransitionType = 1
)

type ViewSwitcherPolicy int

func marshalViewSwitcherPolicy(p uintptr) (interface{}, error) {
	return ViewSwitcherPolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

const (
	// ViewSwitcherPolicyAuto automatically adapt to the best fitting mode
	ViewSwitcherPolicyAuto ViewSwitcherPolicy = 0
	// ViewSwitcherPolicyNarrow force the narrow mode
	ViewSwitcherPolicyNarrow ViewSwitcherPolicy = 1
	// ViewSwitcherPolicyWide force the wide mode
	ViewSwitcherPolicyWide ViewSwitcherPolicy = 2
)

type Swiper interface {
	Caster
	// EmitChildSwitched emits HdySwipeable::child-switched signal. This should be
	// called when the widget switches visible child widget.
	//
	// duration can be 0 if the child is switched without animation.
	EmitChildSwitched(index uint, duration int64)
	// GetCancelProgress gets the progress Swiper will snap back to after the
	// gesture is canceled.
	GetCancelProgress() float64
	// GetDistance gets the swipe distance of Swiper. This corresponds to how many
	// pixels 1 unit represents.
	GetDistance() float64
	// GetProgress gets the current progress of Swiper
	GetProgress() float64
	// GetSwipeArea gets the area Swiper can start a swipe from for the given
	// direction and gesture type. This can be used to restrict swipes to only be
	// possible from a certain area, for example, to only allow edge swipes, or to
	// have a draggable element and ignore swipes elsewhere.
	//
	// Swipe area is only considered for direct swipes (as in, not initiated by
	// SwipeGroup).
	//
	// If not implemented, the default implementation returns the allocation of
	// Swiper, allowing swipes from anywhere.
	GetSwipeArea(navigationDirection NavigationDirection, isDrag bool, rect *gdk.Rectangle)
	// GetSwipeTracker gets the SwipeTracker used by this swipeable widget.
	GetSwipeTracker() *SwipeTracker
	// SwitchChild see HdySwipeable::child-switched.
	SwitchChild(index uint, duration int64)
}

type Swipeable struct {
	Caster
}

// native turns the current *Swipeable into the native C pointer type.
func (s *Swipeable) native() *C.HdySwipeable {
	return (*C.HdySwipeable)(unsafe.Pointer(s.Native()))
}

// EmitChildSwitched emits HdySwipeable::child-switched signal. This should be
// called when the widget switches visible child widget.
//
// duration can be 0 if the child is switched without animation.
func (s *Swipeable) EmitChildSwitched(index uint, duration int64) {
	v1 := C.guint(index)
	v2 := C.gint64(duration)

	C.hdy_swipeable_emit_child_switched(s.native(), v1, v2)
}

// GetCancelProgress gets the progress s will snap back to after the gesture is
// canceled.
func (s *Swipeable) GetCancelProgress() float64 {
	r := float64(C.hdy_swipeable_get_cancel_progress(s.native()))
	return r
}

// GetDistance gets the swipe distance of s. This corresponds to how many pixels
// 1 unit represents.
func (s *Swipeable) GetDistance() float64 {
	r := float64(C.hdy_swipeable_get_distance(s.native()))
	return r
}

// GetProgress gets the current progress of s
func (s *Swipeable) GetProgress() float64 {
	r := float64(C.hdy_swipeable_get_progress(s.native()))
	return r
}

// GetSwipeArea gets the area s can start a swipe from for the given direction
// and gesture type. This can be used to restrict swipes to only be possible
// from a certain area, for example, to only allow edge swipes, or to have a
// draggable element and ignore swipes elsewhere.
//
// Swipe area is only considered for direct swipes (as in, not initiated by
// SwipeGroup).
//
// If not implemented, the default implementation returns the allocation of s,
// allowing swipes from anywhere.
func (s *Swipeable) GetSwipeArea(navigationDirection NavigationDirection, isDrag bool, rect *gdk.Rectangle) {
	v1 := C.HdyNavigationDirection(navigationDirection)
	v2 := cbool(isDrag)
	v3 := (*C.GdkRectangle)(unsafe.Pointer(&rect.GdkRectangle))

	C.hdy_swipeable_get_swipe_area(s.native(), v1, v2, v3)
}

// GetSwipeTracker gets the SwipeTracker used by this swipeable widget.
func (s *Swipeable) GetSwipeTracker() *SwipeTracker {
	r := wrapSwipeTracker(unsafe.Pointer(C.hdy_swipeable_get_swipe_tracker(s.native())))
	return r
}

// SwitchChild see HdySwipeable::child-switched.
func (s *Swipeable) SwitchChild(index uint, duration int64) {
	v1 := C.guint(index)
	v2 := C.gint64(duration)

	C.hdy_swipeable_switch_child(s.native(), v1, v2)
}

// AvatarImageLoadFunc the returned Pixbuf is expected to be square with width
// and height set to size. The image is cropped to a circle without any scaling
// or transformation.
type AvatarImageLoadFunc func(size int) *gdk.Pixbuf

//export callbackAvatarImageLoadFunc
func callbackAvatarImageLoadFunc(size C.gint, userData C.gpointer) *C.GdkPixbuf {
	fn := callback.Get(uintptr(userData))
	if fn == nil {
		panic("callback for AvatarImageLoadFunc not found")
	}

	arg0 := int(size)

	v := fn.(AvatarImageLoadFunc)(arg0)
	if v != nil {
		v.Ref()
	}
	return (*C.GdkPixbuf)(unsafe.Pointer(v.Native()))
}

// ComboRowGetEnumValueNameFunc called for combo rows that are bound to an
// enumeration with (*ComboRow).SetForEnum() for each value from that
// enumeration.
type ComboRowGetEnumValueNameFunc func(value *EnumValueObject) string

//export callbackComboRowGetEnumValueNameFunc
func callbackComboRowGetEnumValueNameFunc(value *C.HdyEnumValueObject, userData C.gpointer) *C.gchar {
	fn := callback.Get(uintptr(userData))
	if fn == nil {
		panic("callback for ComboRowGetEnumValueNameFunc not found")
	}

	arg0 := wrapEnumValueObject(unsafe.Pointer(value))

	v := fn.(ComboRowGetEnumValueNameFunc)(arg0)
	return C.CString(v)
}

// ComboRowGetNameFunc called for combo rows that are bound to a Model with
// (*ComboRow).BindNameModel() for each item that gets added to the model.
type ComboRowGetNameFunc func(item *glib.Object) string

//export callbackComboRowGetNameFunc
func callbackComboRowGetNameFunc(item C.gpointer, userData C.gpointer) *C.gchar {
	fn := callback.Get(uintptr(userData))
	if fn == nil {
		panic("callback for ComboRowGetNameFunc not found")
	}

	arg0 := glib.Take(unsafe.Pointer(item))

	v := fn.(ComboRowGetNameFunc)(arg0)
	return C.CString(v)
}

// EaseOutCubic computes the ease out for t.
func EaseOutCubic(t float64) float64 {
	v1 := C.gdouble(t)
	r := float64(C.hdy_ease_out_cubic(v1))
	return r
}

// EnumValueRowName is a default implementation of ComboRowGetEnumValueNameFunc
// to be used with (*ComboRow).SetForEnum(). If the enumeration has a nickname,
// it will return it, otherwise it will return its name.
func EnumValueRowName(value *EnumValueObject) string {
	v1 := (*C.HdyEnumValueObject)(unsafe.Pointer(value.Native()))

	r := C.GoString(C.hdy_enum_value_row_name(v1, C.gpointer(uintptr(0))))
	return r
}

// GetEnableAnimations returns whether animations are enabled for that widget.
// This should be used when implementing an animated widget to know whether to
// animate it or not.
func GetEnableAnimations(widget gtk.IWidget) bool {
	v1 := cwidget(widget)
	r := gobool(C.hdy_get_enable_animations(v1))
	return r
}

// Init call this function just after initializing GTK, if you are using
// Application it means it must be called when the #GApplication::startup signal
// is emitted. If libhandy has already been initialized, the function will
// simply return.
//
// This makes sure translations, types, themes, and icons for the Handy library
// are set up properly.
func Init() {
	C.hdy_init()
}

type ActionRow struct {
	PreferencesRow

	// Interfaces
	gtk.Actionable
}

// wrapActionRow wraps the given pointer to *ActionRow.
func wrapActionRow(ptr unsafe.Pointer) *ActionRow {
	obj := glib.Take(ptr)
	return &ActionRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
				Bin: gtk.Bin{
					Container: gtk.Container{
						Widget: gtk.Widget{
							InitiallyUnowned: glib.InitiallyUnowned{
								Object: obj,
							},
						},
					},
				},
			},
		},
		Actionable: gtk.Actionable{obj},
	}
}

func marshalActionRow(p uintptr) (interface{}, error) {
	return wrapActionRow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ActionRowNew creates a new ActionRow.
func ActionRowNew() *ActionRow {
	return wrapActionRow(unsafe.Pointer(C.hdy_action_row_new()))
}

// native turns the current *ActionRow into the native C pointer type.
func (a *ActionRow) native() *C.HdyActionRow {
	return (*C.HdyActionRow)(gwidget(&a.PreferencesRow))
}

func (a *ActionRow) Activate() {
	C.hdy_action_row_activate(a.native())
}

// AddPrefix adds a prefix widget to a.
func (a *ActionRow) AddPrefix(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_action_row_add_prefix(a.native(), v1)
}

// GetActivatableWidget gets the widget activated when a is activated.
func (a *ActionRow) GetActivatableWidget() gtk.IWidget {
	r, err := castWidget(C.hdy_action_row_get_activatable_widget(a.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetIconName gets the icon name for a.
func (a *ActionRow) GetIconName() string {
	r := C.GoString(C.hdy_action_row_get_icon_name(a.native()))
	return r
}

// GetSubtitle gets the subtitle for a.
func (a *ActionRow) GetSubtitle() string {
	r := C.GoString(C.hdy_action_row_get_subtitle(a.native()))
	return r
}

// GetSubtitleLines gets the number of lines at the end of which the subtitle
// label will be ellipsized. If the value is 0, the number of lines won't be
// limited.
func (a *ActionRow) GetSubtitleLines() int {
	r := int(C.hdy_action_row_get_subtitle_lines(a.native()))
	return r
}

// GetTitleLines gets the number of lines at the end of which the title label
// will be ellipsized. If the value is 0, the number of lines won't be limited.
func (a *ActionRow) GetTitleLines() int {
	r := int(C.hdy_action_row_get_title_lines(a.native()))
	return r
}

// GetUseUnderline gets whether an embedded underline in the text of the title
// and subtitle labels indicates a mnemonic. See (*ActionRow).SetUseUnderline().
func (a *ActionRow) GetUseUnderline() bool {
	r := gobool(C.hdy_action_row_get_use_underline(a.native()))
	return r
}

// SetActivatableWidget sets the widget to activate when a is activated, either
// by clicking on it, by calling (*ActionRow).Activate(), or via mnemonics in
// the title or the subtitle. See the “use_underline” property to enable
// mnemonics.
//
// The target widget will be activated by emitting the
// GtkWidget::mnemonic-activate signal on it.
func (a *ActionRow) SetActivatableWidget(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_action_row_set_activatable_widget(a.native(), v1)
}

// SetIconName sets the icon name for a.
func (a *ActionRow) SetIconName(iconName string) {
	v1 := C.CString(iconName)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_action_row_set_icon_name(a.native(), v1)
}

// SetSubtitle sets the subtitle for a.
func (a *ActionRow) SetSubtitle(subtitle string) {
	v1 := C.CString(subtitle)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_action_row_set_subtitle(a.native(), v1)
}

// SetSubtitleLines sets the number of lines at the end of which the subtitle
// label will be ellipsized. If the value is 0, the number of lines won't be
// limited.
func (a *ActionRow) SetSubtitleLines(subtitleLines int) {
	v1 := C.gint(subtitleLines)
	C.hdy_action_row_set_subtitle_lines(a.native(), v1)
}

// SetTitleLines sets the number of lines at the end of which the title label
// will be ellipsized. If the value is 0, the number of lines won't be limited.
func (a *ActionRow) SetTitleLines(titleLines int) {
	v1 := C.gint(titleLines)
	C.hdy_action_row_set_title_lines(a.native(), v1)
}

// SetUseUnderline if true, an underline in the text of the title and subtitle
// labels indicates the next character should be used for the mnemonic
// accelerator key.
func (a *ActionRow) SetUseUnderline(useUnderline bool) {
	v1 := cbool(useUnderline)
	C.hdy_action_row_set_use_underline(a.native(), v1)
}

type ApplicationWindow struct {
	gtk.ApplicationWindow

	// Interfaces
	glib.ActionGroup
	glib.ActionMap
}

// wrapApplicationWindow wraps the given pointer to *ApplicationWindow.
func wrapApplicationWindow(ptr unsafe.Pointer) *ApplicationWindow {
	obj := glib.Take(ptr)
	return &ApplicationWindow{
		ApplicationWindow: gtk.ApplicationWindow{
			Window: gtk.Window{
				Bin: gtk.Bin{
					Container: gtk.Container{
						Widget: gtk.Widget{
							InitiallyUnowned: glib.InitiallyUnowned{
								Object: obj,
							},
						},
					},
				},
			},
		},
		ActionGroup: glib.ActionGroup{obj},
		ActionMap:   glib.ActionMap{obj},
	}
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	return wrapApplicationWindow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ApplicationWindowNew creates a new ApplicationWindow.
func ApplicationWindowNew() *ApplicationWindow {
	return wrapApplicationWindow(unsafe.Pointer(C.hdy_application_window_new()))
}

// native turns the current *ApplicationWindow into the native C pointer type.
func (a *ApplicationWindow) native() *C.HdyApplicationWindow {
	return (*C.HdyApplicationWindow)(gwidget(&a.ApplicationWindow))
}

type Avatar struct {
	gtk.DrawingArea
}

// wrapAvatar wraps the given pointer to *Avatar.
func wrapAvatar(ptr unsafe.Pointer) *Avatar {
	obj := glib.Take(ptr)
	return &Avatar{
		DrawingArea: gtk.DrawingArea{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalAvatar(p uintptr) (interface{}, error) {
	return wrapAvatar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// AvatarNew creates a new Avatar.
func AvatarNew(size int, text string, showInitials bool) *Avatar {
	v1 := C.gint(size)
	v2 := C.CString(text)
	defer C.free(unsafe.Pointer(v2))
	v3 := cbool(showInitials)

	return wrapAvatar(unsafe.Pointer(C.hdy_avatar_new(v1, v2, v3)))
}

// native turns the current *Avatar into the native C pointer type.
func (a *Avatar) native() *C.HdyAvatar {
	return (*C.HdyAvatar)(gwidget(&a.DrawingArea))
}

// DrawToPixbuf renders a into a pixbuf at size and scale_factor. This can be
// used to export the fallback avatar.
func (a *Avatar) DrawToPixbuf(size int, scaleFactor int) *gdk.Pixbuf {
	v1 := C.gint(size)
	v2 := C.gint(scaleFactor)

	obj := glib.Take(unsafe.Pointer(C.hdy_avatar_draw_to_pixbuf(a.native(), v1, v2)))
	r := &gdk.Pixbuf{
		Object: obj,
	}
	return r
}

// GetIconName gets the name of the icon in the icon theme to use when the icon
// should be displayed.
func (a *Avatar) GetIconName() string {
	r := C.GoString(C.hdy_avatar_get_icon_name(a.native()))
	return r
}

// GetShowInitials returns whether initials are used for the fallback or the
// icon.
func (a *Avatar) GetShowInitials() bool {
	r := gobool(C.hdy_avatar_get_show_initials(a.native()))
	return r
}

// GetSize returns the size of the avatar.
func (a *Avatar) GetSize() int {
	r := int(C.hdy_avatar_get_size(a.native()))
	return r
}

// GetText get the text used to generate the fallback initials and color
func (a *Avatar) GetText() string {
	r := C.GoString(C.hdy_avatar_get_text(a.native()))
	return r
}

// SetIconName sets the name of the icon in the icon theme to use when the icon
// should be displayed. If no name is set, the avatar-default-symbolic icon will
// be used. If the name doesn't match a valid icon, it is an error and no icon
// will be displayed. If the icon theme is changed, the image will be updated
// automatically.
func (a *Avatar) SetIconName(iconName string) {
	v1 := C.CString(iconName)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_avatar_set_icon_name(a.native(), v1)
}

// SetImageLoadFunc a callback which is called when the custom image need to be
// reloaded for some reason (e.g. scale-factor changes).
//
// This method is deprecated since version 1.2.
func (a *Avatar) SetImageLoadFunc(loadImage AvatarImageLoadFunc) {
	v1 := (*[0]byte)(C.callbackAvatarImageLoadFunc)

	C.hdy_avatar_set_image_load_func(a.native(), v1, C.gpointer(callback.Assign(loadImage)), (*[0]byte)(C.callbackDelete))
}

// SetShowInitials sets whether the initials should be shown on the fallback
// avatar or the icon.
func (a *Avatar) SetShowInitials(showInitials bool) {
	v1 := cbool(showInitials)
	C.hdy_avatar_set_show_initials(a.native(), v1)
}

// SetSize sets the size of the avatar.
func (a *Avatar) SetSize(size int) {
	v1 := C.gint(size)
	C.hdy_avatar_set_size(a.native(), v1)
}

// SetText set the text used to generate the fallback initials color
func (a *Avatar) SetText(text string) {
	v1 := C.CString(text)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_avatar_set_text(a.native(), v1)
}

type Carousel struct {
	gtk.EventBox

	// Interfaces
	gtk.Orientable
	Swiper
}

// wrapCarousel wraps the given pointer to *Carousel.
func wrapCarousel(ptr unsafe.Pointer) *Carousel {
	obj := glib.Take(ptr)
	return &Carousel{
		EventBox: gtk.EventBox{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: glib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
		Orientable: gtk.Orientable{obj},
		Swiper: &Swipeable{
			Caster: &gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCarousel(p uintptr) (interface{}, error) {
	return wrapCarousel(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// CarouselNew create a new Carousel widget.
func CarouselNew() *Carousel {
	return wrapCarousel(unsafe.Pointer(C.hdy_carousel_new()))
}

// native turns the current *Carousel into the native C pointer type.
func (c *Carousel) native() *C.HdyCarousel {
	return (*C.HdyCarousel)(gwidget(&c.EventBox))
}

// GetAllowLongSwipes whether to allow swiping for more than one page at a time.
// If the value is false, each swipe can only move to the adjacent pages.
func (c *Carousel) GetAllowLongSwipes() bool {
	r := gobool(C.hdy_carousel_get_allow_long_swipes(c.native()))
	return r
}

// GetAllowMouseDrag sets whether c can be dragged with mouse pointer
func (c *Carousel) GetAllowMouseDrag() bool {
	r := gobool(C.hdy_carousel_get_allow_mouse_drag(c.native()))
	return r
}

// GetAnimationDuration gets animation duration used by (*Carousel).ScrollTo().
func (c *Carousel) GetAnimationDuration() uint {
	r := uint(C.hdy_carousel_get_animation_duration(c.native()))
	return r
}

// GetInteractive gets whether c can be navigated.
func (c *Carousel) GetInteractive() bool {
	r := gobool(C.hdy_carousel_get_interactive(c.native()))
	return r
}

// GetNPages gets the number of pages in c.
func (c *Carousel) GetNPages() uint {
	r := uint(C.hdy_carousel_get_n_pages(c.native()))
	return r
}

// GetPosition gets current scroll position in c. It's unitless, 1 matches 1
// page.
func (c *Carousel) GetPosition() float64 {
	r := float64(C.hdy_carousel_get_position(c.native()))
	return r
}

// GetRevealDuration gets duration of the animation used when adding or removing
// pages in milliseconds.
func (c *Carousel) GetRevealDuration() uint {
	r := uint(C.hdy_carousel_get_reveal_duration(c.native()))
	return r
}

// GetSpacing gets spacing between pages in pixels.
func (c *Carousel) GetSpacing() uint {
	r := uint(C.hdy_carousel_get_spacing(c.native()))
	return r
}

// Insert inserts child into c at position position.
//
// If position is -1, or larger than the number of pages, child will be appended
// to the end.
func (c *Carousel) Insert(child gtk.IWidget, position int) {
	v1 := cwidget(child)
	v2 := C.gint(position)

	C.hdy_carousel_insert(c.native(), v1, v2)
}

// Prepend prepends child to c
func (c *Carousel) Prepend(child gtk.IWidget) {
	v1 := cwidget(child)
	C.hdy_carousel_prepend(c.native(), v1)
}

// Reorder moves child into position position.
//
// If position is -1, or larger than the number of pages, child will be moved to
// the end.
func (c *Carousel) Reorder(child gtk.IWidget, position int) {
	v1 := cwidget(child)
	v2 := C.gint(position)

	C.hdy_carousel_reorder(c.native(), v1, v2)
}

// ScrollTo scrolls to widget position with an animation.
// Carousel:animation-duration property can be used for controlling the
// duration.
func (c *Carousel) ScrollTo(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_carousel_scroll_to(c.native(), v1)
}

// ScrollToFull scrolls to widget position with an animation.
func (c *Carousel) ScrollToFull(widget gtk.IWidget, duration int64) {
	v1 := cwidget(widget)
	v2 := C.gint64(duration)

	C.hdy_carousel_scroll_to_full(c.native(), v1, v2)
}

// SetAllowLongSwipes sets whether to allow swiping for more than one page at a
// time. If the value is false, each swipe can only move to the adjacent pages.
func (c *Carousel) SetAllowLongSwipes(allowLongSwipes bool) {
	v1 := cbool(allowLongSwipes)
	C.hdy_carousel_set_allow_long_swipes(c.native(), v1)
}

// SetAllowMouseDrag sets whether c can be dragged with mouse pointer. If
// allow_mouse_drag is false, dragging is only available on touch.
func (c *Carousel) SetAllowMouseDrag(allowMouseDrag bool) {
	v1 := cbool(allowMouseDrag)
	C.hdy_carousel_set_allow_mouse_drag(c.native(), v1)
}

// SetAnimationDuration sets animation duration used by (*Carousel).ScrollTo().
func (c *Carousel) SetAnimationDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_carousel_set_animation_duration(c.native(), v1)
}

// SetInteractive sets whether c can be navigated. This can be used to
// temporarily disable a Carousel to only allow swiping in a certain state.
func (c *Carousel) SetInteractive(interactive bool) {
	v1 := cbool(interactive)
	C.hdy_carousel_set_interactive(c.native(), v1)
}

// SetRevealDuration sets duration of the animation used when adding or removing
// pages in milliseconds.
func (c *Carousel) SetRevealDuration(revealDuration uint) {
	v1 := C.guint(revealDuration)
	C.hdy_carousel_set_reveal_duration(c.native(), v1)
}

// SetSpacing sets spacing between pages in pixels.
func (c *Carousel) SetSpacing(spacing uint) {
	v1 := C.guint(spacing)
	C.hdy_carousel_set_spacing(c.native(), v1)
}

type CarouselIndicatorDots struct {
	gtk.DrawingArea

	// Interfaces
	gtk.Orientable
}

// wrapCarouselIndicatorDots wraps the given pointer to *CarouselIndicatorDots.
func wrapCarouselIndicatorDots(ptr unsafe.Pointer) *CarouselIndicatorDots {
	obj := glib.Take(ptr)
	return &CarouselIndicatorDots{
		DrawingArea: gtk.DrawingArea{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
	}
}

func marshalCarouselIndicatorDots(p uintptr) (interface{}, error) {
	return wrapCarouselIndicatorDots(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// CarouselIndicatorDotsNew create a new CarouselIndicatorDots widget.
func CarouselIndicatorDotsNew() *CarouselIndicatorDots {
	return wrapCarouselIndicatorDots(unsafe.Pointer(C.hdy_carousel_indicator_dots_new()))
}

// native turns the current *CarouselIndicatorDots into the native C pointer
// type.
func (c *CarouselIndicatorDots) native() *C.HdyCarouselIndicatorDots {
	return (*C.HdyCarouselIndicatorDots)(gwidget(&c.DrawingArea))
}

// GetCarousel get the Carousel the indicator uses.
//
// See: (*CarouselIndicatorDots).SetCarousel()
func (c *CarouselIndicatorDots) GetCarousel() *Carousel {
	r := wrapCarousel(unsafe.Pointer(C.hdy_carousel_indicator_dots_get_carousel(c.native())))
	return r
}

// SetCarousel sets the Carousel to use.
func (c *CarouselIndicatorDots) SetCarousel(carousel *Carousel) {
	v1 := (*C.HdyCarousel)(unsafe.Pointer(carousel.Widget.Native()))
	C.hdy_carousel_indicator_dots_set_carousel(c.native(), v1)
}

type CarouselIndicatorLines struct {
	gtk.DrawingArea

	// Interfaces
	gtk.Orientable
}

// wrapCarouselIndicatorLines wraps the given pointer to
// *CarouselIndicatorLines.
func wrapCarouselIndicatorLines(ptr unsafe.Pointer) *CarouselIndicatorLines {
	obj := glib.Take(ptr)
	return &CarouselIndicatorLines{
		DrawingArea: gtk.DrawingArea{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
	}
}

func marshalCarouselIndicatorLines(p uintptr) (interface{}, error) {
	return wrapCarouselIndicatorLines(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// CarouselIndicatorLinesNew create a new CarouselIndicatorLines widget.
func CarouselIndicatorLinesNew() *CarouselIndicatorLines {
	return wrapCarouselIndicatorLines(unsafe.Pointer(C.hdy_carousel_indicator_lines_new()))
}

// native turns the current *CarouselIndicatorLines into the native C pointer
// type.
func (c *CarouselIndicatorLines) native() *C.HdyCarouselIndicatorLines {
	return (*C.HdyCarouselIndicatorLines)(gwidget(&c.DrawingArea))
}

// GetCarousel get the Carousel the indicator uses.
//
// See: (*CarouselIndicatorLines).SetCarousel()
func (c *CarouselIndicatorLines) GetCarousel() *Carousel {
	r := wrapCarousel(unsafe.Pointer(C.hdy_carousel_indicator_lines_get_carousel(c.native())))
	return r
}

// SetCarousel sets the Carousel to use.
func (c *CarouselIndicatorLines) SetCarousel(carousel *Carousel) {
	v1 := (*C.HdyCarousel)(unsafe.Pointer(carousel.Widget.Native()))
	C.hdy_carousel_indicator_lines_set_carousel(c.native(), v1)
}

type Clamp struct {
	gtk.Bin

	// Interfaces
	gtk.Orientable
}

// wrapClamp wraps the given pointer to *Clamp.
func wrapClamp(ptr unsafe.Pointer) *Clamp {
	obj := glib.Take(ptr)
	return &Clamp{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
		Orientable: gtk.Orientable{obj},
	}
}

func marshalClamp(p uintptr) (interface{}, error) {
	return wrapClamp(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ClampNew creates a new Clamp.
func ClampNew() *Clamp {
	return wrapClamp(unsafe.Pointer(C.hdy_clamp_new()))
}

// native turns the current *Clamp into the native C pointer type.
func (c *Clamp) native() *C.HdyClamp {
	return (*C.HdyClamp)(gwidget(&c.Bin))
}

// GetMaximumSize gets the maximum size to allocate to the contained child. It
// is the width if c is horizontal, or the height if it is vertical.
func (c *Clamp) GetMaximumSize() int {
	r := int(C.hdy_clamp_get_maximum_size(c.native()))
	return r
}

// GetTighteningThreshold gets the size starting from which the clamp will
// tighten its grip on the child.
func (c *Clamp) GetTighteningThreshold() int {
	r := int(C.hdy_clamp_get_tightening_threshold(c.native()))
	return r
}

// SetMaximumSize sets the maximum size to allocate to the contained child. It
// is the width if c is horizontal, or the height if it is vertical.
func (c *Clamp) SetMaximumSize(maximumSize int) {
	v1 := C.gint(maximumSize)
	C.hdy_clamp_set_maximum_size(c.native(), v1)
}

// SetTighteningThreshold sets the size starting from which the clamp will
// tighten its grip on the child.
func (c *Clamp) SetTighteningThreshold(tighteningThreshold int) {
	v1 := C.gint(tighteningThreshold)
	C.hdy_clamp_set_tightening_threshold(c.native(), v1)
}

type ComboRow struct {
	ActionRow

	// Interfaces
	gtk.Actionable
}

// wrapComboRow wraps the given pointer to *ComboRow.
func wrapComboRow(ptr unsafe.Pointer) *ComboRow {
	obj := glib.Take(ptr)
	return &ComboRow{
		ActionRow: ActionRow{
			PreferencesRow: PreferencesRow{
				ListBoxRow: gtk.ListBoxRow{
					Bin: gtk.Bin{
						Container: gtk.Container{
							Widget: gtk.Widget{
								InitiallyUnowned: glib.InitiallyUnowned{
									Object: obj,
								},
							},
						},
					},
				},
			},
		},
		Actionable: gtk.Actionable{obj},
	}
}

func marshalComboRow(p uintptr) (interface{}, error) {
	return wrapComboRow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ComboRowNew creates a new ComboRow.
func ComboRowNew() *ComboRow {
	return wrapComboRow(unsafe.Pointer(C.hdy_combo_row_new()))
}

// native turns the current *ComboRow into the native C pointer type.
func (c *ComboRow) native() *C.HdyComboRow {
	return (*C.HdyComboRow)(gwidget(&c.ActionRow))
}

// BindModel binds model to c.
//
// If c was already bound to a model, that previous binding is destroyed.
//
// The contents of c are cleared and then filled with widgets that represent
// items from model. c is updated whenever model changes. If model is nil, c is
// left empty.
func (c *ComboRow) BindModel(model *glib.ListModel) {
	v1 := (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.hdy_combo_row_bind_model(c.native(), v1, nil, nil, C.gpointer(uintptr(0)), (*[0]byte)(C.callbackDelete))
}

// BindNameModel binds model to c.
//
// If c was already bound to a model, that previous binding is destroyed.
//
// The contents of c are cleared and then filled with widgets that represent
// items from model. c is updated whenever model changes. If model is nil, c is
// left empty.
//
// This is more convenient to use than (*ComboRow).BindModel() if you want to
// represent items of the model with names.
func (c *ComboRow) BindNameModel(model *glib.ListModel, getNameFunc ComboRowGetNameFunc) {
	v1 := (*C.GListModel)(unsafe.Pointer(model.Native()))
	v2 := (*[0]byte)(C.callbackComboRowGetNameFunc)

	C.hdy_combo_row_bind_name_model(c.native(), v1, v2, C.gpointer(callback.Assign(getNameFunc)), (*[0]byte)(C.callbackDelete))
}

// GetModel gets the model bound to c, or nil if none is bound.
func (c *ComboRow) GetModel() *glib.ListModel {
	obj := glib.Take(unsafe.Pointer(C.hdy_combo_row_get_model(c.native())))
	r := &glib.ListModel{
		Object: obj,
	}
	return r
}

// GetSelectedIndex gets the index of the selected item in its Model.
func (c *ComboRow) GetSelectedIndex() int {
	r := int(C.hdy_combo_row_get_selected_index(c.native()))
	return r
}

// GetUseSubtitle gets whether the current value of c should be displayed as its
// subtitle.
func (c *ComboRow) GetUseSubtitle() bool {
	r := gobool(C.hdy_combo_row_get_use_subtitle(c.native()))
	return r
}

// SetForEnum creates a model for enum_type and binds it to c. The items of the
// model will be EnumValueObject objects.
//
// If c was already bound to a model, that previous binding is destroyed.
//
// The contents of c are cleared and then filled with widgets that represent
// items from model. c is updated whenever model changes. If model is nil, c is
// left empty.
//
// This is more convenient to use than (*ComboRow).BindNameModel() if you want
// to represent values of an enumeration with names.
//
// See EnumValueRowName().
func (c *ComboRow) SetForEnum(enumType glib.Type, getNameFunc ComboRowGetEnumValueNameFunc) {
	v1 := C.GType(enumType)
	v2 := (*[0]byte)(C.callbackComboRowGetEnumValueNameFunc)

	C.hdy_combo_row_set_for_enum(c.native(), v1, v2, C.gpointer(callback.Assign(getNameFunc)), (*[0]byte)(C.callbackDelete))
}

// SetGetNameFunc sets a closure to convert items into names. See
// HdyComboRow:use-subtitle.
func (c *ComboRow) SetGetNameFunc(getNameFunc ComboRowGetNameFunc) {
	v1 := (*[0]byte)(C.callbackComboRowGetNameFunc)

	C.hdy_combo_row_set_get_name_func(c.native(), v1, C.gpointer(callback.Assign(getNameFunc)), (*[0]byte)(C.callbackDelete))
}

// SetSelectedIndex sets the index of the selected item in its Model.
func (c *ComboRow) SetSelectedIndex(selectedIndex int) {
	v1 := C.gint(selectedIndex)
	C.hdy_combo_row_set_selected_index(c.native(), v1)
}

// SetUseSubtitle sets whether the current value of c should be displayed as its
// subtitle.
//
// If true, you should not access HdyActionRow:subtitle.
func (c *ComboRow) SetUseSubtitle(useSubtitle bool) {
	v1 := cbool(useSubtitle)
	C.hdy_combo_row_set_use_subtitle(c.native(), v1)
}

type Deck struct {
	gtk.Container

	// Interfaces
	gtk.Orientable
	Swiper
}

// wrapDeck wraps the given pointer to *Deck.
func wrapDeck(ptr unsafe.Pointer) *Deck {
	obj := glib.Take(ptr)
	return &Deck{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
		Swiper: &Swipeable{
			Caster: &gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalDeck(p uintptr) (interface{}, error) {
	return wrapDeck(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// DeckNew creates a new Deck.
func DeckNew() *Deck {
	return wrapDeck(unsafe.Pointer(C.hdy_deck_new()))
}

// native turns the current *Deck into the native C pointer type.
func (d *Deck) native() *C.HdyDeck {
	return (*C.HdyDeck)(gwidget(&d.Container))
}

// GetAdjacentChild gets the previous or next child, or nil if it doesn't exist.
// This will be the same widget (*Deck).Navigate() will navigate to.
func (d *Deck) GetAdjacentChild(direction NavigationDirection) gtk.IWidget {
	v1 := C.HdyNavigationDirection(direction)
	r, err := castWidget(C.hdy_deck_get_adjacent_child(d.native(), v1))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetCanSwipeBack returns whether the Deck allows swiping to the previous
// child.
func (d *Deck) GetCanSwipeBack() bool {
	r := gobool(C.hdy_deck_get_can_swipe_back(d.native()))
	return r
}

// GetCanSwipeForward returns whether the Deck allows swiping to the next child.
func (d *Deck) GetCanSwipeForward() bool {
	r := gobool(C.hdy_deck_get_can_swipe_forward(d.native()))
	return r
}

// GetChildByName finds the child of d with the name given as the argument.
// Returns nil if there is no child with this name.
func (d *Deck) GetChildByName(name string) gtk.IWidget {
	v1 := C.CString(name)
	defer C.free(unsafe.Pointer(v1))
	r, err := castWidget(C.hdy_deck_get_child_by_name(d.native(), v1))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetHomogeneous gets whether d is homogeneous for the given orientation. See
// (*Deck).SetHomogeneous().
func (d *Deck) GetHomogeneous(orientation gtk.Orientation) bool {
	v1 := C.GtkOrientation(orientation)
	r := gobool(C.hdy_deck_get_homogeneous(d.native(), v1))
	return r
}

// GetInterpolateSize returns whether the Deck is set up to interpolate between
// the sizes of children on page switch.
func (d *Deck) GetInterpolateSize() bool {
	r := gobool(C.hdy_deck_get_interpolate_size(d.native()))
	return r
}

// GetTransitionDuration returns the amount of time (in milliseconds) that
// transitions between children in d will take.
func (d *Deck) GetTransitionDuration() uint {
	r := uint(C.hdy_deck_get_transition_duration(d.native()))
	return r
}

// GetTransitionRunning returns whether d is currently in a transition from one
// page to another.
func (d *Deck) GetTransitionRunning() bool {
	r := gobool(C.hdy_deck_get_transition_running(d.native()))
	return r
}

// GetTransitionType gets the type of animation that will be used for
// transitions between children in d.
func (d *Deck) GetTransitionType() DeckTransitionType {
	r := DeckTransitionType(C.hdy_deck_get_transition_type(d.native()))
	return r
}

// GetVisibleChild gets the visible child widget.
func (d *Deck) GetVisibleChild() gtk.IWidget {
	r, err := castWidget(C.hdy_deck_get_visible_child(d.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetVisibleChildName gets the name of the currently visible child widget.
func (d *Deck) GetVisibleChildName() string {
	r := C.GoString(C.hdy_deck_get_visible_child_name(d.native()))
	return r
}

// InsertChildAfter inserts child in the position after sibling in the list of
// children. If sibling is nil, insert child at the first position.
func (d *Deck) InsertChildAfter(child gtk.IWidget, sibling gtk.IWidget) {
	v1 := cwidget(child)
	v2 := cwidget(sibling)

	C.hdy_deck_insert_child_after(d.native(), v1, v2)
}

// Navigate switches to the previous or next child, similar to performing a
// swipe gesture to go in direction.
func (d *Deck) Navigate(direction NavigationDirection) bool {
	v1 := C.HdyNavigationDirection(direction)
	r := gobool(C.hdy_deck_navigate(d.native(), v1))
	return r
}

// Prepend inserts child at the first position in d.
func (d *Deck) Prepend(child gtk.IWidget) {
	v1 := cwidget(child)
	C.hdy_deck_prepend(d.native(), v1)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// children. If sibling is nil, move child to the first position.
func (d *Deck) ReorderChildAfter(child gtk.IWidget, sibling gtk.IWidget) {
	v1 := cwidget(child)
	v2 := cwidget(sibling)

	C.hdy_deck_reorder_child_after(d.native(), v1, v2)
}

// SetCanSwipeBack sets whether or not d allows switching to the previous child
// via a swipe gesture.
func (d *Deck) SetCanSwipeBack(canSwipeBack bool) {
	v1 := cbool(canSwipeBack)
	C.hdy_deck_set_can_swipe_back(d.native(), v1)
}

// SetCanSwipeForward sets whether or not d allows switching to the next child
// via a swipe gesture.
func (d *Deck) SetCanSwipeForward(canSwipeForward bool) {
	v1 := cbool(canSwipeForward)
	C.hdy_deck_set_can_swipe_forward(d.native(), v1)
}

// SetHomogeneous sets the Deck to be homogeneous or not for the given
// orientation. If it is homogeneous, the Deck will request the same width or
// height for all its children depending on the orientation. If it isn't, the
// deck may change width or height when a different child becomes visible.
func (d *Deck) SetHomogeneous(orientation gtk.Orientation, homogeneous bool) {
	v1 := C.GtkOrientation(orientation)
	v2 := cbool(homogeneous)

	C.hdy_deck_set_homogeneous(d.native(), v1, v2)
}

// SetInterpolateSize sets whether or not d will interpolate its size when
// changing the visible child. If the Deck:interpolate-size property is set to
// true, d will interpolate its size between the current one and the one it'll
// take after changing the visible child, according to the set transition
// duration.
func (d *Deck) SetInterpolateSize(interpolateSize bool) {
	v1 := cbool(interpolateSize)
	C.hdy_deck_set_interpolate_size(d.native(), v1)
}

// SetTransitionDuration sets the duration that transitions between children in
// d will take.
func (d *Deck) SetTransitionDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_deck_set_transition_duration(d.native(), v1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between children in d.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the child that is about to become
// current.
func (d *Deck) SetTransitionType(transition DeckTransitionType) {
	v1 := C.HdyDeckTransitionType(transition)
	C.hdy_deck_set_transition_type(d.native(), v1)
}

// SetVisibleChild makes visible_child visible using a transition determined by
// HdyDeck:transition-type and HdyDeck:transition-duration. The transition can
// be cancelled by the user, in which case visible child will change back to the
// previously visible child.
func (d *Deck) SetVisibleChild(visibleChild gtk.IWidget) {
	v1 := cwidget(visibleChild)
	C.hdy_deck_set_visible_child(d.native(), v1)
}

// SetVisibleChildName makes the child with the name name visible.
//
// See (*Deck).SetVisibleChild() for more details.
func (d *Deck) SetVisibleChildName(name string) {
	v1 := C.CString(name)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_deck_set_visible_child_name(d.native(), v1)
}

type EnumValueObject struct {
	*glib.Object
}

// wrapEnumValueObject wraps the given pointer to *EnumValueObject.
func wrapEnumValueObject(ptr unsafe.Pointer) *EnumValueObject {
	obj := glib.Take(ptr)
	return &EnumValueObject{
		Object: obj,
	}
}

func marshalEnumValueObject(p uintptr) (interface{}, error) {
	return wrapEnumValueObject(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// native turns the current *EnumValueObject into the native C pointer type.
func (e *EnumValueObject) native() *C.HdyEnumValueObject {
	return (*C.HdyEnumValueObject)(unsafe.Pointer(e.Native()))
}

func (e *EnumValueObject) GetName() string {
	r := C.GoString(C.hdy_enum_value_object_get_name(e.native()))
	return r
}
func (e *EnumValueObject) GetNick() string {
	r := C.GoString(C.hdy_enum_value_object_get_nick(e.native()))
	return r
}
func (e *EnumValueObject) GetValue() int {
	r := int(C.hdy_enum_value_object_get_value(e.native()))
	return r
}

type ExpanderRow struct {
	PreferencesRow

	// Interfaces
	gtk.Actionable
}

// wrapExpanderRow wraps the given pointer to *ExpanderRow.
func wrapExpanderRow(ptr unsafe.Pointer) *ExpanderRow {
	obj := glib.Take(ptr)
	return &ExpanderRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
				Bin: gtk.Bin{
					Container: gtk.Container{
						Widget: gtk.Widget{
							InitiallyUnowned: glib.InitiallyUnowned{
								Object: obj,
							},
						},
					},
				},
			},
		},
		Actionable: gtk.Actionable{obj},
	}
}

func marshalExpanderRow(p uintptr) (interface{}, error) {
	return wrapExpanderRow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ExpanderRowNew creates a new ExpanderRow.
func ExpanderRowNew() *ExpanderRow {
	return wrapExpanderRow(unsafe.Pointer(C.hdy_expander_row_new()))
}

// native turns the current *ExpanderRow into the native C pointer type.
func (e *ExpanderRow) native() *C.HdyExpanderRow {
	return (*C.HdyExpanderRow)(gwidget(&e.PreferencesRow))
}

// AddAction adds an action widget to e.
func (e *ExpanderRow) AddAction(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_expander_row_add_action(e.native(), v1)
}

// AddPrefix adds a prefix widget to e.
func (e *ExpanderRow) AddPrefix(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_expander_row_add_prefix(e.native(), v1)
}

// GetEnableExpansion gets whether the expansion of e is enabled.
func (e *ExpanderRow) GetEnableExpansion() bool {
	r := gobool(C.hdy_expander_row_get_enable_expansion(e.native()))
	return r
}
func (e *ExpanderRow) GetExpanded() bool {
	r := gobool(C.hdy_expander_row_get_expanded(e.native()))
	return r
}

// GetIconName gets the icon name for e.
func (e *ExpanderRow) GetIconName() string {
	r := C.GoString(C.hdy_expander_row_get_icon_name(e.native()))
	return r
}

// GetShowEnableSwitch gets whether the switch enabling the expansion of e is
// visible.
func (e *ExpanderRow) GetShowEnableSwitch() bool {
	r := gobool(C.hdy_expander_row_get_show_enable_switch(e.native()))
	return r
}

// GetSubtitle gets the subtitle for e.
func (e *ExpanderRow) GetSubtitle() string {
	r := C.GoString(C.hdy_expander_row_get_subtitle(e.native()))
	return r
}

// GetUseUnderline gets whether an embedded underline in the text of the title
// and subtitle labels indicates a mnemonic. See
// (*ExpanderRow).SetUseUnderline().
func (e *ExpanderRow) GetUseUnderline() bool {
	r := gobool(C.hdy_expander_row_get_use_underline(e.native()))
	return r
}

// SetEnableExpansion sets whether the expansion of e is enabled.
func (e *ExpanderRow) SetEnableExpansion(enableExpansion bool) {
	v1 := cbool(enableExpansion)
	C.hdy_expander_row_set_enable_expansion(e.native(), v1)
}
func (e *ExpanderRow) SetExpanded(expanded bool) {
	v1 := cbool(expanded)
	C.hdy_expander_row_set_expanded(e.native(), v1)
}

// SetIconName sets the icon name for e.
func (e *ExpanderRow) SetIconName(iconName string) {
	v1 := C.CString(iconName)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_expander_row_set_icon_name(e.native(), v1)
}

// SetShowEnableSwitch sets whether the switch enabling the expansion of e is
// visible.
func (e *ExpanderRow) SetShowEnableSwitch(showEnableSwitch bool) {
	v1 := cbool(showEnableSwitch)
	C.hdy_expander_row_set_show_enable_switch(e.native(), v1)
}

// SetSubtitle sets the subtitle for e.
func (e *ExpanderRow) SetSubtitle(subtitle string) {
	v1 := C.CString(subtitle)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_expander_row_set_subtitle(e.native(), v1)
}

// SetUseUnderline if true, an underline in the text of the title and subtitle
// labels indicates the next character should be used for the mnemonic
// accelerator key.
func (e *ExpanderRow) SetUseUnderline(useUnderline bool) {
	v1 := cbool(useUnderline)
	C.hdy_expander_row_set_use_underline(e.native(), v1)
}

type Flap struct {
	gtk.Container

	// Interfaces
	gtk.Orientable
	Swiper
}

// wrapFlap wraps the given pointer to *Flap.
func wrapFlap(ptr unsafe.Pointer) *Flap {
	obj := glib.Take(ptr)
	return &Flap{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
		Swiper: &Swipeable{
			Caster: &gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalFlap(p uintptr) (interface{}, error) {
	return wrapFlap(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// FlapNew creates a new Flap.
func FlapNew() *Flap {
	return wrapFlap(unsafe.Pointer(C.hdy_flap_new()))
}

// native turns the current *Flap into the native C pointer type.
func (f *Flap) native() *C.HdyFlap {
	return (*C.HdyFlap)(gwidget(&f.Container))
}

// GetContent gets the content widget for f
func (f *Flap) GetContent() gtk.IWidget {
	r, err := castWidget(C.hdy_flap_get_content(f.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetFlap gets the flap widget for f
func (f *Flap) GetFlap() gtk.IWidget {
	r, err := castWidget(C.hdy_flap_get_flap(f.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetFlapPosition gets the flap position for f.
func (f *Flap) GetFlapPosition() gtk.PackType {
	r := gtk.PackType(C.hdy_flap_get_flap_position(f.native()))
	return r
}

// GetFoldDuration returns the amount of time (in milliseconds) that fold
// transitions in f will take.
func (f *Flap) GetFoldDuration() uint {
	r := uint(C.hdy_flap_get_fold_duration(f.native()))
	return r
}

// GetFoldPolicy gets the current fold policy of f. See (*Flap).SetFoldPolicy().
func (f *Flap) GetFoldPolicy() FlapFoldPolicy {
	r := FlapFoldPolicy(C.hdy_flap_get_fold_policy(f.native()))
	return r
}

// GetFolded gets whether f is currently folded.
//
// See Flap:fold-policy.
func (f *Flap) GetFolded() bool {
	r := gobool(C.hdy_flap_get_folded(f.native()))
	return r
}

// GetLocked gets whether f is locked.
func (f *Flap) GetLocked() bool {
	r := gobool(C.hdy_flap_get_locked(f.native()))
	return r
}

// GetModal gets whether the f is modal. See (*Flap).SetModal().
func (f *Flap) GetModal() bool {
	r := gobool(C.hdy_flap_get_modal(f.native()))
	return r
}

// GetRevealDuration returns the amount of time (in milliseconds) that reveal
// transitions in f will take.
func (f *Flap) GetRevealDuration() uint {
	r := uint(C.hdy_flap_get_reveal_duration(f.native()))
	return r
}

// GetRevealFlap gets whether the flap widget is revealed for f.
func (f *Flap) GetRevealFlap() bool {
	r := gobool(C.hdy_flap_get_reveal_flap(f.native()))
	return r
}

// GetRevealProgress gets the current reveal transition progress for f. 0 means
// fully hidden, 1 means fully revealed. See Flap:reveal-flap.
func (f *Flap) GetRevealProgress() float64 {
	r := float64(C.hdy_flap_get_reveal_progress(f.native()))
	return r
}

// GetSeparator gets the separator widget for f.
func (f *Flap) GetSeparator() gtk.IWidget {
	r, err := castWidget(C.hdy_flap_get_separator(f.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetSwipeToClose gets whether f can be closed with a swipe gesture.
func (f *Flap) GetSwipeToClose() bool {
	r := gobool(C.hdy_flap_get_swipe_to_close(f.native()))
	return r
}

// GetSwipeToOpen gets whether f can be opened with a swipe gesture.
func (f *Flap) GetSwipeToOpen() bool {
	r := gobool(C.hdy_flap_get_swipe_to_open(f.native()))
	return r
}

// GetTransitionType gets the type of animation that will be used for reveal and
// fold transitions in f.
func (f *Flap) GetTransitionType() FlapTransitionType {
	r := FlapTransitionType(C.hdy_flap_get_transition_type(f.native()))
	return r
}

// SetContent sets the content widget for f, always displayed when unfolded, and
// partially visible when folded.
func (f *Flap) SetContent(content gtk.IWidget) {
	v1 := cwidget(content)
	C.hdy_flap_set_content(f.native(), v1)
}

// SetFlap sets the flap widget for f, only visible when Flap:reveal-progress is
// greater than 0.
func (f *Flap) SetFlap(flap gtk.IWidget) {
	v1 := cwidget(flap)
	C.hdy_flap_set_flap(f.native(), v1)
}

// SetFlapPosition sets the flap position for f. If GTK_PACK_START, the flap is
// displayed before the content, if GTK_PACK_END, it's displayed after the
// content.
func (f *Flap) SetFlapPosition(position gtk.PackType) {
	v1 := C.GtkPackType(position)
	C.hdy_flap_set_flap_position(f.native(), v1)
}

// SetFoldDuration sets the duration that fold transitions in f will take.
func (f *Flap) SetFoldDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_flap_set_fold_duration(f.native(), v1)
}

// SetFoldPolicy sets the current fold policy for f. See FlapFoldPolicy for
// available policies.
func (f *Flap) SetFoldPolicy(policy FlapFoldPolicy) {
	v1 := C.HdyFlapFoldPolicy(policy)
	C.hdy_flap_set_fold_policy(f.native(), v1)
}

// SetLocked sets whether f is locked.
//
// If false, folding f when the flap is revealed automatically closes it, and
// unfolding it when the flap is not revealed opens it. If true,
// Flap:reveal-flap value never changes on its own.
func (f *Flap) SetLocked(locked bool) {
	v1 := cbool(locked)
	C.hdy_flap_set_locked(f.native(), v1)
}

// SetModal sets whether the f can be closed with a click.
//
// If modal is true, clicking the content widget while flap is revealed, or
// pressing Escape key, will close the flap. If false, clicks are passed through
// to the content widget.
func (f *Flap) SetModal(modal bool) {
	v1 := cbool(modal)
	C.hdy_flap_set_modal(f.native(), v1)
}

// SetRevealDuration sets the duration that reveal transitions in f will take.
func (f *Flap) SetRevealDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_flap_set_reveal_duration(f.native(), v1)
}

// SetRevealFlap sets whether the flap widget is revealed for f.
func (f *Flap) SetRevealFlap(revealFlap bool) {
	v1 := cbool(revealFlap)
	C.hdy_flap_set_reveal_flap(f.native(), v1)
}

// SetSeparator sets the separator widget for f, displayed between content and
// flap when there's no shadow to display. When exactly it's visible depends on
// the Flap:transition-type value. If nil, no separator will be used.
func (f *Flap) SetSeparator(separator gtk.IWidget) {
	v1 := cwidget(separator)
	C.hdy_flap_set_separator(f.native(), v1)
}

// SetSwipeToClose sets whether f can be closed with a swipe gesture.
//
// The area that can be swiped depends on the Flap:transition-type value.
func (f *Flap) SetSwipeToClose(swipeToClose bool) {
	v1 := cbool(swipeToClose)
	C.hdy_flap_set_swipe_to_close(f.native(), v1)
}

// SetSwipeToOpen sets whether f can be opened with a swipe gesture.
//
// The area that can be swiped depends on the Flap:transition-type value.
func (f *Flap) SetSwipeToOpen(swipeToOpen bool) {
	v1 := cbool(swipeToOpen)
	C.hdy_flap_set_swipe_to_open(f.native(), v1)
}

// SetTransitionType sets the type of animation that will be used for reveal and
// fold transitions in f.
//
// Flap:flap is transparent by default, which means the content will be seen
// through it with HDY_FLAP_TRANSITION_TYPE_OVER transitions; add the
// .background style class to it if this is unwanted.
func (f *Flap) SetTransitionType(transitionType FlapTransitionType) {
	v1 := C.HdyFlapTransitionType(transitionType)
	C.hdy_flap_set_transition_type(f.native(), v1)
}

type HeaderBar struct {
	gtk.Container
}

// wrapHeaderBar wraps the given pointer to *HeaderBar.
func wrapHeaderBar(ptr unsafe.Pointer) *HeaderBar {
	obj := glib.Take(ptr)
	return &HeaderBar{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	return wrapHeaderBar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// HeaderBarNew creates a new HeaderBar widget.
func HeaderBarNew() *HeaderBar {
	return wrapHeaderBar(unsafe.Pointer(C.hdy_header_bar_new()))
}

// native turns the current *HeaderBar into the native C pointer type.
func (h *HeaderBar) native() *C.HdyHeaderBar {
	return (*C.HdyHeaderBar)(gwidget(&h.Container))
}

// GetCenteringPolicy gets the policy h follows to horizontally align its center
// widget.
func (h *HeaderBar) GetCenteringPolicy() CenteringPolicy {
	r := CenteringPolicy(C.hdy_header_bar_get_centering_policy(h.native()))
	return r
}

// GetCustomTitle retrieves the custom title widget of the header. See
// (*HeaderBar).SetCustomTitle().
func (h *HeaderBar) GetCustomTitle() gtk.IWidget {
	r, err := castWidget(C.hdy_header_bar_get_custom_title(h.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetDecorationLayout gets the decoration layout set with
// (*HeaderBar).SetDecorationLayout().
func (h *HeaderBar) GetDecorationLayout() string {
	r := C.GoString(C.hdy_header_bar_get_decoration_layout(h.native()))
	return r
}

// GetHasSubtitle retrieves whether the header bar reserves space for a
// subtitle, regardless if one is currently set or not.
func (h *HeaderBar) GetHasSubtitle() bool {
	r := gobool(C.hdy_header_bar_get_has_subtitle(h.native()))
	return r
}

// GetInterpolateSize gets whether h should interpolate its size on visible
// child change.
//
// See (*HeaderBar).SetInterpolateSize().
func (h *HeaderBar) GetInterpolateSize() bool {
	r := gobool(C.hdy_header_bar_get_interpolate_size(h.native()))
	return r
}

// GetShowCloseButton returns whether this header bar shows the standard window
// decorations.
func (h *HeaderBar) GetShowCloseButton() bool {
	r := gobool(C.hdy_header_bar_get_show_close_button(h.native()))
	return r
}

// GetSubtitle retrieves the subtitle of the header. See
// (*HeaderBar).SetSubtitle().
func (h *HeaderBar) GetSubtitle() string {
	r := C.GoString(C.hdy_header_bar_get_subtitle(h.native()))
	return r
}

// GetTitle retrieves the title of the header. See (*HeaderBar).SetTitle().
func (h *HeaderBar) GetTitle() string {
	r := C.GoString(C.hdy_header_bar_get_title(h.native()))
	return r
}

// GetTransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in h will take.
func (h *HeaderBar) GetTransitionDuration() uint {
	r := uint(C.hdy_header_bar_get_transition_duration(h.native()))
	return r
}

// GetTransitionRunning returns whether the h is currently in a transition from
// one page to another.
func (h *HeaderBar) GetTransitionRunning() bool {
	r := gobool(C.hdy_header_bar_get_transition_running(h.native()))
	return r
}

// PackEnd adds child to h:, packed with reference to the end of the h:.
func (h *HeaderBar) PackEnd(child gtk.IWidget) {
	v1 := cwidget(child)
	C.hdy_header_bar_pack_end(h.native(), v1)
}

// PackStart adds child to h:, packed with reference to the start of the h:.
func (h *HeaderBar) PackStart(child gtk.IWidget) {
	v1 := cwidget(child)
	C.hdy_header_bar_pack_start(h.native(), v1)
}

// SetCenteringPolicy sets the policy h must follow to horizontally align its
// center widget.
func (h *HeaderBar) SetCenteringPolicy(centeringPolicy CenteringPolicy) {
	v1 := C.HdyCenteringPolicy(centeringPolicy)
	C.hdy_header_bar_set_centering_policy(h.native(), v1)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes any
// title set by (*HeaderBar).SetTitle() or (*HeaderBar).SetSubtitle(). To
// achieve the same style as the builtin title and subtitle, use the “title”
// and “subtitle” style classes.
//
// You should set the custom title to nil, for the header title label to be
// visible again.
func (h *HeaderBar) SetCustomTitle(titleWidget gtk.IWidget) {
	v1 := cwidget(titleWidget)
	C.hdy_header_bar_set_custom_title(h.native(), v1)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close, icon (the window icon)
// and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left,
// and minimize, maximize and close buttons on the right.
func (h *HeaderBar) SetDecorationLayout(layout string) {
	v1 := C.CString(layout)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_header_bar_set_decoration_layout(h.native(), v1)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
func (h *HeaderBar) SetHasSubtitle(setting bool) {
	v1 := cbool(setting)
	C.hdy_header_bar_set_has_subtitle(h.native(), v1)
}

// SetInterpolateSize sets whether or not h will interpolate the size of its
// opposing orientation when changing the visible child. If true, h will
// interpolate its size between the one of the previous visible child and the
// one of the new visible child, according to the set transition duration and
// the orientation, e.g. if h is horizontal, it will interpolate the its height.
func (h *HeaderBar) SetInterpolateSize(interpolateSize bool) {
	v1 := cbool(interpolateSize)
	C.hdy_header_bar_set_interpolate_size(h.native(), v1)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
func (h *HeaderBar) SetShowCloseButton(setting bool) {
	v1 := cbool(setting)
	C.hdy_header_bar_set_show_close_button(h.native(), v1)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a user
// an additional detail to help them identify the current view.
//
// Note that HdyHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the HeaderBar:has-subtitle
// property to false.
func (h *HeaderBar) SetSubtitle(subtitle string) {
	v1 := C.CString(subtitle)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_header_bar_set_subtitle(h.native(), v1)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the application
// name.
func (h *HeaderBar) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_header_bar_set_title(h.native(), v1)
}

// SetTransitionDuration sets the duration that transitions between pages in h
// will take.
func (h *HeaderBar) SetTransitionDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_header_bar_set_transition_duration(h.native(), v1)
}

type HeaderGroup struct {
	*glib.Object
}

// wrapHeaderGroup wraps the given pointer to *HeaderGroup.
func wrapHeaderGroup(ptr unsafe.Pointer) *HeaderGroup {
	obj := glib.Take(ptr)
	return &HeaderGroup{
		Object: obj,
	}
}

func marshalHeaderGroup(p uintptr) (interface{}, error) {
	return wrapHeaderGroup(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// HeaderGroupNew creates a new HeaderGroup.
func HeaderGroupNew() *HeaderGroup {
	return wrapHeaderGroup(unsafe.Pointer(C.hdy_header_group_new()))
}

// native turns the current *HeaderGroup into the native C pointer type.
func (h *HeaderGroup) native() *C.HdyHeaderGroup {
	return (*C.HdyHeaderGroup)(unsafe.Pointer(h.Native()))
}

// AddGtkHeaderBar adds header_bar to h. When the widget is destroyed or no
// longer referenced elsewhere, it will be removed from the header group.
func (h *HeaderGroup) AddGtkHeaderBar(headerBar *gtk.HeaderBar) {
	v1 := (*C.GtkHeaderBar)(unsafe.Pointer(headerBar.Widget.Native()))
	C.hdy_header_group_add_gtk_header_bar(h.native(), v1)
}

// AddHeaderBar adds header_bar to h. When the widget is destroyed or no longer
// referenced elsewhere, it will be removed from the header group.
func (h *HeaderGroup) AddHeaderBar(headerBar *HeaderBar) {
	v1 := (*C.HdyHeaderBar)(unsafe.Pointer(headerBar.Widget.Native()))
	C.hdy_header_group_add_header_bar(h.native(), v1)
}

// AddHeaderGroup adds header_group to h. When the nested group is no longer
// referenced elsewhere, it will be removed from the header group.
func (h *HeaderGroup) AddHeaderGroup(headerGroup *HeaderGroup) {
	v1 := (*C.HdyHeaderGroup)(unsafe.Pointer(headerGroup.Native()))
	C.hdy_header_group_add_header_group(h.native(), v1)
}

// GetChildren returns the list of children associated with h.
func (h *HeaderGroup) GetChildren() *glib.SList {
	r := glib.WrapSList(uintptr(unsafe.Pointer(C.hdy_header_group_get_children(h.native()))))
	return r
}

// GetDecorateAll gets whether the elements of the group should all receive the
// full decoration.
func (h *HeaderGroup) GetDecorateAll() bool {
	r := gobool(C.hdy_header_group_get_decorate_all(h.native()))
	return r
}

// RemoveChild removes child from h.
func (h *HeaderGroup) RemoveChild(child *HeaderGroupChild) {
	v1 := (*C.HdyHeaderGroupChild)(unsafe.Pointer(child.Native()))
	C.hdy_header_group_remove_child(h.native(), v1)
}

// RemoveGtkHeaderBar removes header_bar from h.
func (h *HeaderGroup) RemoveGtkHeaderBar(headerBar *gtk.HeaderBar) {
	v1 := (*C.GtkHeaderBar)(unsafe.Pointer(headerBar.Widget.Native()))
	C.hdy_header_group_remove_gtk_header_bar(h.native(), v1)
}

// RemoveHeaderBar removes header_bar from h.
func (h *HeaderGroup) RemoveHeaderBar(headerBar *HeaderBar) {
	v1 := (*C.HdyHeaderBar)(unsafe.Pointer(headerBar.Widget.Native()))
	C.hdy_header_group_remove_header_bar(h.native(), v1)
}

// RemoveHeaderGroup removes a nested HeaderGroup from a HeaderGroup
func (h *HeaderGroup) RemoveHeaderGroup(headerGroup *HeaderGroup) {
	v1 := (*C.HdyHeaderGroup)(unsafe.Pointer(headerGroup.Native()))
	C.hdy_header_group_remove_header_group(h.native(), v1)
}

// SetDecorateAll sets whether the elements of the group should all receive the
// full decoration.
func (h *HeaderGroup) SetDecorateAll(decorateAll bool) {
	v1 := cbool(decorateAll)
	C.hdy_header_group_set_decorate_all(h.native(), v1)
}

type HeaderGroupChild struct {
	*glib.Object
}

// wrapHeaderGroupChild wraps the given pointer to *HeaderGroupChild.
func wrapHeaderGroupChild(ptr unsafe.Pointer) *HeaderGroupChild {
	obj := glib.Take(ptr)
	return &HeaderGroupChild{
		Object: obj,
	}
}

func marshalHeaderGroupChild(p uintptr) (interface{}, error) {
	return wrapHeaderGroupChild(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// native turns the current *HeaderGroupChild into the native C pointer type.
func (h *HeaderGroupChild) native() *C.HdyHeaderGroupChild {
	return (*C.HdyHeaderGroupChild)(unsafe.Pointer(h.Native()))
}

// GetChildType gets the child type.
func (h *HeaderGroupChild) GetChildType() HeaderGroupChildType {
	r := HeaderGroupChildType(C.hdy_header_group_child_get_child_type(h.native()))
	return r
}

// GetGtkHeaderBar gets the child HeaderBar. Use
// (*HeaderGroupChild).GetChildType() to check the child type.
func (h *HeaderGroupChild) GetGtkHeaderBar() *gtk.HeaderBar {
	obj := glib.Take(unsafe.Pointer(C.hdy_header_group_child_get_gtk_header_bar(h.native())))
	r := &gtk.HeaderBar{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
	return r
}

// GetHeaderBar gets the child HeaderBar. Use (*HeaderGroupChild).GetChildType()
// to check the child type.
func (h *HeaderGroupChild) GetHeaderBar() *HeaderBar {
	r := wrapHeaderBar(unsafe.Pointer(C.hdy_header_group_child_get_header_bar(h.native())))
	return r
}

// GetHeaderGroup gets the child HeaderGroup. Use
// (*HeaderGroupChild).GetChildType() to check the child type.
func (h *HeaderGroupChild) GetHeaderGroup() *HeaderGroup {
	r := wrapHeaderGroup(unsafe.Pointer(C.hdy_header_group_child_get_header_group(h.native())))
	return r
}

type Keypad struct {
	gtk.Bin
}

// wrapKeypad wraps the given pointer to *Keypad.
func wrapKeypad(ptr unsafe.Pointer) *Keypad {
	obj := glib.Take(ptr)
	return &Keypad{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalKeypad(p uintptr) (interface{}, error) {
	return wrapKeypad(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// KeypadNew create a new Keypad widget.
func KeypadNew(symbolsVisible bool, lettersVisible bool) *Keypad {
	v1 := cbool(symbolsVisible)
	v2 := cbool(lettersVisible)

	return wrapKeypad(unsafe.Pointer(C.hdy_keypad_new(v1, v2)))
}

// native turns the current *Keypad into the native C pointer type.
func (k *Keypad) native() *C.HdyKeypad {
	return (*C.HdyKeypad)(gwidget(&k.Bin))
}

// GetColumnSpacing returns the amount of space between the columns of k.
func (k *Keypad) GetColumnSpacing() uint {
	r := uint(C.hdy_keypad_get_column_spacing(k.native()))
	return r
}

// GetEndAction returns the widget for the lower right corner (or left, in RTL
// locales) of k.
func (k *Keypad) GetEndAction() gtk.IWidget {
	r, err := castWidget(C.hdy_keypad_get_end_action(k.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetEntry get the connected entry. See (*Keypad).SetEntry() for details.
func (k *Keypad) GetEntry() *gtk.Entry {
	obj := glib.Take(unsafe.Pointer(C.hdy_keypad_get_entry(k.native())))
	r := &gtk.Entry{
		Widget: gtk.Widget{
			InitiallyUnowned: glib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
	return r
}

// GetLettersVisible returns whether k should display the standard letters below
// the digits on its buttons.
func (k *Keypad) GetLettersVisible() bool {
	r := gobool(C.hdy_keypad_get_letters_visible(k.native()))
	return r
}

// GetRowSpacing returns the amount of space between the rows of k.
func (k *Keypad) GetRowSpacing() uint {
	r := uint(C.hdy_keypad_get_row_spacing(k.native()))
	return r
}

// GetStartAction returns the widget for the lower left corner (or right, in RTL
// locales) of k.
func (k *Keypad) GetStartAction() gtk.IWidget {
	r, err := castWidget(C.hdy_keypad_get_start_action(k.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetSymbolsVisible returns whether k should display the standard letters below
// the digits on its buttons.
//
// Returns Whether k should display the hash and asterisk buttons, and should
// display the plus symbol at the bottom of its 0 button.
func (k *Keypad) GetSymbolsVisible() bool {
	r := gobool(C.hdy_keypad_get_symbols_visible(k.native()))
	return r
}

// SetColumnSpacing sets the amount of space between columns of k.
func (k *Keypad) SetColumnSpacing(spacing uint) {
	v1 := C.guint(spacing)
	C.hdy_keypad_set_column_spacing(k.native(), v1)
}

// SetEndAction sets the widget for the lower right corner (or left, in RTL
// locales) of k.
func (k *Keypad) SetEndAction(endAction gtk.IWidget) {
	v1 := cwidget(endAction)
	C.hdy_keypad_set_end_action(k.native(), v1)
}

// SetEntry binds entry to k and blocks any input which wouldn't be possible to
// type with with the keypad.
func (k *Keypad) SetEntry(entry *gtk.Entry) {
	v1 := (*C.GtkEntry)(unsafe.Pointer(entry.Widget.Native()))
	C.hdy_keypad_set_entry(k.native(), v1)
}

// SetLettersVisible sets whether k should display the standard letters below
// the digits on its buttons.
func (k *Keypad) SetLettersVisible(lettersVisible bool) {
	v1 := cbool(lettersVisible)
	C.hdy_keypad_set_letters_visible(k.native(), v1)
}

// SetRowSpacing sets the amount of space between rows of k.
func (k *Keypad) SetRowSpacing(spacing uint) {
	v1 := C.guint(spacing)
	C.hdy_keypad_set_row_spacing(k.native(), v1)
}

// SetStartAction sets the widget for the lower left corner (or right, in RTL
// locales) of k.
func (k *Keypad) SetStartAction(startAction gtk.IWidget) {
	v1 := cwidget(startAction)
	C.hdy_keypad_set_start_action(k.native(), v1)
}

// SetSymbolsVisible sets whether k should display the hash and asterisk
// buttons, and should display the plus symbol at the bottom of its 0 button.
func (k *Keypad) SetSymbolsVisible(symbolsVisible bool) {
	v1 := cbool(symbolsVisible)
	C.hdy_keypad_set_symbols_visible(k.native(), v1)
}

type Leaflet struct {
	gtk.Container

	// Interfaces
	gtk.Orientable
	Swiper
}

// wrapLeaflet wraps the given pointer to *Leaflet.
func wrapLeaflet(ptr unsafe.Pointer) *Leaflet {
	obj := glib.Take(ptr)
	return &Leaflet{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
		Swiper: &Swipeable{
			Caster: &gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalLeaflet(p uintptr) (interface{}, error) {
	return wrapLeaflet(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// LeafletNew creates a new Leaflet.
func LeafletNew() *Leaflet {
	return wrapLeaflet(unsafe.Pointer(C.hdy_leaflet_new()))
}

// native turns the current *Leaflet into the native C pointer type.
func (l *Leaflet) native() *C.HdyLeaflet {
	return (*C.HdyLeaflet)(gwidget(&l.Container))
}

// GetAdjacentChild gets the previous or next child that doesn't have
// 'navigatable' child property set to false, or nil if it doesn't exist. This
// will be the same widget (*Leaflet).Navigate() will navigate to.
func (l *Leaflet) GetAdjacentChild(direction NavigationDirection) gtk.IWidget {
	v1 := C.HdyNavigationDirection(direction)
	r, err := castWidget(C.hdy_leaflet_get_adjacent_child(l.native(), v1))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetCanSwipeBack returns whether the Leaflet allows swiping to the previous
// child.
func (l *Leaflet) GetCanSwipeBack() bool {
	r := gobool(C.hdy_leaflet_get_can_swipe_back(l.native()))
	return r
}

// GetCanSwipeForward returns whether the Leaflet allows swiping to the next
// child.
func (l *Leaflet) GetCanSwipeForward() bool {
	r := gobool(C.hdy_leaflet_get_can_swipe_forward(l.native()))
	return r
}

// GetChildByName finds the child of l with the name given as the argument.
// Returns nil if there is no child with this name.
func (l *Leaflet) GetChildByName(name string) gtk.IWidget {
	v1 := C.CString(name)
	defer C.free(unsafe.Pointer(v1))
	r, err := castWidget(C.hdy_leaflet_get_child_by_name(l.native(), v1))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetChildTransitionDuration returns the amount of time (in milliseconds) that
// transitions between children in l will take.
func (l *Leaflet) GetChildTransitionDuration() uint {
	r := uint(C.hdy_leaflet_get_child_transition_duration(l.native()))
	return r
}

// GetChildTransitionRunning returns whether l is currently in a transition from
// one page to another.
func (l *Leaflet) GetChildTransitionRunning() bool {
	r := gobool(C.hdy_leaflet_get_child_transition_running(l.native()))
	return r
}

// GetFolded gets whether l is folded.
func (l *Leaflet) GetFolded() bool {
	r := gobool(C.hdy_leaflet_get_folded(l.native()))
	return r
}

// GetHomogeneous gets whether l is homogeneous for the given fold and
// orientation. See (*Leaflet).SetHomogeneous().
func (l *Leaflet) GetHomogeneous(folded bool, orientation gtk.Orientation) bool {
	v1 := cbool(folded)
	v2 := C.GtkOrientation(orientation)

	r := gobool(C.hdy_leaflet_get_homogeneous(l.native(), v1, v2))
	return r
}

// GetInterpolateSize returns whether the Leaflet is set up to interpolate
// between the sizes of children on page switch.
func (l *Leaflet) GetInterpolateSize() bool {
	r := gobool(C.hdy_leaflet_get_interpolate_size(l.native()))
	return r
}

// GetModeTransitionDuration returns the amount of time (in milliseconds) that
// transitions between modes in l will take.
func (l *Leaflet) GetModeTransitionDuration() uint {
	r := uint(C.hdy_leaflet_get_mode_transition_duration(l.native()))
	return r
}

// GetTransitionType gets the type of animation that will be used for
// transitions between modes and children in l.
func (l *Leaflet) GetTransitionType() LeafletTransitionType {
	r := LeafletTransitionType(C.hdy_leaflet_get_transition_type(l.native()))
	return r
}

// GetVisibleChild gets the visible child widget.
func (l *Leaflet) GetVisibleChild() gtk.IWidget {
	r, err := castWidget(C.hdy_leaflet_get_visible_child(l.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetVisibleChildName gets the name of the currently visible child widget.
func (l *Leaflet) GetVisibleChildName() string {
	r := C.GoString(C.hdy_leaflet_get_visible_child_name(l.native()))
	return r
}

// InsertChildAfter inserts child in the position after sibling in the list of
// children. If sibling is nil, insert child at the first position.
func (l *Leaflet) InsertChildAfter(child gtk.IWidget, sibling gtk.IWidget) {
	v1 := cwidget(child)
	v2 := cwidget(sibling)

	C.hdy_leaflet_insert_child_after(l.native(), v1, v2)
}

// Navigate switches to the previous or next child that doesn't have
// 'navigatable' child property set to false, similar to performing a swipe
// gesture to go in direction.
func (l *Leaflet) Navigate(direction NavigationDirection) bool {
	v1 := C.HdyNavigationDirection(direction)
	r := gobool(C.hdy_leaflet_navigate(l.native(), v1))
	return r
}

// Prepend inserts child at the first position in l.
func (l *Leaflet) Prepend(child gtk.IWidget) {
	v1 := cwidget(child)
	C.hdy_leaflet_prepend(l.native(), v1)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// children. If sibling is nil, move child to the first position.
func (l *Leaflet) ReorderChildAfter(child gtk.IWidget, sibling gtk.IWidget) {
	v1 := cwidget(child)
	v2 := cwidget(sibling)

	C.hdy_leaflet_reorder_child_after(l.native(), v1, v2)
}

// SetCanSwipeBack sets whether or not l allows switching to the previous child
// that has 'navigatable' child property set to true via a swipe gesture
func (l *Leaflet) SetCanSwipeBack(canSwipeBack bool) {
	v1 := cbool(canSwipeBack)
	C.hdy_leaflet_set_can_swipe_back(l.native(), v1)
}

// SetCanSwipeForward sets whether or not l allows switching to the next child
// that has 'navigatable' child property set to true via a swipe gesture.
func (l *Leaflet) SetCanSwipeForward(canSwipeForward bool) {
	v1 := cbool(canSwipeForward)
	C.hdy_leaflet_set_can_swipe_forward(l.native(), v1)
}

// SetChildTransitionDuration sets the duration that transitions between
// children in l will take.
func (l *Leaflet) SetChildTransitionDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_leaflet_set_child_transition_duration(l.native(), v1)
}

// SetHomogeneous sets the Leaflet to be homogeneous or not for the given fold
// and orientation. If it is homogeneous, the Leaflet will request the same
// width or height for all its children depending on the orientation. If it
// isn't and it is folded, the leaflet may change width or height when a
// different child becomes visible.
func (l *Leaflet) SetHomogeneous(folded bool, orientation gtk.Orientation, homogeneous bool) {
	v1 := cbool(folded)
	v2 := C.GtkOrientation(orientation)
	v3 := cbool(homogeneous)

	C.hdy_leaflet_set_homogeneous(l.native(), v1, v2, v3)
}

// SetInterpolateSize sets whether or not l will interpolate its size when
// changing the visible child. If the Leaflet:interpolate-size property is set
// to true, l will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set transition
// duration.
func (l *Leaflet) SetInterpolateSize(interpolateSize bool) {
	v1 := cbool(interpolateSize)
	C.hdy_leaflet_set_interpolate_size(l.native(), v1)
}

// SetModeTransitionDuration sets the duration that transitions between modes in
// l will take.
func (l *Leaflet) SetModeTransitionDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_leaflet_set_mode_transition_duration(l.native(), v1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between modes and children in l.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the mode or child that is about to
// become current.
func (l *Leaflet) SetTransitionType(transition LeafletTransitionType) {
	v1 := C.HdyLeafletTransitionType(transition)
	C.hdy_leaflet_set_transition_type(l.native(), v1)
}

// SetVisibleChild makes visible_child visible using a transition determined by
// HdyLeaflet:transition-type and HdyLeaflet:child-transition-duration. The
// transition can be cancelled by the user, in which case visible child will
// change back to the previously visible child.
func (l *Leaflet) SetVisibleChild(visibleChild gtk.IWidget) {
	v1 := cwidget(visibleChild)
	C.hdy_leaflet_set_visible_child(l.native(), v1)
}

// SetVisibleChildName makes the child with the name name visible.
//
// See (*Leaflet).SetVisibleChild() for more details.
func (l *Leaflet) SetVisibleChildName(name string) {
	v1 := C.CString(name)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_leaflet_set_visible_child_name(l.native(), v1)
}

type PreferencesGroup struct {
	gtk.Bin
}

// wrapPreferencesGroup wraps the given pointer to *PreferencesGroup.
func wrapPreferencesGroup(ptr unsafe.Pointer) *PreferencesGroup {
	obj := glib.Take(ptr)
	return &PreferencesGroup{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPreferencesGroup(p uintptr) (interface{}, error) {
	return wrapPreferencesGroup(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// PreferencesGroupNew creates a new PreferencesGroup.
func PreferencesGroupNew() *PreferencesGroup {
	return wrapPreferencesGroup(unsafe.Pointer(C.hdy_preferences_group_new()))
}

// native turns the current *PreferencesGroup into the native C pointer type.
func (p *PreferencesGroup) native() *C.HdyPreferencesGroup {
	return (*C.HdyPreferencesGroup)(gwidget(&p.Bin))
}

func (p *PreferencesGroup) GetDescription() string {
	r := C.GoString(C.hdy_preferences_group_get_description(p.native()))
	return r
}

// GetTitle gets the title of p.
func (p *PreferencesGroup) GetTitle() string {
	r := C.GoString(C.hdy_preferences_group_get_title(p.native()))
	return r
}

// SetDescription sets the description for p.
func (p *PreferencesGroup) SetDescription(description string) {
	v1 := C.CString(description)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_preferences_group_set_description(p.native(), v1)
}

// SetTitle sets the title for p.
func (p *PreferencesGroup) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_preferences_group_set_title(p.native(), v1)
}

type PreferencesPage struct {
	gtk.Bin
}

// wrapPreferencesPage wraps the given pointer to *PreferencesPage.
func wrapPreferencesPage(ptr unsafe.Pointer) *PreferencesPage {
	obj := glib.Take(ptr)
	return &PreferencesPage{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPreferencesPage(p uintptr) (interface{}, error) {
	return wrapPreferencesPage(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// PreferencesPageNew creates a new PreferencesPage.
func PreferencesPageNew() *PreferencesPage {
	return wrapPreferencesPage(unsafe.Pointer(C.hdy_preferences_page_new()))
}

// native turns the current *PreferencesPage into the native C pointer type.
func (p *PreferencesPage) native() *C.HdyPreferencesPage {
	return (*C.HdyPreferencesPage)(gwidget(&p.Bin))
}

// GetIconName gets the icon name for p, or nil.
func (p *PreferencesPage) GetIconName() string {
	r := C.GoString(C.hdy_preferences_page_get_icon_name(p.native()))
	return r
}

// GetTitle gets the title of p, or nil.
func (p *PreferencesPage) GetTitle() string {
	r := C.GoString(C.hdy_preferences_page_get_title(p.native()))
	return r
}

// SetIconName sets the icon name for p.
func (p *PreferencesPage) SetIconName(iconName string) {
	v1 := C.CString(iconName)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_preferences_page_set_icon_name(p.native(), v1)
}

// SetTitle sets the title of p.
func (p *PreferencesPage) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_preferences_page_set_title(p.native(), v1)
}

type PreferencesRow struct {
	gtk.ListBoxRow

	// Interfaces
	gtk.Actionable
}

// wrapPreferencesRow wraps the given pointer to *PreferencesRow.
func wrapPreferencesRow(ptr unsafe.Pointer) *PreferencesRow {
	obj := glib.Take(ptr)
	return &PreferencesRow{
		ListBoxRow: gtk.ListBoxRow{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: glib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
		Actionable: gtk.Actionable{obj},
	}
}

func marshalPreferencesRow(p uintptr) (interface{}, error) {
	return wrapPreferencesRow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// PreferencesRowNew creates a new PreferencesRow.
func PreferencesRowNew() *PreferencesRow {
	return wrapPreferencesRow(unsafe.Pointer(C.hdy_preferences_row_new()))
}

// native turns the current *PreferencesRow into the native C pointer type.
func (p *PreferencesRow) native() *C.HdyPreferencesRow {
	return (*C.HdyPreferencesRow)(gwidget(&p.ListBoxRow))
}

// GetTitle gets the title of the preference represented by p.
func (p *PreferencesRow) GetTitle() string {
	r := C.GoString(C.hdy_preferences_row_get_title(p.native()))
	return r
}

// GetUseUnderline gets whether an embedded underline in the text of the title
// indicates a mnemonic. See (*PreferencesRow).SetUseUnderline().
func (p *PreferencesRow) GetUseUnderline() bool {
	r := gobool(C.hdy_preferences_row_get_use_underline(p.native()))
	return r
}

// SetTitle sets the title of the preference represented by p.
func (p *PreferencesRow) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_preferences_row_set_title(p.native(), v1)
}

// SetUseUnderline if true, an underline in the text of the title indicates the
// next character should be used for the mnemonic accelerator key.
func (p *PreferencesRow) SetUseUnderline(useUnderline bool) {
	v1 := cbool(useUnderline)
	C.hdy_preferences_row_set_use_underline(p.native(), v1)
}

type PreferencesWindow struct {
	Window
}

// wrapPreferencesWindow wraps the given pointer to *PreferencesWindow.
func wrapPreferencesWindow(ptr unsafe.Pointer) *PreferencesWindow {
	obj := glib.Take(ptr)
	return &PreferencesWindow{
		Window: Window{
			Window: gtk.Window{
				Bin: gtk.Bin{
					Container: gtk.Container{
						Widget: gtk.Widget{
							InitiallyUnowned: glib.InitiallyUnowned{
								Object: obj,
							},
						},
					},
				},
			},
		},
	}
}

func marshalPreferencesWindow(p uintptr) (interface{}, error) {
	return wrapPreferencesWindow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// PreferencesWindowNew creates a new PreferencesWindow.
func PreferencesWindowNew() *PreferencesWindow {
	return wrapPreferencesWindow(unsafe.Pointer(C.hdy_preferences_window_new()))
}

// native turns the current *PreferencesWindow into the native C pointer type.
func (p *PreferencesWindow) native() *C.HdyPreferencesWindow {
	return (*C.HdyPreferencesWindow)(gwidget(&p.Window))
}

// CloseSubpage closes the current subpage to return back to the preferences, if
// there is no presented subpage, this does nothing.
func (p *PreferencesWindow) CloseSubpage() {
	C.hdy_preferences_window_close_subpage(p.native())
}

// GetCanSwipeBack returns whether or not p allows switching from a subpage to
// the preferences via a swipe gesture.
func (p *PreferencesWindow) GetCanSwipeBack() bool {
	r := gobool(C.hdy_preferences_window_get_can_swipe_back(p.native()))
	return r
}

// GetSearchEnabled gets whether search is enabled for p.
func (p *PreferencesWindow) GetSearchEnabled() bool {
	r := gobool(C.hdy_preferences_window_get_search_enabled(p.native()))
	return r
}

// PresentSubpage sets subpage as the window's subpage and present it. The
// transition can be cancelled by the user, in which case visible child will
// change back to the previously visible child.
func (p *PreferencesWindow) PresentSubpage(subpage gtk.IWidget) {
	v1 := cwidget(subpage)
	C.hdy_preferences_window_present_subpage(p.native(), v1)
}

// SetCanSwipeBack sets whether or not p allows switching from a subpage to the
// preferences via a swipe gesture.
func (p *PreferencesWindow) SetCanSwipeBack(canSwipeBack bool) {
	v1 := cbool(canSwipeBack)
	C.hdy_preferences_window_set_can_swipe_back(p.native(), v1)
}

// SetSearchEnabled sets whether search is enabled for p.
func (p *PreferencesWindow) SetSearchEnabled(searchEnabled bool) {
	v1 := cbool(searchEnabled)
	C.hdy_preferences_window_set_search_enabled(p.native(), v1)
}

type SearchBar struct {
	gtk.Bin
}

// wrapSearchBar wraps the given pointer to *SearchBar.
func wrapSearchBar(ptr unsafe.Pointer) *SearchBar {
	obj := glib.Take(ptr)
	return &SearchBar{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	return wrapSearchBar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// SearchBarNew creates a SearchBar. You will need to tell it about which widget
// is going to be your text entry using (*SearchBar).ConnectEntry().
func SearchBarNew() *SearchBar {
	return wrapSearchBar(unsafe.Pointer(C.hdy_search_bar_new()))
}

// native turns the current *SearchBar into the native C pointer type.
func (s *SearchBar) native() *C.HdySearchBar {
	return (*C.HdySearchBar)(gwidget(&s.Bin))
}

// ConnectEntry connects the Entry widget passed as the one to be used in this
// search bar. The entry should be a descendant of the search bar. This is only
// required if the entry isn’t the direct child of the search bar (as in our
// main example).
func (s *SearchBar) ConnectEntry(entry *gtk.Entry) {
	v1 := (*C.GtkEntry)(unsafe.Pointer(entry.Widget.Native()))
	C.hdy_search_bar_connect_entry(s.native(), v1)
}

// GetSearchMode returns whether the search mode is on or off.
func (s *SearchBar) GetSearchMode() bool {
	r := gobool(C.hdy_search_bar_get_search_mode(s.native()))
	return r
}

// GetShowCloseButton returns whether the close button is shown.
func (s *SearchBar) GetShowCloseButton() bool {
	r := gobool(C.hdy_search_bar_get_show_close_button(s.native()))
	return r
}

// HandleEvent function should be called when the top-level window which
// contains the search bar received a key event.
//
// If the key event is handled by the search bar, the bar will be shown, the
// entry populated with the entered text and GDK_EVENT_STOP will be returned.
// The caller should ensure that events are not propagated further.
//
// If no entry has been connected to the search bar, using
// (*SearchBar).ConnectEntry(), this function will return immediately with a
// warning.
//
// Showing the search bar on key presses
//
//    static gboolean
//    on_key_press_event (GtkWidget *widget,
//                        GdkEvent  *event,
//                        gpointer   user_data)
//    {
//      HdySearchBar *bar = HDY_SEARCH_BAR (user_data);
//      return hdy_search_bar_handle_event (self, event);
//    }
//
//    static void
//    create_toplevel (void)
//    {
//      GtkWidget *window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      GtkWindow *search_bar = hdy_search_bar_new ();
//
//     // Add more widgets to the window...
//
//      g_signal_connect (window,
//                       "key-press-event",
//                        G_CALLBACK (on_key_press_event),
//                        search_bar);
//    }
//
func (s *SearchBar) HandleEvent(event *gdk.Event) bool {
	v1 := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	r := gobool(C.hdy_search_bar_handle_event(s.native(), v1))
	return r
}

// SetSearchMode switches the search mode on or off.
func (s *SearchBar) SetSearchMode(searchMode bool) {
	v1 := cbool(searchMode)
	C.hdy_search_bar_set_search_mode(s.native(), v1)
}

// SetShowCloseButton shows or hides the close button. Applications that already
// have a “search” toggle button should not show a close button in their
// search bar, as it duplicates the role of the toggle button.
func (s *SearchBar) SetShowCloseButton(visible bool) {
	v1 := cbool(visible)
	C.hdy_search_bar_set_show_close_button(s.native(), v1)
}

type Squeezer struct {
	gtk.Container

	// Interfaces
	gtk.Orientable
}

// wrapSqueezer wraps the given pointer to *Squeezer.
func wrapSqueezer(ptr unsafe.Pointer) *Squeezer {
	obj := glib.Take(ptr)
	return &Squeezer{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
		Orientable: gtk.Orientable{obj},
	}
}

func marshalSqueezer(p uintptr) (interface{}, error) {
	return wrapSqueezer(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// SqueezerNew creates a new Squeezer container.
func SqueezerNew() *Squeezer {
	return wrapSqueezer(unsafe.Pointer(C.hdy_squeezer_new()))
}

// native turns the current *Squeezer into the native C pointer type.
func (s *Squeezer) native() *C.HdySqueezer {
	return (*C.HdySqueezer)(gwidget(&s.Container))
}

// GetChildEnabled gets whether child is enabled.
//
// See (*Squeezer).SetChildEnabled().
func (s *Squeezer) GetChildEnabled(child gtk.IWidget) bool {
	v1 := cwidget(child)
	r := gobool(C.hdy_squeezer_get_child_enabled(s.native(), v1))
	return r
}

// GetHomogeneous gets whether s is homogeneous.
//
// See (*Squeezer).SetHomogeneous().
func (s *Squeezer) GetHomogeneous() bool {
	r := gobool(C.hdy_squeezer_get_homogeneous(s.native()))
	return r
}

// GetInterpolateSize gets whether s should interpolate its size on visible
// child change.
//
// See (*Squeezer).SetInterpolateSize().
func (s *Squeezer) GetInterpolateSize() bool {
	r := gobool(C.hdy_squeezer_get_interpolate_size(s.native()))
	return r
}

// GetTransitionDuration gets the amount of time (in milliseconds) that
// transitions between children in s will take.
func (s *Squeezer) GetTransitionDuration() uint {
	r := uint(C.hdy_squeezer_get_transition_duration(s.native()))
	return r
}

// GetTransitionRunning gets whether s is currently in a transition from one
// child to another.
func (s *Squeezer) GetTransitionRunning() bool {
	r := gobool(C.hdy_squeezer_get_transition_running(s.native()))
	return r
}

// GetTransitionType gets the type of animation that will be used for
// transitions between children in s.
func (s *Squeezer) GetTransitionType() SqueezerTransitionType {
	r := SqueezerTransitionType(C.hdy_squeezer_get_transition_type(s.native()))
	return r
}

// GetVisibleChild gets the currently visible child of s, or nil if there are no
// visible children.
func (s *Squeezer) GetVisibleChild() gtk.IWidget {
	r, err := castWidget(C.hdy_squeezer_get_visible_child(s.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetXAlign gets the Squeezer:xalign property for s.
func (s *Squeezer) GetXAlign() float32 {
	r := float32(C.hdy_squeezer_get_xalign(s.native()))
	return r
}

// GetYAlign gets the Squeezer:yalign property for s.
func (s *Squeezer) GetYAlign() float32 {
	r := float32(C.hdy_squeezer_get_yalign(s.native()))
	return r
}

// SetChildEnabled make s enable or disable child. If a child is disabled, it
// will be ignored when looking for the child fitting the available size best.
// This allows to programmatically and prematurely hide a child of s even if it
// fits in the available space.
//
// This can be used e.g. to ensure a certain child is hidden below a certain
// window width, or any other constraint you find suitable.
func (s *Squeezer) SetChildEnabled(child gtk.IWidget, enabled bool) {
	v1 := cwidget(child)
	v2 := cbool(enabled)

	C.hdy_squeezer_set_child_enabled(s.native(), v1, v2)
}

// SetHomogeneous sets s to be homogeneous or not. If it is homogeneous, s will
// request the same size for all its children for its opposite orientation, e.g.
// if s is oriented horizontally and is homogeneous, it will request the same
// height for all its children. If it isn't, s may change size when a different
// child becomes visible.
func (s *Squeezer) SetHomogeneous(homogeneous bool) {
	v1 := cbool(homogeneous)
	C.hdy_squeezer_set_homogeneous(s.native(), v1)
}

// SetInterpolateSize sets whether or not s will interpolate the size of its
// opposing orientation when changing the visible child. If true, s will
// interpolate its size between the one of the previous visible child and the
// one of the new visible child, according to the set transition duration and
// the orientation, e.g. if s is horizontal, it will interpolate the its height.
func (s *Squeezer) SetInterpolateSize(interpolateSize bool) {
	v1 := cbool(interpolateSize)
	C.hdy_squeezer_set_interpolate_size(s.native(), v1)
}

// SetTransitionDuration sets the duration that transitions between children in
// s will take.
func (s *Squeezer) SetTransitionDuration(duration uint) {
	v1 := C.guint(duration)
	C.hdy_squeezer_set_transition_duration(s.native(), v1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between children in s. Available types include various kinds of
// fades and slides.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the child that is about to become
// current.
func (s *Squeezer) SetTransitionType(transition SqueezerTransitionType) {
	v1 := C.HdySqueezerTransitionType(transition)
	C.hdy_squeezer_set_transition_type(s.native(), v1)
}

// SetXAlign sets the Squeezer:xalign property for s.
func (s *Squeezer) SetXAlign(xalign float32) {
	v1 := C.gfloat(xalign)
	C.hdy_squeezer_set_xalign(s.native(), v1)
}

// SetYAlign sets the Squeezer:yalign property for s.
func (s *Squeezer) SetYAlign(yalign float32) {
	v1 := C.gfloat(yalign)
	C.hdy_squeezer_set_yalign(s.native(), v1)
}

type StatusPage struct {
	gtk.Bin
}

// wrapStatusPage wraps the given pointer to *StatusPage.
func wrapStatusPage(ptr unsafe.Pointer) *StatusPage {
	obj := glib.Take(ptr)
	return &StatusPage{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalStatusPage(p uintptr) (interface{}, error) {
	return wrapStatusPage(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// StatusPageNew creates a new StatusPage.
func StatusPageNew() *StatusPage {
	return wrapStatusPage(unsafe.Pointer(C.hdy_status_page_new()))
}

// native turns the current *StatusPage into the native C pointer type.
func (s *StatusPage) native() *C.HdyStatusPage {
	return (*C.HdyStatusPage)(gwidget(&s.Bin))
}

// GetDescription gets the description for s.
func (s *StatusPage) GetDescription() string {
	r := C.GoString(C.hdy_status_page_get_description(s.native()))
	return r
}

// GetIconName gets the icon name for s.
func (s *StatusPage) GetIconName() string {
	r := C.GoString(C.hdy_status_page_get_icon_name(s.native()))
	return r
}

// GetTitle gets the title for s.
func (s *StatusPage) GetTitle() string {
	r := C.GoString(C.hdy_status_page_get_title(s.native()))
	return r
}

// SetDescription sets the description for s.
func (s *StatusPage) SetDescription(description string) {
	v1 := C.CString(description)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_status_page_set_description(s.native(), v1)
}

// SetIconName sets the icon name for s.
func (s *StatusPage) SetIconName(iconName string) {
	v1 := C.CString(iconName)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_status_page_set_icon_name(s.native(), v1)
}

// SetTitle sets the title for s.
func (s *StatusPage) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_status_page_set_title(s.native(), v1)
}

type SwipeGroup struct {
	*glib.Object
}

// wrapSwipeGroup wraps the given pointer to *SwipeGroup.
func wrapSwipeGroup(ptr unsafe.Pointer) *SwipeGroup {
	obj := glib.Take(ptr)
	return &SwipeGroup{
		Object: obj,
	}
}

func marshalSwipeGroup(p uintptr) (interface{}, error) {
	return wrapSwipeGroup(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// SwipeGroupNew create a new SwipeGroup object.
func SwipeGroupNew() *SwipeGroup {
	return wrapSwipeGroup(unsafe.Pointer(C.hdy_swipe_group_new()))
}

// native turns the current *SwipeGroup into the native C pointer type.
func (s *SwipeGroup) native() *C.HdySwipeGroup {
	return (*C.HdySwipeGroup)(unsafe.Pointer(s.Native()))
}

// AddSwipeable when the widget is destroyed or no longer referenced elsewhere,
// it will be removed from the swipe group.
func (s *SwipeGroup) AddSwipeable(swipeable Swiper) {
	v1 := (*C.HdySwipeable)(unsafe.Pointer(swipeable.Native()))
	C.hdy_swipe_group_add_swipeable(s.native(), v1)
}

// GetSwipeables returns the list of swipeables associated with s.
func (s *SwipeGroup) GetSwipeables() *glib.SList {
	r := glib.WrapSList(uintptr(unsafe.Pointer(C.hdy_swipe_group_get_swipeables(s.native()))))
	return r
}

// RemoveSwipeable removes a widget from a SwipeGroup.
func (s *SwipeGroup) RemoveSwipeable(swipeable Swiper) {
	v1 := (*C.HdySwipeable)(unsafe.Pointer(swipeable.Native()))
	C.hdy_swipe_group_remove_swipeable(s.native(), v1)
}

type SwipeTracker struct {
	*glib.Object

	// Interfaces
	gtk.Orientable
}

// wrapSwipeTracker wraps the given pointer to *SwipeTracker.
func wrapSwipeTracker(ptr unsafe.Pointer) *SwipeTracker {
	obj := glib.Take(ptr)
	return &SwipeTracker{
		Object:     obj,
		Orientable: gtk.Orientable{obj},
	}
}

func marshalSwipeTracker(p uintptr) (interface{}, error) {
	return wrapSwipeTracker(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// SwipeTrackerNew create a new SwipeTracker object on widget.
func SwipeTrackerNew(swipeable Swipeable) *SwipeTracker {
	v1 := (*C.HdySwipeable)(unsafe.Pointer(swipeable.Native()))
	return wrapSwipeTracker(unsafe.Pointer(C.hdy_swipe_tracker_new(v1)))
}

// native turns the current *SwipeTracker into the native C pointer type.
func (s *SwipeTracker) native() *C.HdySwipeTracker {
	return (*C.HdySwipeTracker)(unsafe.Pointer(s.Native()))
}

// GetAllowLongSwipes whether to allow swiping for more than one snap point at a
// time. If the value is false, each swipe can only move to the adjacent snap
// points.
func (s *SwipeTracker) GetAllowLongSwipes() bool {
	r := gobool(C.hdy_swipe_tracker_get_allow_long_swipes(s.native()))
	return r
}

// GetAllowMouseDrag get whether s can be dragged with mouse pointer.
func (s *SwipeTracker) GetAllowMouseDrag() bool {
	r := gobool(C.hdy_swipe_tracker_get_allow_mouse_drag(s.native()))
	return r
}

// GetEnabled get whether s is enabled. When it's not enabled, no events will be
// processed. Generally widgets will want to expose this via a property.
func (s *SwipeTracker) GetEnabled() bool {
	r := gobool(C.hdy_swipe_tracker_get_enabled(s.native()))
	return r
}

// GetReversed get whether s is reversing the swipe direction.
func (s *SwipeTracker) GetReversed() bool {
	r := gobool(C.hdy_swipe_tracker_get_reversed(s.native()))
	return r
}

// GetSwipeable get s's swipeable widget.
func (s *SwipeTracker) GetSwipeable() Swiper {
	obj := glib.Take(unsafe.Pointer(C.hdy_swipe_tracker_get_swipeable(s.native())))
	r := &Swipeable{
		Caster: &gtk.Widget{
			InitiallyUnowned: glib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
	return r
}

// SetAllowLongSwipes sets whether to allow swiping for more than one snap point
// at a time. If the value is false, each swipe can only move to the adjacent
// snap points.
func (s *SwipeTracker) SetAllowLongSwipes(allowLongSwipes bool) {
	v1 := cbool(allowLongSwipes)
	C.hdy_swipe_tracker_set_allow_long_swipes(s.native(), v1)
}

// SetAllowMouseDrag set whether s can be dragged with mouse pointer. This
// should usually be false.
func (s *SwipeTracker) SetAllowMouseDrag(allowMouseDrag bool) {
	v1 := cbool(allowMouseDrag)
	C.hdy_swipe_tracker_set_allow_mouse_drag(s.native(), v1)
}

// SetEnabled set whether s is enabled. When it's not enabled, no events will be
// processed. Usually widgets will want to expose this via a property.
func (s *SwipeTracker) SetEnabled(enabled bool) {
	v1 := cbool(enabled)
	C.hdy_swipe_tracker_set_enabled(s.native(), v1)
}

// SetReversed set whether to reverse the swipe direction. If s is horizontal,
// can be used for supporting RTL text direction.
func (s *SwipeTracker) SetReversed(reversed bool) {
	v1 := cbool(reversed)
	C.hdy_swipe_tracker_set_reversed(s.native(), v1)
}

// ShiftPosition move the current progress value by delta. This can be used to
// adjust the current position if snap points move during the gesture.
func (s *SwipeTracker) ShiftPosition(delta float64) {
	v1 := C.gdouble(delta)
	C.hdy_swipe_tracker_shift_position(s.native(), v1)
}

type TabBar struct {
	gtk.Bin
}

// wrapTabBar wraps the given pointer to *TabBar.
func wrapTabBar(ptr unsafe.Pointer) *TabBar {
	obj := glib.Take(ptr)
	return &TabBar{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalTabBar(p uintptr) (interface{}, error) {
	return wrapTabBar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// TabBarNew creates a new TabBar widget.
func TabBarNew() *TabBar {
	return wrapTabBar(unsafe.Pointer(C.hdy_tab_bar_new()))
}

// native turns the current *TabBar into the native C pointer type.
func (t *TabBar) native() *C.HdyTabBar {
	return (*C.HdyTabBar)(gwidget(&t.Bin))
}

// GetAutohide gets whether the tabs automatically hide, see
// (*TabBar).SetAutohide().
func (t *TabBar) GetAutohide() bool {
	r := gobool(C.hdy_tab_bar_get_autohide(t.native()))
	return r
}

// GetEndActionWidget gets the widget shown after the tabs.
func (t *TabBar) GetEndActionWidget() gtk.IWidget {
	r, err := castWidget(C.hdy_tab_bar_get_end_action_widget(t.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetExpandTabs gets whether tabs should expand, see (*TabBar).SetExpandTabs().
func (t *TabBar) GetExpandTabs() bool {
	r := gobool(C.hdy_tab_bar_get_expand_tabs(t.native()))
	return r
}

// GetInverted gets whether tabs use inverted layout, see
// (*TabBar).SetInverted().
func (t *TabBar) GetInverted() bool {
	r := gobool(C.hdy_tab_bar_get_inverted(t.native()))
	return r
}

// GetIsOverflowing gets whether t is overflowing.
func (t *TabBar) GetIsOverflowing() bool {
	r := gobool(C.hdy_tab_bar_get_is_overflowing(t.native()))
	return r
}

// GetStartActionWidget gets the widget shown before the tabs.
func (t *TabBar) GetStartActionWidget() gtk.IWidget {
	r, err := castWidget(C.hdy_tab_bar_get_start_action_widget(t.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetTabsRevealed gets the value of the TabBar:tabs-revealed property.
func (t *TabBar) GetTabsRevealed() bool {
	r := gobool(C.hdy_tab_bar_get_tabs_revealed(t.native()))
	return r
}

// GetView gets the TabView t controls.
func (t *TabBar) GetView() *TabView {
	r := wrapTabView(unsafe.Pointer(C.hdy_tab_bar_get_view(t.native())))
	return r
}

// SetAutohide sets whether the tabs automatically hide.
//
// If autohide is true, the tab bar disappears when the associated TabView has 0
// or 1 tab, no pinned tabs, and no tab is being transferred.
//
// Autohide is enabled by default.
//
// See TabBar:tabs-revealed.
func (t *TabBar) SetAutohide(autohide bool) {
	v1 := cbool(autohide)
	C.hdy_tab_bar_set_autohide(t.native(), v1)
}

// SetEndActionWidget sets the widget to show after the tabs.
func (t *TabBar) SetEndActionWidget(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_tab_bar_set_end_action_widget(t.native(), v1)
}

// SetExpandTabs sets whether tabs should expand.
//
// If expand_tabs is true, the tabs will always vary width filling the whole
// width when possible, otherwise tabs will always have the minimum possible
// size.
//
// Expand is enabled by default.
func (t *TabBar) SetExpandTabs(expandTabs bool) {
	v1 := cbool(expandTabs)
	C.hdy_tab_bar_set_expand_tabs(t.native(), v1)
}

// SetInverted sets whether tabs tabs use inverted layout.
//
// If inverted is true, non-pinned tabs will have the close button at the
// beginning and the indicator at the end rather than the opposite.
func (t *TabBar) SetInverted(inverted bool) {
	v1 := cbool(inverted)
	C.hdy_tab_bar_set_inverted(t.native(), v1)
}

// SetStartActionWidget sets the widget to show before the tabs.
func (t *TabBar) SetStartActionWidget(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_tab_bar_set_start_action_widget(t.native(), v1)
}

// SetView sets the TabView t controls.
func (t *TabBar) SetView(view *TabView) {
	v1 := (*C.HdyTabView)(unsafe.Pointer(view.Widget.Native()))
	C.hdy_tab_bar_set_view(t.native(), v1)
}

type TabPage struct {
	*glib.Object
}

// wrapTabPage wraps the given pointer to *TabPage.
func wrapTabPage(ptr unsafe.Pointer) *TabPage {
	obj := glib.Take(ptr)
	return &TabPage{
		Object: obj,
	}
}

func marshalTabPage(p uintptr) (interface{}, error) {
	return wrapTabPage(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// native turns the current *TabPage into the native C pointer type.
func (t *TabPage) native() *C.HdyTabPage {
	return (*C.HdyTabPage)(unsafe.Pointer(t.Native()))
}

// GetChild gets the child of t.
func (t *TabPage) GetChild() gtk.IWidget {
	r, err := castWidget(C.hdy_tab_page_get_child(t.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// GetIcon gets the icon of t, see (*TabPage).SetIcon().
func (t *TabPage) GetIcon() *glib.Icon {
	obj := glib.Take(unsafe.Pointer(C.hdy_tab_page_get_icon(t.native())))
	r := &glib.Icon{
		Object: obj,
	}
	return r
}

// GetIndicatorActivatable gets whether the indicator of t is activatable, see
// (*TabPage).SetIndicatorActivatable().
func (t *TabPage) GetIndicatorActivatable() bool {
	r := gobool(C.hdy_tab_page_get_indicator_activatable(t.native()))
	return r
}

// GetIndicatorIcon gets the indicator icon of t, see
// (*TabPage).SetIndicatorIcon().
func (t *TabPage) GetIndicatorIcon() *glib.Icon {
	obj := glib.Take(unsafe.Pointer(C.hdy_tab_page_get_indicator_icon(t.native())))
	r := &glib.Icon{
		Object: obj,
	}
	return r
}

// GetLoading gets whether t is loading, see (*TabPage).SetLoading().
func (t *TabPage) GetLoading() bool {
	r := gobool(C.hdy_tab_page_get_loading(t.native()))
	return r
}

// GetNeedsAttention gets whether t needs attention, see
// (*TabPage).SetNeedsAttention().
func (t *TabPage) GetNeedsAttention() bool {
	r := gobool(C.hdy_tab_page_get_needs_attention(t.native()))
	return r
}

// GetParent gets the parent page of t, or nil if the t does not have a parent.
//
// See (*TabView).AddPage() and (*TabView).ClosePage().
func (t *TabPage) GetParent() *TabPage {
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_page_get_parent(t.native())))
	return r
}

// GetPinned gets whether t is pinned. See (*TabView).SetPagePinned().
func (t *TabPage) GetPinned() bool {
	r := gobool(C.hdy_tab_page_get_pinned(t.native()))
	return r
}

// GetSelected gets whether t is selected. See (*TabView).SetSelectedPage().
func (t *TabPage) GetSelected() bool {
	r := gobool(C.hdy_tab_page_get_selected(t.native()))
	return r
}

// GetTitle gets the title of t, see (*TabPage).SetTitle().
func (t *TabPage) GetTitle() string {
	r := C.GoString(C.hdy_tab_page_get_title(t.native()))
	return r
}

// GetTooltip gets the tooltip of t, see (*TabPage).SetTooltip().
func (t *TabPage) GetTooltip() string {
	r := C.GoString(C.hdy_tab_page_get_tooltip(t.native()))
	return r
}

// SetIcon sets the icon of t, displayed next to the title.
//
// TabBar will not show the icon if TabPage:loading is set to true, or if t is
// pinned and TabPage:indicator-icon is set.
func (t *TabPage) SetIcon(icon *glib.Icon) {
	v1 := (*C.GIcon)(unsafe.Pointer(icon.Native()))
	C.hdy_tab_page_set_icon(t.native(), v1)
}

// SetIndicatorActivatable sets whether the indicator of t is activatable.
//
// If set to true, TabView::indicator-activated will be emitted when the
// indicator is clicked.
//
// If TabPage:indicator-icon is not set, does nothing.
func (t *TabPage) SetIndicatorActivatable(activatable bool) {
	v1 := cbool(activatable)
	C.hdy_tab_page_set_indicator_activatable(t.native(), v1)
}

// SetIndicatorIcon sets the indicator icon of t.
//
// A common use case is an audio or camera indicator in a web browser.
//
// TabPage will show it at the beginning of the tab, alongside icon representing
// TabPage:icon or loading spinner.
//
// If the page is pinned, the indicator will be shown instead of icon or
// spinner.
//
// If TabPage:indicator-activatable is set to true, indicator icon can act as a
// button.
func (t *TabPage) SetIndicatorIcon(indicatorIcon *glib.Icon) {
	v1 := (*C.GIcon)(unsafe.Pointer(indicatorIcon.Native()))
	C.hdy_tab_page_set_indicator_icon(t.native(), v1)
}

// SetLoading sets wether t is loading.
//
// If set to true, TabBar will display a spinner in place of icon.
//
// If t is pinned and TabPage:indicator-icon is set, the loading status will not
// be visible.
func (t *TabPage) SetLoading(loading bool) {
	v1 := cbool(loading)
	C.hdy_tab_page_set_loading(t.native(), v1)
}

// SetNeedsAttention sets whether t needs attention.
//
// TabBar will display a glow under the tab representing t if set to true. If
// the tab is not visible, the corresponding edge of the tab bar will be
// highlighted.
func (t *TabPage) SetNeedsAttention(needsAttention bool) {
	v1 := cbool(needsAttention)
	C.hdy_tab_page_set_needs_attention(t.native(), v1)
}

// SetTitle sets the title of t.
//
// TabBar will display it in the center of the tab representing t unless it's
// pinned, and will use it as a tooltip unless TabPage:tooltip is set.
func (t *TabPage) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_tab_page_set_title(t.native(), v1)
}

// SetTooltip sets the tooltip of t, marked up with the Pango text markup
// language.
//
// If not set, TabBar will use TabPage:title as a tooltip instead.
func (t *TabPage) SetTooltip(tooltip string) {
	v1 := C.CString(tooltip)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_tab_page_set_tooltip(t.native(), v1)
}

type TabView struct {
	gtk.Bin
}

// wrapTabView wraps the given pointer to *TabView.
func wrapTabView(ptr unsafe.Pointer) *TabView {
	obj := glib.Take(ptr)
	return &TabView{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalTabView(p uintptr) (interface{}, error) {
	return wrapTabView(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// TabViewNew creates a new TabView widget.
func TabViewNew() *TabView {
	return wrapTabView(unsafe.Pointer(C.hdy_tab_view_new()))
}

// native turns the current *TabView into the native C pointer type.
func (t *TabView) native() *C.HdyTabView {
	return (*C.HdyTabView)(gwidget(&t.Bin))
}

// AddPage adds child to t with parent as the parent.
//
// This function can be used to automatically position new pages, and to select
// the correct page when this page is closed while being selected (see
// (*TabView).ClosePage()).
//
// If parent is nil, this function is equivalent to (*TabView).Append().
func (t *TabView) AddPage(child gtk.IWidget, parent *TabPage) *TabPage {
	v1 := cwidget(child)
	v2 := (*C.HdyTabPage)(unsafe.Pointer(parent.Native()))

	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_add_page(t.native(), v1, v2)))
	return r
}

// Append inserts child as the last non-pinned page.
func (t *TabView) Append(child gtk.IWidget) *TabPage {
	v1 := cwidget(child)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_append(t.native(), v1)))
	return r
}

// AppendPinned inserts child as the last pinned page.
func (t *TabView) AppendPinned(child gtk.IWidget) *TabPage {
	v1 := cwidget(child)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_append_pinned(t.native(), v1)))
	return r
}

// CloseOtherPages requests to close all pages other than page.
func (t *TabView) CloseOtherPages(page *TabPage) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	C.hdy_tab_view_close_other_pages(t.native(), v1)
}

// ClosePage requests to close page.
//
// Calling this function will result in TabView::close-page signal being emitted
// for page. Closing the page can then be confirmed or denied via
// (*TabView).ClosePageFinish().
//
// If the page is waiting for a (*TabView).ClosePageFinish() call, this function
// will do nothing.
//
// The default handler for TabView::close-page will immediately confirm closing
// the page if it's non-pinned, or reject it if it's pinned. This behavior can
// be changed by registering your own handler for that signal.
//
// If page was selected, another page will be selected instead:
//
// If the TabPage:parent value is nil, the next page will be selected when
// possible, or if the page was already last, the previous page will be selected
// instead.
//
// If it's not nil, the previous page will be selected if it's a descendant
// (possibly indirect) of the parent. If both the previous page and the parent
// are pinned, the parent will be selected instead.
func (t *TabView) ClosePage(page *TabPage) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	C.hdy_tab_view_close_page(t.native(), v1)
}

// ClosePageFinish completes a (*TabView).ClosePage() call for page.
//
// If confirm is true, page will be closed. If it's false, ite will be reverted
// to its previous state and (*TabView).ClosePage() can be called for it again.
//
// This function should not be called unless a custom handler for
// TabView::close-page is used.
func (t *TabView) ClosePageFinish(page *TabPage, confirm bool) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	v2 := cbool(confirm)

	C.hdy_tab_view_close_page_finish(t.native(), v1, v2)
}

// ClosePagesAfter requests to close all pages after page.
func (t *TabView) ClosePagesAfter(page *TabPage) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	C.hdy_tab_view_close_pages_after(t.native(), v1)
}

// ClosePagesBefore requests to close all pages before page.
func (t *TabView) ClosePagesBefore(page *TabPage) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	C.hdy_tab_view_close_pages_before(t.native(), v1)
}

// GetDefaultIcon gets default icon of t, see (*TabView).SetDefaultIcon().
func (t *TabView) GetDefaultIcon() *glib.Icon {
	obj := glib.Take(unsafe.Pointer(C.hdy_tab_view_get_default_icon(t.native())))
	r := &glib.Icon{
		Object: obj,
	}
	return r
}

// GetIsTransferringPage whether a page is being transferred.
//
// Gets the value of TabView:is-transferring-page property.
func (t *TabView) GetIsTransferringPage() bool {
	r := gobool(C.hdy_tab_view_get_is_transferring_page(t.native()))
	return r
}

// GetMenuModel gets the tab context menu model for t, see
// (*TabView).SetMenuModel().
func (t *TabView) GetMenuModel() *glib.MenuModel {
	obj := glib.Take(unsafe.Pointer(C.hdy_tab_view_get_menu_model(t.native())))
	r := &glib.MenuModel{
		Object: obj,
	}
	return r
}

// GetNPages gets the number of pages in t.
func (t *TabView) GetNPages() int {
	r := int(C.hdy_tab_view_get_n_pages(t.native()))
	return r
}

// GetNPinnedPages gets the number of pinned pages in t.
//
// See (*TabView).SetPagePinned().
func (t *TabView) GetNPinnedPages() int {
	r := int(C.hdy_tab_view_get_n_pinned_pages(t.native()))
	return r
}

// GetNthPage gets the TabPage representing the child at position.
func (t *TabView) GetNthPage(position int) *TabPage {
	v1 := C.gint(position)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_get_nth_page(t.native(), v1)))
	return r
}

// GetPage gets the TabPage object representing child.
func (t *TabView) GetPage(child gtk.IWidget) *TabPage {
	v1 := cwidget(child)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_get_page(t.native(), v1)))
	return r
}

// GetPagePosition finds the position of page in t, starting from 0.
func (t *TabView) GetPagePosition(page *TabPage) int {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	r := int(C.hdy_tab_view_get_page_position(t.native(), v1))
	return r
}

// GetPages returns a Model containing the pages of t. This model can be used to
// keep an up to date view of the pages.
func (t *TabView) GetPages() *glib.ListModel {
	obj := glib.Take(unsafe.Pointer(C.hdy_tab_view_get_pages(t.native())))
	r := &glib.ListModel{
		Object: obj,
	}
	return r
}

// GetSelectedPage gets the currently selected page in t.
func (t *TabView) GetSelectedPage() *TabPage {
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_get_selected_page(t.native())))
	return r
}

// GetShortcutWidget gets the shortcut widget for t, see
// (*TabView).SetShortcutWidget().
func (t *TabView) GetShortcutWidget() gtk.IWidget {
	r, err := castWidget(C.hdy_tab_view_get_shortcut_widget(t.native()))
	if err != nil {
		panic("cast widget *C.GtkWidget failed: " + err.Error())
	}
	return r
}

// Insert inserts a non-pinned page at position.
//
// It's an error to try to insert a page before a pinned page, in that case
// (*TabView).InsertPinned() should be used instead.
func (t *TabView) Insert(child gtk.IWidget, position int) *TabPage {
	v1 := cwidget(child)
	v2 := C.gint(position)

	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_insert(t.native(), v1, v2)))
	return r
}

// InsertPinned inserts a pinned page at position.
//
// It's an error to try to insert a pinned page after a non-pinned page, in that
// case (*TabView).Insert() should be used instead.
func (t *TabView) InsertPinned(child gtk.IWidget, position int) *TabPage {
	v1 := cwidget(child)
	v2 := C.gint(position)

	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_insert_pinned(t.native(), v1, v2)))
	return r
}

// Prepend inserts child as the first non-pinned page.
func (t *TabView) Prepend(child gtk.IWidget) *TabPage {
	v1 := cwidget(child)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_prepend(t.native(), v1)))
	return r
}

// PrependPinned inserts child as the first pinned page.
func (t *TabView) PrependPinned(child gtk.IWidget) *TabPage {
	v1 := cwidget(child)
	r := wrapTabPage(unsafe.Pointer(C.hdy_tab_view_prepend_pinned(t.native(), v1)))
	return r
}

// ReorderBackward reorders page to before its previous page if possible.
func (t *TabView) ReorderBackward(page *TabPage) bool {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	r := gobool(C.hdy_tab_view_reorder_backward(t.native(), v1))
	return r
}

// ReorderFirst reorders page to the first possible position.
func (t *TabView) ReorderFirst(page *TabPage) bool {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	r := gobool(C.hdy_tab_view_reorder_first(t.native(), v1))
	return r
}

// ReorderForward reorders page to after its next page if possible.
func (t *TabView) ReorderForward(page *TabPage) bool {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	r := gobool(C.hdy_tab_view_reorder_forward(t.native(), v1))
	return r
}

// ReorderLast reorders page to the last possible position.
func (t *TabView) ReorderLast(page *TabPage) bool {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	r := gobool(C.hdy_tab_view_reorder_last(t.native(), v1))
	return r
}

// ReorderPage reorders page to position.
//
// It's a programmer error to try to reorder a pinned page after a non-pinned
// one, or a non-pinned page before a pinned one.
func (t *TabView) ReorderPage(page *TabPage, position int) bool {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	v2 := C.gint(position)

	r := gobool(C.hdy_tab_view_reorder_page(t.native(), v1, v2))
	return r
}

// SelectNextPage selects the page after the currently selected page.
//
// If the last page was already selected, this function does nothing.
func (t *TabView) SelectNextPage() bool {
	r := gobool(C.hdy_tab_view_select_next_page(t.native()))
	return r
}

// SelectPreviousPage selects the page before the currently selected page.
//
// If the first page was already selected, this function does nothing.
func (t *TabView) SelectPreviousPage() bool {
	r := gobool(C.hdy_tab_view_select_previous_page(t.native()))
	return r
}

// SetDefaultIcon sets default page icon for t.
//
// If a page doesn't provide its own icon via TabPage:icon, default icon may be
// used instead for contexts where having an icon is necessary.
//
// TabBar will use default icon for pinned tabs in case the page is not loading,
// doesn't have an icon and an indicator. Default icon is never used for tabs
// that aren't pinned.
//
// By default, 'hdy-tab-icon-missing-symbolic' icon is used.
func (t *TabView) SetDefaultIcon(defaultIcon *glib.Icon) {
	v1 := (*C.GIcon)(unsafe.Pointer(defaultIcon.Native()))
	C.hdy_tab_view_set_default_icon(t.native(), v1)
}

// SetMenuModel sets the tab context menu model for t.
//
// When a context menu is shown for a tab, it will be constructed from the
// provided menu model. Use TabView::setup-menu signal to set up the menu
// actions for the particular tab.
func (t *TabView) SetMenuModel(menuModel *glib.MenuModel) {
	v1 := (*C.GMenuModel)(unsafe.Pointer(menuModel.Native()))
	C.hdy_tab_view_set_menu_model(t.native(), v1)
}

// SetPagePinned pins or unpins page.
//
// Pinned pages are guaranteed to be placed before all non-pinned pages; at any
// given moment the first TabView:n-pinned-pages pages in t are guaranteed to be
// pinned.
//
// When a page is pinned or unpinned, it's automatically reordered: pinning a
// page moves it after other pinned pages; unpinning a page moves it before
// other non-pinned pages.
//
// Pinned pages can still be reordered between each other.
//
// TabBar will display pinned pages in a compact form, never showing the title
// or close button, and only showing a single icon, selected in the following
// order:
//
// 1. TabPage:indicator-icon 2. A spinner if TabPage:loading is true 3.
// TabPage:icon 4. TabView:default-icon
//
// Pinned pages cannot be closed by default, see TabView::close-page for how to
// override that behavior.
func (t *TabView) SetPagePinned(page *TabPage, pinned bool) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	v2 := cbool(pinned)

	C.hdy_tab_view_set_page_pinned(t.native(), v1, v2)
}

// SetSelectedPage sets the currently selected page in t.
func (t *TabView) SetSelectedPage(selectedPage *TabPage) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(selectedPage.Native()))
	C.hdy_tab_view_set_selected_page(t.native(), v1)
}

// SetShortcutWidget sets the shortcut widget for t.
//
// Registers the following shortcuts on widget: * Ctrl+Page Up - switch to the
// previous page * Ctrl+Page Down - switch to the next page * Ctrl+Home - switch
// to the first page * Ctrl+End - switch to the last page * Ctrl+Shift+Page Up -
// move the current page backward * Ctrl+Shift+Page Down - move the current page
// forward * Ctrl+Shift+Home - move the current page at the start *
// Ctrl+Shift+End - move the current page at the end * Ctrl+Tab - switch to the
// next page, with looping * Ctrl+Shift+Tab - switch to the previous page, with
// looping * Alt+1-9 - switch to pages 1-9 * Alt+0 - switch to page 10
//
// These shortcuts are always available on t, this function is useful if they
// should be available globally.
func (t *TabView) SetShortcutWidget(widget gtk.IWidget) {
	v1 := cwidget(widget)
	C.hdy_tab_view_set_shortcut_widget(t.native(), v1)
}

// TransferPage transfers page from t to other_view. The page object will be
// reused.
//
// It's a programmer error to try to insert a pinned page after a non-pinned
// one, or a non-pinned page before a pinned one.
func (t *TabView) TransferPage(page *TabPage, otherView *TabView, position int) {
	v1 := (*C.HdyTabPage)(unsafe.Pointer(page.Native()))
	v2 := (*C.HdyTabView)(unsafe.Pointer(otherView.Widget.Native()))
	v3 := C.gint(position)

	C.hdy_tab_view_transfer_page(t.native(), v1, v2, v3)
}

type TitleBar struct {
	gtk.Bin
}

// wrapTitleBar wraps the given pointer to *TitleBar.
func wrapTitleBar(ptr unsafe.Pointer) *TitleBar {
	obj := glib.Take(ptr)
	return &TitleBar{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalTitleBar(p uintptr) (interface{}, error) {
	return wrapTitleBar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// TitleBarNew creates a new TitleBar.
func TitleBarNew() *TitleBar {
	return wrapTitleBar(unsafe.Pointer(C.hdy_title_bar_new()))
}

// native turns the current *TitleBar into the native C pointer type.
func (t *TitleBar) native() *C.HdyTitleBar {
	return (*C.HdyTitleBar)(gwidget(&t.Bin))
}

// GetSelectionMode returns whether whether t is in selection mode.
func (t *TitleBar) GetSelectionMode() bool {
	r := gobool(C.hdy_title_bar_get_selection_mode(t.native()))
	return r
}

// SetSelectionMode sets whether t is in selection mode.
func (t *TitleBar) SetSelectionMode(selectionMode bool) {
	v1 := cbool(selectionMode)
	C.hdy_title_bar_set_selection_mode(t.native(), v1)
}

type ValueObject struct {
	*glib.Object
}

// wrapValueObject wraps the given pointer to *ValueObject.
func wrapValueObject(ptr unsafe.Pointer) *ValueObject {
	obj := glib.Take(ptr)
	return &ValueObject{
		Object: obj,
	}
}

func marshalValueObject(p uintptr) (interface{}, error) {
	return wrapValueObject(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ValueObjectNew create a new ValueObject.
func ValueObjectNew(value *glib.Value) *ValueObject {
	v1 := (*C.GValue)(unsafe.Pointer(value.Native()))
	return wrapValueObject(unsafe.Pointer(C.hdy_value_object_new(v1)))
}

// ValueObjectNewString creates a new ValueObject. This is a convenience method
// to create a ValueObject that stores a string.
func ValueObjectNewString(string string) *ValueObject {
	v1 := C.CString(string)
	defer C.free(unsafe.Pointer(v1))
	return wrapValueObject(unsafe.Pointer(C.hdy_value_object_new_string(v1)))
}

// ValueObjectNewTakeString creates a new ValueObject. This is a convenience
// method to create a ValueObject that stores a string taking ownership of it.
func ValueObjectNewTakeString(string string) *ValueObject {
	v1 := C.CString(string)
	defer C.free(unsafe.Pointer(v1))
	return wrapValueObject(unsafe.Pointer(C.hdy_value_object_new_take_string(v1)))
}

// native turns the current *ValueObject into the native C pointer type.
func (v *ValueObject) native() *C.HdyValueObject {
	return (*C.HdyValueObject)(unsafe.Pointer(v.Native()))
}

// CopyValue copy data from the contained #GValue into dest.
func (v *ValueObject) CopyValue(dest *glib.Value) {
	v1 := (*C.GValue)(unsafe.Pointer(dest.Native()))
	C.hdy_value_object_copy_value(v.native(), v1)
}

// DupString returns a copy of the contained string if the value is of type
// TYPE_STRING.
func (v *ValueObject) DupString() string {
	r := C.GoString(C.hdy_value_object_dup_string(v.native()))
	return r
}

// GetString returns the contained string if the value is of type TYPE_STRING.
func (v *ValueObject) GetString() string {
	r := C.GoString(C.hdy_value_object_get_string(v.native()))
	return r
}

// GetValue return the contained value.
func (v *ValueObject) GetValue() *glib.Value {
	r := glib.ValueFromNative((unsafe.Pointer(C.hdy_value_object_get_value(v.native()))))
	return r
}

type ViewSwitcher struct {
	gtk.Bin
}

// wrapViewSwitcher wraps the given pointer to *ViewSwitcher.
func wrapViewSwitcher(ptr unsafe.Pointer) *ViewSwitcher {
	obj := glib.Take(ptr)
	return &ViewSwitcher{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalViewSwitcher(p uintptr) (interface{}, error) {
	return wrapViewSwitcher(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ViewSwitcherNew creates a new ViewSwitcher widget.
func ViewSwitcherNew() *ViewSwitcher {
	return wrapViewSwitcher(unsafe.Pointer(C.hdy_view_switcher_new()))
}

// native turns the current *ViewSwitcher into the native C pointer type.
func (v *ViewSwitcher) native() *C.HdyViewSwitcher {
	return (*C.HdyViewSwitcher)(gwidget(&v.Bin))
}

// GetNarrowEllipsize get the ellipsizing position of the narrow mode label. See
// (*ViewSwitcher).SetNarrowEllipsize().
func (v *ViewSwitcher) GetNarrowEllipsize() pango.EllipsizeMode {
	r := pango.EllipsizeMode(C.hdy_view_switcher_get_narrow_ellipsize(v.native()))
	return r
}

// GetPolicy gets the policy of v.
func (v *ViewSwitcher) GetPolicy() ViewSwitcherPolicy {
	r := ViewSwitcherPolicy(C.hdy_view_switcher_get_policy(v.native()))
	return r
}

// GetStack get the Stack being controlled by the ViewSwitcher.
//
// See: (*ViewSwitcher).SetStack()
func (v *ViewSwitcher) GetStack() *gtk.Stack {
	obj := glib.Take(unsafe.Pointer(C.hdy_view_switcher_get_stack(v.native())))
	r := &gtk.Stack{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
	return r
}

// SetNarrowEllipsize set the mode used to ellipsize the text in narrow mode if
// there is not enough space to render the entire string.
func (v *ViewSwitcher) SetNarrowEllipsize(mode pango.EllipsizeMode) {
	v1 := C.PangoEllipsizeMode(mode)
	C.hdy_view_switcher_set_narrow_ellipsize(v.native(), v1)
}

// SetPolicy sets the policy of v.
func (v *ViewSwitcher) SetPolicy(policy ViewSwitcherPolicy) {
	v1 := C.HdyViewSwitcherPolicy(policy)
	C.hdy_view_switcher_set_policy(v.native(), v1)
}

// SetStack sets the Stack to control.
func (v *ViewSwitcher) SetStack(stack *gtk.Stack) {
	v1 := (*C.GtkStack)(unsafe.Pointer(stack.Widget.Native()))
	C.hdy_view_switcher_set_stack(v.native(), v1)
}

type ViewSwitcherBar struct {
	gtk.Bin
}

// wrapViewSwitcherBar wraps the given pointer to *ViewSwitcherBar.
func wrapViewSwitcherBar(ptr unsafe.Pointer) *ViewSwitcherBar {
	obj := glib.Take(ptr)
	return &ViewSwitcherBar{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalViewSwitcherBar(p uintptr) (interface{}, error) {
	return wrapViewSwitcherBar(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ViewSwitcherBarNew creates a new ViewSwitcherBar widget.
func ViewSwitcherBarNew() *ViewSwitcherBar {
	return wrapViewSwitcherBar(unsafe.Pointer(C.hdy_view_switcher_bar_new()))
}

// native turns the current *ViewSwitcherBar into the native C pointer type.
func (v *ViewSwitcherBar) native() *C.HdyViewSwitcherBar {
	return (*C.HdyViewSwitcherBar)(gwidget(&v.Bin))
}

// GetPolicy gets the policy of v.
func (v *ViewSwitcherBar) GetPolicy() ViewSwitcherPolicy {
	r := ViewSwitcherPolicy(C.hdy_view_switcher_bar_get_policy(v.native()))
	return r
}

// GetReveal gets whether v should be revealed or not.
func (v *ViewSwitcherBar) GetReveal() bool {
	r := gobool(C.hdy_view_switcher_bar_get_reveal(v.native()))
	return r
}

// GetStack get the Stack being controlled by the ViewSwitcher.
func (v *ViewSwitcherBar) GetStack() *gtk.Stack {
	obj := glib.Take(unsafe.Pointer(C.hdy_view_switcher_bar_get_stack(v.native())))
	r := &gtk.Stack{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
	return r
}

// SetPolicy sets the policy of v.
func (v *ViewSwitcherBar) SetPolicy(policy ViewSwitcherPolicy) {
	v1 := C.HdyViewSwitcherPolicy(policy)
	C.hdy_view_switcher_bar_set_policy(v.native(), v1)
}

// SetReveal sets whether v should be revealed or not.
func (v *ViewSwitcherBar) SetReveal(reveal bool) {
	v1 := cbool(reveal)
	C.hdy_view_switcher_bar_set_reveal(v.native(), v1)
}

// SetStack sets the Stack to control.
func (v *ViewSwitcherBar) SetStack(stack *gtk.Stack) {
	v1 := (*C.GtkStack)(unsafe.Pointer(stack.Widget.Native()))
	C.hdy_view_switcher_bar_set_stack(v.native(), v1)
}

type ViewSwitcherTitle struct {
	gtk.Bin
}

// wrapViewSwitcherTitle wraps the given pointer to *ViewSwitcherTitle.
func wrapViewSwitcherTitle(ptr unsafe.Pointer) *ViewSwitcherTitle {
	obj := glib.Take(ptr)
	return &ViewSwitcherTitle{
		Bin: gtk.Bin{
			Container: gtk.Container{
				Widget: gtk.Widget{
					InitiallyUnowned: glib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalViewSwitcherTitle(p uintptr) (interface{}, error) {
	return wrapViewSwitcherTitle(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// ViewSwitcherTitleNew creates a new ViewSwitcherTitle widget.
func ViewSwitcherTitleNew() *ViewSwitcherTitle {
	return wrapViewSwitcherTitle(unsafe.Pointer(C.hdy_view_switcher_title_new()))
}

// native turns the current *ViewSwitcherTitle into the native C pointer type.
func (v *ViewSwitcherTitle) native() *C.HdyViewSwitcherTitle {
	return (*C.HdyViewSwitcherTitle)(gwidget(&v.Bin))
}

// GetPolicy gets the policy of v.
func (v *ViewSwitcherTitle) GetPolicy() ViewSwitcherPolicy {
	r := ViewSwitcherPolicy(C.hdy_view_switcher_title_get_policy(v.native()))
	return r
}

// GetStack get the Stack being controlled by the ViewSwitcher.
func (v *ViewSwitcherTitle) GetStack() *gtk.Stack {
	obj := glib.Take(unsafe.Pointer(C.hdy_view_switcher_title_get_stack(v.native())))
	r := &gtk.Stack{
		Container: gtk.Container{
			Widget: gtk.Widget{
				InitiallyUnowned: glib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
	return r
}

// GetSubtitle gets the subtitle of v. See (*ViewSwitcherTitle).SetSubtitle().
func (v *ViewSwitcherTitle) GetSubtitle() string {
	r := C.GoString(C.hdy_view_switcher_title_get_subtitle(v.native()))
	return r
}

// GetTitle gets the title of v. See (*ViewSwitcherTitle).SetTitle().
func (v *ViewSwitcherTitle) GetTitle() string {
	r := C.GoString(C.hdy_view_switcher_title_get_title(v.native()))
	return r
}

// GetTitleVisible get whether the title label of v is visible.
func (v *ViewSwitcherTitle) GetTitleVisible() bool {
	r := gobool(C.hdy_view_switcher_title_get_title_visible(v.native()))
	return r
}

// GetViewSwitcherEnabled gets whether v's view switcher is enabled.
//
// See (*ViewSwitcherTitle).SetViewSwitcherEnabled().
func (v *ViewSwitcherTitle) GetViewSwitcherEnabled() bool {
	r := gobool(C.hdy_view_switcher_title_get_view_switcher_enabled(v.native()))
	return r
}

// SetPolicy sets the policy of v.
func (v *ViewSwitcherTitle) SetPolicy(policy ViewSwitcherPolicy) {
	v1 := C.HdyViewSwitcherPolicy(policy)
	C.hdy_view_switcher_title_set_policy(v.native(), v1)
}

// SetStack sets the Stack to control.
func (v *ViewSwitcherTitle) SetStack(stack *gtk.Stack) {
	v1 := (*C.GtkStack)(unsafe.Pointer(stack.Widget.Native()))
	C.hdy_view_switcher_title_set_stack(v.native(), v1)
}

// SetSubtitle sets the subtitle of v. The subtitle should give a user
// additional details.
func (v *ViewSwitcherTitle) SetSubtitle(subtitle string) {
	v1 := C.CString(subtitle)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_view_switcher_title_set_subtitle(v.native(), v1)
}

// SetTitle sets the title of v. The title should give a user additional
// details. A good title should not include the application name.
func (v *ViewSwitcherTitle) SetTitle(title string) {
	v1 := C.CString(title)
	defer C.free(unsafe.Pointer(v1))
	C.hdy_view_switcher_title_set_title(v.native(), v1)
}

// SetViewSwitcherEnabled make v enable or disable its view switcher. If it is
// disabled, the title will be displayed instead. This allows to
// programmatically and prematurely hide the view switcher of v even if it fits
// in the available space.
//
// This can be used e.g. to ensure the view switcher is hidden below a certain
// window width, or any other constraint you find suitable.
func (v *ViewSwitcherTitle) SetViewSwitcherEnabled(enabled bool) {
	v1 := cbool(enabled)
	C.hdy_view_switcher_title_set_view_switcher_enabled(v.native(), v1)
}

type Window struct {
	gtk.Window
}

// wrapWindow wraps the given pointer to *Window.
func wrapWindow(ptr unsafe.Pointer) *Window {
	obj := glib.Take(ptr)
	return &Window{
		Window: gtk.Window{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: glib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalWindow(p uintptr) (interface{}, error) {
	return wrapWindow(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// WindowNew creates a new Window.
func WindowNew() *Window {
	return wrapWindow(unsafe.Pointer(C.hdy_window_new()))
}

// native turns the current *Window into the native C pointer type.
func (w *Window) native() *C.HdyWindow {
	return (*C.HdyWindow)(gwidget(&w.Window))
}

type WindowHandle struct {
	gtk.EventBox
}

// wrapWindowHandle wraps the given pointer to *WindowHandle.
func wrapWindowHandle(ptr unsafe.Pointer) *WindowHandle {
	obj := glib.Take(ptr)
	return &WindowHandle{
		EventBox: gtk.EventBox{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: glib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalWindowHandle(p uintptr) (interface{}, error) {
	return wrapWindowHandle(unsafe.Pointer(C.g_value_get_object((*C.GValue)(unsafe.Pointer(p))))), nil
}

// WindowHandleNew creates a new WindowHandle.
func WindowHandleNew() *WindowHandle {
	return wrapWindowHandle(unsafe.Pointer(C.hdy_window_handle_new()))
}

// native turns the current *WindowHandle into the native C pointer type.
func (w *WindowHandle) native() *C.HdyWindowHandle {
	return (*C.HdyWindowHandle)(gwidget(&w.EventBox))
}
