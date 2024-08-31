package id3

import (
	"encoding/hex"
	"errors"
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
	Flags      Id3HeaderFlags
	Size       int
}

type Id3Tag struct {
	Header Id3Header
}

const ID3_HEADER_SIZE_MAX_BIT_COUNT = 7

func NewId3HeaderFlags(unsync bool, extended bool, experimental bool) Id3HeaderFlags {
	return Id3HeaderFlags{
		Unsynchronisation:     unsync,
		ExtendedHeader:        extended,
		ExperimentalIndicator: experimental,
	}
}

func NewId3Header(version string, identifier string, flags Id3HeaderFlags, size int) Id3Header {
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

	identifier := identifier(rawHeader)

	version := version(rawHeader)

	flags := flags(rawHeader)

	size, _ := size(rawHeader)

	header := NewId3Header(version, string(identifier), flags, int(size))

	return NewId3Tag(header), nil
}

func identifier(fullRawHeader []byte) string {

	return string(fullRawHeader[0:3])
}

func version(fullRawHeader []byte) string {
	major := hex.EncodeToString(fullRawHeader[3:4])
	minor := hex.EncodeToString(fullRawHeader[4:5])

	return strings.Join([]string{major, ".", minor}, "")
}

func flags(fullRawHeader []byte) Id3HeaderFlags {
	flagsByte := fullRawHeader[6:7]

	unsynchronisationFlag := false
	extendedFlag := false
	experimentalFlag := false

	if (flagsByte[0]>>7)&1 == 1 {
		unsynchronisationFlag = true
	}

	if (flagsByte[0]>>6)&1 == 1 {
		extendedFlag = true
	}

	if (flagsByte[0]>>5)&1 == 1 {
		experimentalFlag = true
	}

	return NewId3HeaderFlags(unsynchronisationFlag, extendedFlag, experimentalFlag)
}

func size(fullRawHeader []byte) (uint32, error) {
	// https: //github.com/mikkyang/id3-go

	theSize := fullRawHeader[7:11]

	i := uint32(0)
	for _, b := range theSize {
		if b >= (1 << ID3_HEADER_SIZE_MAX_BIT_COUNT) {
			err := errors.New("byte integer: exceed max bit")
			return uint32(255), err
		}

		i = (i << ID3_HEADER_SIZE_MAX_BIT_COUNT) | uint32(b)
	}

	return i, nil
}
