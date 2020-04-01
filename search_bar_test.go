package handy

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestSearchBar(t *testing.T) {
	t.Run("glib properties", makeTestProperty(SearchBarNew()))

	t.Run("connect entry", func(t *testing.T) {
		e, err := gtk.EntryNew()
		if err != nil {
			t.Fatal("Failed to make *gtk.Entry")
		}

		b := SearchBarNew()
		b.ConnectEntry(e)
	})

	t.Run("search mode", func(t *testing.T) {
		b := SearchBarNew()
		b.SetSearchMode(true)

		if !b.GetSearchMode() {
			t.Fatal("Unexpected search mode false")
		}
	})

	t.Run("close button", func(t *testing.T) {
		b := SearchBarNew()
		b.SetShowCloseButton(true)

		if !b.GetShowCloseButton() {
			t.Fatal("Unexpected show close button false")
		}
	})
}
