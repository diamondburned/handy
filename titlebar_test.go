package handy

import "testing"

func TestTitleBar(t *testing.T) {
	t.Run("glib properties", makeTestProperty(TitleBarNew()))

	var b = TitleBarNew()

	b.SetSelectionMode(true)
	if !b.GetSelectionMode() {
		t.Fatal("GetSelectionMode returned false")
	}
}
