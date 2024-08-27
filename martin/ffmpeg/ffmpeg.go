package ffmpeg

import (
	"errors"
	"image"
	_ "image/png" //png doesn't work without this
	"os"
)

type Image struct {
	Path string
}

type Martin struct {
	Width  int
	Height int
	Image  Image
}

var (
	ErrInvalidMartinPath = errors.New("Could not find Martin at path")
)

func NewMartin(path string) (*Martin, error) {

	reader, err := os.Open(path)
	if err != nil {
		return nil, ErrInvalidMartinPath
	}
	defer reader.Close()

	martinFile, _, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, err
	}

	return &Martin{
		Width:  martinFile.Width,
		Height: martinFile.Height,
		Image: Image{
			Path: path,
		},
	}, nil

}
