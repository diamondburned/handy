// Package handy is a Golang binding to libhandy, which is a Gtk library that
// helps with mobile UI development.
//
// Copyright (C) 2020 diamondburned
package handy

// #cgo pkg-config: libhandy-0.0
// #cgo CPPFLAGS: -DHANDY_USE_UNSTABLE_API
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

type Fold int

const (
	FOLD_UNFOLDED Fold = C.HDY_FOLD_UNFOLDED
	FOLD_FOLDED   Fold = C.HDY_FOLD_FOLDED
)
