#include <glib-2.0/glib.h>

// Helper function because cgo thinks gpointer is a different type, even though
// it's just a typedef to void*.
static gpointer conptr(void *ptr) { return ptr; }
