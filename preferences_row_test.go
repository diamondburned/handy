package handy

import "testing"

func TestPreferencesRow(t *testing.T) {
	t.Run("glib properties", makeTestProperty(PreferencesRowNew()))

	t.Run("title", func(t *testing.T) {
		var p = PreferencesRowNew()

		p.SetTitle("hime arikawa")
		if v := p.GetTitle(); v != "hime arikawa" {
			t.Fatal("Unexpected title:", v)
		}
	})

	t.Run("underline", func(t *testing.T) {
		var p = PreferencesRowNew()

		p.SetUseUnderline(true)
		if !p.GetUseUnderline() {
			t.Fatal("Unexpected false for GetUseUnderline")
		}

		p.SetUseUnderline(false)
		if p.GetUseUnderline() {
			t.Fatal("Unexpected true for GetUseUnderline")
		}
	})
}
