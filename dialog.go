package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (d *Dialog) native() *C.HdyDialog {
	return C.HDY_DIALOG(gwidget(d))
}

type Dialog struct {
	gtk.Dialog
}

func DialogNew(w gtk.IWindow) *Dialog {
	v := C.hdy_dialog_new((*C.GtkWindow)(unsafe.Pointer(w.ToWindow().Native())))
	obj := glib.Take(unsafe.Pointer(v))
	return &Dialog{gtk.Dialog{gtk.Window{gtk.Bin{container(obj)}}}}
}

func (d *Dialog) GetNarrow() bool {
	v := C.hdy_dialog_get_narrow(d.native())
	return gobool(v)
}
