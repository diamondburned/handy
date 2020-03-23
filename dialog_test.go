package handy

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestDialog(t *testing.T) {
	w, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		t.Fatal("Fialed to make window:", err)
	}

	t.Run("glib properties", func(t *testing.T) {
		test := makeTestProperty(DialogNew(w))
		test(t)
	})

	d := DialogNew(w)

	// Apparently this is true by default, but there's no reliable way to test
	// this, since we can't set it beforehand.
	if d.GetNarrow() {
		t.Fatal("Unexpected GetNarrow false")
	}
}
