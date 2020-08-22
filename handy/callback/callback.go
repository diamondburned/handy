package callback

import (
	"sync"
	"sync/atomic"
)

var (
	registry = sync.Map{}   // uintptr -> func(Any) (Any)
	serial   = new(uintptr) // userData
)

func Assign(callback interface{}) uintptr {
	id := atomic.AddUintptr(serial, 1)
	registry.Store(id, callback)
	return uintptr(id)
}

func Get(ptr uintptr) interface{} {
	v, _ := registry.Load(ptr)
	return v
}

func Delete(ptr uintptr) {
	registry.Delete(ptr)
}
