package util

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func DecodeImageByPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func DecodeImageByFile(file io.Reader) (image.Image, error) {
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, nil
}
