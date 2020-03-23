package callback

// #cgo pkg-config: glib-2.0
// #include <glib-2.0/glib.h>
import "C"

import (
	"sync"
	"unsafe"
)

var (
	registry = map[int]*call{}
	regMutex = sync.RWMutex{}

	serial int
)

type call struct {
	ptr unsafe.Pointer
	fn  interface{}
}

func Assign(ptr unsafe.Pointer, callback interface{}) C.gpointer {
	regMutex.Lock()
	defer regMutex.Unlock()

	id := serial
	serial++

	registry[id] = &call{
		ptr: ptr,
		fn:  callback,
	}

	return C.gpointer(uintptr(id))
}

func Get(ptr C.gpointer) interface{} {
	regMutex.RLock()
	defer regMutex.RUnlock()

	if v, ok := registry[int(uintptr(ptr))]; ok {
		return v.fn
	}
	return nil
}

func Delete(ptr unsafe.Pointer) {
	regMutex.Lock()
	defer regMutex.Unlock()

	for i, call := range registry {
		if call.ptr == ptr {
			delete(registry, i)
		}
	}
}
