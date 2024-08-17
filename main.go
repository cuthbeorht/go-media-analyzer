package main

import (
	"fmt"

	"github.com/cuthbeorht/go-media-analyzer/internal/media"
)

func main() {
	testFile := "/Volumes/DropboxData/New_Music/music/Transformers Discography/(2007) Transformers The Score/01 - Autobots.mp3"
	media := media.NewMedia("Foo", "MP3")
	fmt.Println("Media: ", media, testFile)
}
