// Package handy is a Golang binding to libhandy, which is a Gtk library that
// helps with mobile UI development.
//
// Copyright (C) 2020 diamondburned
package handy

// #cgo pkg-config: libhandy-0.0
// #cgo CPPFLAGS: -DHANDY_USE_UNSTABLE_API
// #include <handy.h>
import "C"

type Fold int

const (
	FOLD_UNFOLDED Fold = C.HDY_FOLD_UNFOLDED
	FOLD_FOLDED   Fold = C.HDY_FOLD_FOLDED
)
