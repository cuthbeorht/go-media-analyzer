package media

import (
	"testing"
)

func TestNewMedia_ExpectValidNewMedia(t *testing.T) {
	media := NewMedia("Foo", "MP3")

	if media != NewMedia("Foo", "MP3") {
		t.Fatal("Expected NewMedia{Foo, MP3}")
	}
}

func TestOpenMedia_ExpectDummyMedia(t *testing.T) {
	media := OpenMedia("/dummy/path")

	if media != NewMedia("Foo", "MP3") {
		t.Fatalf("Expecting %s to equal NewMedia{Name: Foo, FileType: MP3}", media)
	}
}
