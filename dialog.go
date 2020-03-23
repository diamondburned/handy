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

// Dialog is an adaptive dialog.
//
// Description
//
// A GtkDialog that adapts to smaller displays. In the smaller view a HdyDialog
// matches its size to that of its parent and for "Presentation Dialogs" uses a
// back button rather than close button to dismiss.
//
// It's recommended that dialog contents are wrapped in a GtkScrolledWindow to
// ensure they don't overflow the screen.
//
// HdyDialog works best when “use-header-bar” is TRUE (which is the case when
// using hdy_dialog_new()).
//
// If you want to replace the titlebar by your own, we recommend using
// HdyHeaderBar as it will retain the abiity to present a back button when the
// dialog is small. HdyHeaderBar doesn't have to be its direct child and you can
// use any complex contraption you like as the dialog's titlebar.
//
// https://developer.puri.sm/projects/libhandy/unstable/HdyDialog.html
type Dialog struct {
	gtk.Dialog
}

// DialogNew creates a HdyDialog with "transient-for" set to parent.
func DialogNew(parent gtk.IWindow) *Dialog {
	v := C.hdy_dialog_new((*C.GtkWindow)(unsafe.Pointer(parent.ToWindow().Native())))
	obj := glib.Take(unsafe.Pointer(v))
	return &Dialog{gtk.Dialog{gtk.Window{gtk.Bin{container(obj)}}}}
}

// GetNarrow gets whether self is narrow.
func (d *Dialog) GetNarrow() bool {
	v := C.hdy_dialog_get_narrow(d.native())
	return gobool(v)
}
