package media

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
	return NewMedia("Foo", "MP3")
}
