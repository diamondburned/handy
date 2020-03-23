package handy

// #include <handy.h>
// extern GdkPixbuf* avatarImageLoadFunc(gint size, gpointer key);
import "C"

/*
import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/handy/callback"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type AvatarImageLoadFunc func(size int) *gdk.Pixbuf

// export avatarImageLoadFunc
func avatarImageLoadFunc(size C.gint, key C.gpointer) *C.GdkPixbuf {
	fn := callback.Get(key).(AvatarImageLoadFunc)
	return fn(int(size)).NativePrivate()
}

type Avatar struct {
	gtk.DrawingArea
}

func (a *Avatar) native() *C.HdyAvatar {
	return C.HDY_AVATAR(gwidget(a))
}

func AvatarNew() *Avatar {
	v := C.hdy_avatar_new()
	obj := glib.Take(unsafe.Pointer(v))

	a := &Avatar{gtk.DrawingArea{widget(obj)}}

	// We need to do this to clear off all the handlers:
	runtime.SetFinalizer(a, func() {
		callback.Delete(nwidget(a))
	})

	return a
}

func (a *Avatar) GetText() string {
	c := C.hdy_avatar_get_text(a.native())
	return C.GoString(c)
}

func (a *Avatar) SetText(text string) {
	C.hdy_avatar_set_text(a.native(), C.CString(text))
}

func (a *Avatar) GetShowInitials() bool {
	c := C.hdy_avatar_get_show_initials(a.native())
	return gobool(c)
}

func (a *Avatar) SetShowInitials(showInitials bool) {
	C.hdy_avatar_set_show_initials(a.native(), cbool(showInitials))
}

func (a *Avatar) SetImageLoadFunc(fn AvatarImageLoadFunc) {
	callback.Assign(nwidget(a), fn)
	C.hdy_avatar_set_image_load_func(
		a.native(), (C.HdyAvatarImageLoadFunc)(avatarImageLoadFunc), C.NULL, C.NULL)
}

func (a *Avatar) GetSize() int {
	c := C.hdy_avatar_get_size(a.native())
	return int(c)
}

func (a *Avatar) SetSize(size int) {
	C.hdy_avatar_set_size(a.native(), C.int(size))
}
*/
