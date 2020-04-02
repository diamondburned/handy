package handy

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestActionRow(t *testing.T) {
	t.Run("glib properties", makeTestProperty(ActionRowNew()))

	t.Run("(sub)title", func(t *testing.T) {
		a := ActionRowNew()

		a.SetTitle("astolfo")
		if v := a.GetTitle(); v != "astolfo" {
			t.Fatal("Unexpected title:", v)
		}

		a.SetSubtitle("best trap")
		if v := a.GetSubtitle(); v != "best trap" {
			t.Fatal("Unexpected subtitle:", v)
		}
	})

	t.Run("icon name", func(t *testing.T) {
		a := ActionRowNew()

		a.SetIconName("face-monkey-symbolic")
		if v := a.GetIconName(); v != "face-monkey-symbolic" {
			t.Fatal("Unexpected icon name:", v)
		}
	})

	t.Run("widget tests", func(t *testing.T) {
		a := ActionRowNew()

		l, _ := gtk.LabelNew("hideri big forehead")
		l.SetName("hideri label")

		a.SetActivatableWidget(l)
		w := a.GetActivatableWidget()

		if n, _ := w.GetName(); n != "hideri label" {
			t.Fatal("Unexpected name:", n)
		}

		// Can't test these:
		x, _ := gtk.LabelNew("the jar")
		a.AddAction(x)

		y, _ := gtk.LabelNew("the other jar")
		a.AddPrefix(y)
	})

	t.Run("underline", func(t *testing.T) {
		a := ActionRowNew()
		a.SetUseUnderline(true)
		if !a.GetUseUnderline() {
			t.Fatal("Unexpected false for underline")
		}
	})
}
