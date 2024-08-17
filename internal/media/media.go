package media

import (
	"errors"
	"os"

	"github.com/cuthbeorht/go-media-analyzer/internal/utils"
	"github.com/h2non/filetype"
)

type Media struct {
	Name     string
	FileType string
	FullPath string
}

func NewMedia(name string, fileType string, fullPath string) Media {
	return Media{
		Name:     name,
		FileType: fileType,
		FullPath: fullPath,
	}
}

func (m *Media) equals(other *Media) bool {

	if m.Name == other.Name && m.FileType == other.FileType && m.FullPath == other.FullPath {
		return true
	}

	return false
}

func OpenMedia(fullPath string) (*Media, error) {

	binaryData, err := os.Open(fullPath)
	utils.CheckError(err)
	defer binaryData.Close()

	fileStats, err := binaryData.Stat()
	utils.CheckError(err)

	if fileStats.Size() == 0 {
		return nil, errors.New("expected file to have a size greater than zero")
	}

	buffer, err := os.ReadFile(fullPath)
	utils.CheckError(err)
	kind, _ := filetype.Match(buffer)

	newMedia := NewMedia(fileStats.Name(), kind.Extension, fullPath)
	return &newMedia, nil
}
