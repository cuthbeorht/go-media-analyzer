package id3

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Id3HeaderFlags struct {
	Unsynchronisation     bool
	ExtendedHeader        bool
	ExperimentalIndicator bool
}

type Id3Header struct {
	Version    string
	Identifier string
	Flags      string
	Size       int
}

type Id3Tag struct {
	Header Id3Header
}

func NewId3HeaderFlags(unsync bool, extended bool, experimental bool) Id3HeaderFlags {
	return Id3HeaderFlags{
		Unsynchronisation:     unsync,
		ExtendedHeader:        extended,
		ExperimentalIndicator: experimental,
	}
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

	major := hex.EncodeToString(rawHeader[3:4])
	minor := hex.EncodeToString(rawHeader[4:5])

	flags, _ := strconv.Atoi(string(rawHeader[6:7]))
	fmt.Println("Flags: ", flags)

	size := rawHeader[7:11]
	fmt.Println("Size: ", size)

	header := NewId3Header(strings.Join([]string{major, ".", minor}, ""), string(identifier), "", 0)

	return NewId3Tag(header), nil
}
