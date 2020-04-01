package handy

import "testing"

func TestPreferencesGroup(t *testing.T) {
	t.Run("glib properties", makeTestProperty(PreferencesGroupNew()))

	t.Run("title", func(t *testing.T) {
		var p = PreferencesGroupNew()

		p.SetTitle("hime arikawa")
		if v := p.GetTitle(); v != "hime arikawa" {
			t.Fatal("Unexpected title:", v)
		}
	})

	t.Run("description", func(t *testing.T) {
		var p = PreferencesGroupNew()

		p.SetDescription("nagisa")
		if v := p.GetDescription(); v != "nagisa" {
			t.Fatal("Unexpected description:", v)
		}
	})
}
