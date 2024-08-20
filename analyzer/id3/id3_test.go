package id3

import (
	"os"
	"testing"

	"github.com/cuthbeorht/go-media-analyzer/internal/utils"
)

func TestReadId3_ValidData(t *testing.T) {

	buffer, err := os.ReadFile("../../fixtures/sample.mp3")
	utils.CheckError(err)

	tag, _ := ReadId3(buffer)

	if tag.Header.Identifier != "ID3" || tag.Header.Version != "03.00" {
		t.Fatal("Expected Identifier should be ID3. It is ", tag.Header.Identifier)
	}
}
