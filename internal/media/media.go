package media

import (
	"errors"
	"os"

	"github.com/cuthbeorht/go-media-analyzer/internal/utils"
)

type Media struct {
	Name     string
	FileType string
}

func NewMedia(name string, fileType string) Media {
	return Media{
		Name:     name,
		FileType: fileType,
	}
}

func OpenMedia(fullPath string) (*Media, error) {

	binaryData, err := os.Open(fullPath)
	utils.CheckError(err)

	fileStats, err := binaryData.Stat()
	utils.CheckError(err)

	if fileStats.Size() == 0 {
		return nil, errors.New("expected file to have a size greater than zero")
	}
	newMedia := NewMedia("Foo", "MP3")
	return &newMedia, nil
}
