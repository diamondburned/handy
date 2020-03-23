package handy

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestLeaflet(t *testing.T) {
	t.Run("glib properties", func(t *testing.T) {
		var l = LeafletNew()
		if err := l.SetProperty("name", "astolfo"); err != nil {
			t.Fatal("Failed to set property using glib:", err)
		}

		v, err := l.GetProperty("name")
		if err != nil {
			t.Fatal("Failed to get property using glib:", err)
		}

		name, ok := v.(string)
		if !ok {
			t.Fatal("Returned property is not of type string")
		}

		if name != "astolfo" {
			t.Fatal("Unexpected name returned:", name)
		}
	})

	t.Run("fold", func(t *testing.T) {
		var l = LeafletNew()

		if f := l.GetFold(); f != FOLD_UNFOLDED {
			t.Fatal("Unexpected Leaflet GetFold:", f)
		}
	})

	t.Run("visible child", func(t *testing.T) {
		var l = LeafletNew()

		// Make a fake widget:
		b, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		if err != nil {
			t.Fatal("Failed to make box:", err)
		}

		// Give it a name:
		b.SetName("astolfo")

		l.Add(b)
		l.SetVisibleChild(b)

		// Check the widget and its name:
		n, err := l.GetVisibleChild().GetName()
		if err != nil {
			t.Fatal("Failed to get name:", err)
		}

		if n != "astolfo" {
			t.Fatal("Wrong name:", n)
		}

		// This fails.

		// t.Run("visible child name", func(t *testing.T) {
		// 	// Child has to be visible:
		// 	b.Show()

		// 	// l.SetVisibleChildName("astolfo")

		// 	n = l.GetVisibleChildName()
		// 	if n != "astolfo" {
		// 		t.Fatal("Visible child name mismatch:", n)
		// 	}
		// })
	})

	t.Run("homogeneous", func(t *testing.T) {
		var l = LeafletNew()

		// Set:
		l.SetHomogeneous(FOLD_FOLDED, gtk.ORIENTATION_VERTICAL, true)

		// Get and check:
		if !l.GetHomogeneous(FOLD_FOLDED, gtk.ORIENTATION_VERTICAL) {
			t.Fatal("Homogeneous returned false")
		}
	})

	t.Run("transition type", func(t *testing.T) {
		var l = LeafletNew()

		// Set:
		l.SetTransitionType(LEAFLET_TRANSITION_TYPE_OVER)

		// Check
		if trans := l.GetTransitionType(); trans != LEAFLET_TRANSITION_TYPE_OVER {
			t.Fatal("Unexpected transition type:", trans)
		}
	})

	t.Run("transition duration", func(t *testing.T) {
		var l = LeafletNew()

		// Set:
		l.SetModeTransitionDuration(300) // ms

		// Check
		if ms := l.GetModeTransitionDuration(); ms != 300 {
			t.Fatal("Unexpected transition duration:", ms)
		}
	})

	// transition running is skipped

	t.Run("interpolate size", func(t *testing.T) {
		var l = LeafletNew()

		l.SetInterpolateSize(true)
		if !l.GetInterpolateSize() {
			t.Fatal("Interpolate size returned unexpected false")
		}
	})

	t.Run("can swipe", func(t *testing.T) {
		var l = LeafletNew()

		l.SetCanSwipeBack(true)
		if !l.GetCanSwipeBack() {
			t.Fatal("Can swipe back returned unexpected false")
		}

		l.SetCanSwipeForward(true)
		if !l.GetCanSwipeForward() {
			t.Fatal("Can swipe forward returned unexpected false")
		}
	})
}
