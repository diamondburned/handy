package handy

import "testing"

func TestColumn(t *testing.T) {
	t.Run("glib properties", makeTestProperty(ColumnNew()))

	var c = ColumnNew()

	c.SetMaximumWidth(69)
	if w := c.GetMaximumWidth(); w != 69 {
		t.Fatal("Unexpected maximum width:", w)
	}

	c.SetLinearGrowthWidth(420)
	if w := c.GetLinearGrowthWidth(); w != 420 {
		t.Fatal("Unexpected linear growth width:", w)
	}
}
