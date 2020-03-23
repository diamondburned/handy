package handy

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.Init(nil)
}

type propertySetter interface {
	SetProperty(string, interface{}) error
	GetProperty(string) (interface{}, error)
}

// lazy function
func makeTestProperty(c propertySetter) func(*testing.T) {
	return func(t *testing.T) {
		if err := c.SetProperty("name", "Astolfo"); err != nil {
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

		if name != "Astolfo" {
			t.Fatal("Unexpected name returned:", name)
		}
	}
}
