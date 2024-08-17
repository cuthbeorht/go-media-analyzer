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

func TestOpenMedia_ExpectEmptyFile(t *testing.T) {
	_, error := OpenMedia("/tmp/fixtures/empty.mp3")

	if error == nil {
		t.Fatal("Expecting error")
	}
}
