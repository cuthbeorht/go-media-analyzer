package id3

type Id3Header struct {
	Version    string
	Identifier string
	Flags      string
	Size       int
}

type Id3Tag struct {
	Header Id3Header
}

func NewId3Header(version string, identifier string, flags string, size int) Id3Header {
	return Id3Header{
		Version:    version,
		Identifier: identifier,
		Flags:      flags,
		Size:       size,
	}
}

func NewId3Tag(header Id3Header) Id3Tag {
	return Id3Tag{
		Header: header,
	}
}

func ReadId3(buff []byte) (Id3Tag, error) {
	rawHeader := buff[0:10]

	identifier := rawHeader[0:3]

	header := NewId3Header("", string(identifier), "", 0)

	return NewId3Tag(header), nil
}
