package handy

// #include <handy.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Column is a container letting its child grow up to a given width.
//
// Description
//
// The HdyColumn widget limits the size of the widget it contains to a given
// maximum width. The expansion of the child from its minimum to its maximum
// size is eased out for a smooth transition.
// If the child requires more than the requested maximum width, it will be
// allocated the minimum width it can fit in instead.
type Column struct {
	gtk.Bin
}

func (c *Column) native() *C.HdyColumn {
	return C.HDY_COLUMN(gwidget(c))
}

func ColumnNew() *Column {
	v := C.hdy_column_new()
	obj := glib.Take(unsafe.Pointer(v))
	return &Column{gtk.Bin{container(obj)}}
}

func (c *Column) GetMaximumWidth() int {
	v := C.hdy_column_get_maximum_width(c.native())
	return int(v)
}

func (c *Column) SetMaximumWidth(maximumWidth int) {
	C.hdy_column_set_maximum_width(c.native(), C.gint(maximumWidth))
}

func (c *Column) GetLinearGrowthWidth() int {
	v := C.hdy_column_get_linear_growth_width(c.native())
	return int(v)
}

func (c *Column) SetLinearGrowthWidth(linearGrowthWidth int) {
	C.hdy_column_set_linear_growth_width(c.native(), C.gint(linearGrowthWidth))
}
