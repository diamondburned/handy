package handy

// #include <handy.h>
// #include "util.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func gpointer(ptr unsafe.Pointer) C.gpointer {
	return C.conptr(ptr)
}

func nwidget(w gtk.IWidget) unsafe.Pointer {
	return unsafe.Pointer(w.ToWidget().Native())
}

func cwidget(w gtk.IWidget) *C.GtkWidget {
	return (*C.GtkWidget)(nwidget(w))
}

func gwidget(w gtk.IWidget) C.gpointer {
	return gpointer(unsafe.Pointer(w.ToWidget().GObject))
}

func widget(obj *glib.Object) *gtk.Widget {
	return &gtk.Widget{glib.InitiallyUnowned{obj}}
}

func container(obj *glib.Object) gtk.Container {
	return gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}
}

func cbool(val bool) C.gboolean {
	if val {
		return C.TRUE
	}
	return C.FALSE
}

func gobool(val C.gboolean) bool {
	return val != C.FALSE
}
