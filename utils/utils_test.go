package utils

import "testing"

func TestForceUrl(t *testing.T) {
	t.Run("Without protocol", func(t *testing.T) {
		got := ForceUrl("reddit.com/r/unixporn")
		want := "http://reddit.com/r/unixporn"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("with protocol", func(t *testing.T) {
		got := ForceUrl("https://reddit.com/r/unixporn")
		want := "https://reddit.com/r/unixporn"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("feed url", func(t *testing.T) {
		got := ForceUrl("feed://someblog.xml")
		want := "http://someblog.xml"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
