package handy

import "testing"

func TestExpanderRow(t *testing.T) {
	t.Run("glib properties", makeTestProperty(ExpanderRowNew()))

	var e = ExpanderRowNew()

	e.SetTitle("hime arikawa")
	if title := e.GetTitle(); title != "hime arikawa" {
		t.Fatal("Unexpected title:", title)
	}

	e.SetExpanded(true)
	if !e.GetExpanded() {
		t.Fatal("expanded false")
	}

	e.SetEnableExpansion(true)
	if !e.GetEnableExpansion() {
		t.Fatal("enable expansion false")
	}

	e.SetShowEnableSwitch(true)
	if !e.GetShowEnableSwitch() {
		t.Fatal("show enable switch false")
	}
}
