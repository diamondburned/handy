package handy

// #include <handy.h>
// #include "util.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gtk"
)

func gpointer(ptr unsafe.Pointer) C.gpointer {
	return C.conptr(ptr)
}

func native(w gtk.IWidget) unsafe.Pointer {
	return unsafe.Pointer(w.ToWidget().Native())
}

func cwidget(w gtk.IWidget) *C.GtkWidget {
	return (*C.GtkWidget)(native(w))
}

func cbool(val bool) C.gboolean {
	if val {
		return C.TRUE
	}
	return C.FALSE
}
