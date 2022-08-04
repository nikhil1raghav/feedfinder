package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForceUrl(t *testing.T) {
	t.Run("Without protocol", func(t *testing.T) {
		got := ForceUrl("reddit.com/r/unixporn")
		want := "http://reddit.com/r/unixporn"
		assert.Equal(t, got, want)
	})
	t.Run("with protocol", func(t *testing.T) {
		got := ForceUrl("https://reddit.com/r/unixporn")
		want := "https://reddit.com/r/unixporn"
		assert.Equal(t, got, want)
	})
	t.Run("feed url", func(t *testing.T) {
		got := ForceUrl("feed://someblog.xml")
		want := "http://someblog.xml"
		assert.Equal(t, got, want)
	})
}

func TestJoinUrl(t *testing.T) {
	baseUrl := "http://youtube.com"
	completeUrl := "http://youtube.com/t/terms"
	partialUrl := "/t/terms"
	partialWithoutSlash := "t/terms"

	t.Run("Join partial url", func(t *testing.T) {
		got, err := JoinUrl(baseUrl, partialUrl)
		if err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, completeUrl, got)
	})

	t.Run("Join guess word", func(t *testing.T) {
		got, err := JoinUrl(baseUrl, "atom.xml")
		if err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, "http://youtube.com/atom.xml", got)
	})

	t.Run("Join partial without slash", func(t *testing.T) {
		got, err := JoinUrl(baseUrl, partialWithoutSlash)
		if err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, completeUrl, got)
	})
}
