package handy

import "testing"

func TestPreferencesPage(t *testing.T) {
	t.Run("glib properties", makeTestProperty(PreferencesPageNew()))

	t.Run("icon name", func(t *testing.T) {
		var p = PreferencesPageNew()

		p.SetIconName("")
		if i := p.GetIconName(); i != "" {
			t.Fatal("Unexpected icon name:", i)
		}

		p.SetIconName("astolfo")
		if i := p.GetIconName(); i != "astolfo" {
			t.Fatal("Unexpected icon name:", i)
		}
	})

	t.Run("title", func(t *testing.T) {
		var p = PreferencesPageNew()

		p.SetTitle("")
		if i := p.GetTitle(); i != "" {
			t.Fatal("Unexpected title:", i)
		}

		p.SetTitle("hideri")
		if i := p.GetTitle(); i != "hideri" {
			t.Fatal("Unexpected title:", i)
		}
	})
}
