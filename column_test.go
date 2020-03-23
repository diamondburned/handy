package handy

import "testing"

func TestColumn(t *testing.T) {
	t.Run("glib properties", func(t *testing.T) {
		var c = ColumnNew()
		if err := c.SetProperty("name", "Hackadoll No. 3"); err != nil {
			t.Fatal("Failed to set property using glib:", err)
		}

		v, err := c.GetProperty("name")
		if err != nil {
			t.Fatal("Failed to get property using glib:", err)
		}

		name, ok := v.(string)
		if !ok {
			t.Fatal("Returned property is not of type string")
		}

		if name != "Hackadoll No. 3" {
			t.Fatal("Unexpected name returned:", name)
		}
	})

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
