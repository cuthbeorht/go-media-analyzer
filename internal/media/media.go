package media

import (
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

func OpenMedia(fullPath string) Media {

	_, err := os.ReadFile(fullPath)
	utils.CheckError(err)

	return NewMedia("Foo", "MP3")
}
