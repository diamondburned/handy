package handy

import "testing"

func TestExpanderRow(t *testing.T) {
	t.Run("glib properties", makeTestProperty(ExpanderRowNew()))

	var e = ExpanderRowNew()

	e.SetTitle("hime arikawa")
	if title := e.GetTitle(); title != "hime arikawa" {
		t.Fatal("Unexpected title:", title)
	}

	e.SetExpanded(false)
	if e.GetExpanded() {
		t.Fatal("expanded true")
	}

	e.SetEnableExpansion(false)
	if e.GetEnableExpansion() {
		t.Fatal("enable expansion true")
	}

	e.SetShowEnableSwitch(false)
	if e.GetShowEnableSwitch() {
		t.Fatal("show enable switch true")
	}
}
