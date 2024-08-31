package id3

import (
	"os"
	"testing"

	"github.com/cuthbeorht/go-media-analyzer/internal/utils"
)

func TestReadId3_ExpectCorrectIdentifier(t *testing.T) {

	buffer, err := os.ReadFile("../../test/fixtures/sample.mp3")
	utils.CheckError(err)

	tag, _ := ReadId3(buffer)

	if tag.Header.Identifier != "ID3" {
		t.Fatal("Expected Identifier should be ID3. It is ", tag.Header.Identifier)
	}
}

func TestReadId3_ExpectCorrectVersion(t *testing.T) {

	buffer, err := os.ReadFile("../../test/fixtures/sample.mp3")
	utils.CheckError(err)

	tag, _ := ReadId3(buffer)

	if tag.Header.Version != "03.00" {
		t.Fatal("Expected Identifier should be ID3. It is ", tag.Header.Identifier)
	}
}

func TestReadId3_ExpectFalseHeaderFlag(t *testing.T) {
	buffer, err := os.ReadFile("../../test/fixtures/sample.mp3")
	utils.CheckError(err)

	tag, _ := ReadId3(buffer)

	if tag.Header.Flags.ExperimentalIndicator || tag.Header.Flags.ExtendedHeader || tag.Header.Flags.Unsynchronisation {
		t.Fatal("Expected Header flags to all be false. They are ", tag.Header.Flags)
	}
}

func TestReadId3_ExpectCorrectSize(t *testing.T) {
	buffer, err := os.ReadFile("../../test/fixtures/sample.mp3")
	utils.CheckError(err)

	tag, _ := ReadId3(buffer)

	if tag.Header.Size != 156244 {
		t.Fatal("Expected size of 156244. Got ", tag.Header.Size)
	}
}
