package media

import (
	"testing"
)

// func setup() func() {
// 	_, err := exec.Command("cp -R fixtures /tmp").Output()
// 	utils.CheckError(err)

// 	return func() {
// 		_, err := exec.Command("rm -rf /tmp/fixtures").Output()
// 		utils.CheckError(err)
// 	}
// }

func TestNewMedia_ExpectValidNewMedia(t *testing.T) {

	media := NewMedia("Foo", "MP3", "/tmp/fixtures/sample.mp3")

	if media != NewMedia("Foo", "MP3", "/tmp/fixtures/sample.mp3") {
		t.Fatal("Expected NewMedia{Foo, MP3}")
	}
}

func TestOpenMedia_ExpectEmptyFile(t *testing.T) {
	_, error := OpenMedia("/tmp/fixtures/empty.mp3")

	if error == nil {
		t.Fatal("Expecting error")
	}
}

func TestOpenMedia_ExpectMP3File(t *testing.T) {
	media, _ := OpenMedia("/tmp/fixtures/sample.mp3")

	expectedMedia := NewMedia("sample.mp3", "mp3", "/tmp/fixtures/sample.mp3")

	if !expectedMedia.equals(media) {
		t.Fatalf("Expeceted media, %s, to match %s", media, expectedMedia)
	}

}
