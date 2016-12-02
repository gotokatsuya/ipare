package util

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

type ResizeOption struct {
	Width  int
	Height int
}

func NewResizeOption(size ...int) ResizeOption {
	switch len(size) {
	case 0:
		return ResizeOption{}
	case 1:
		return ResizeOption{
			Width:  size[0],
			Height: size[0],
		}
	default:
		return ResizeOption{
			Width:  size[0],
			Height: size[1],
		}
	}
}

func (r ResizeOption) ResizeImage(src image.Image) image.Image {
	return imaging.Resize(
		src,
		r.Width,
		r.Height,
		imaging.Lanczos,
	)
}

func AverageGrayscaledRGBA(r, g, b, a uint32) uint64 {
	return uint64(r)
}

func GrayscaleRGBA(r, g, b, a uint32) uint8 {
	f := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return uint8(f + 0.5)
}

func GrayscaleImage(src image.Image) image.Image {
	return imaging.Grayscale(src)
}

func CenterPoint(src image.Image) (int, int) {
	size := src.Bounds().Size()
	return size.X / 2, size.Y / 2
}

func Copy(src image.Image) image.Image {
	return imaging.Clone(src)
}

func Open(filename string) (image.Image, error) {
	return imaging.Open(filename)
}

func Decode(r io.Reader) (image.Image, error) {
	return imaging.Decode(r)
}

func AdjustValuesForPixelPartAlgorithm(centerX, centerY, partSize int) (adjustBeginPoint, adjustPartSize int, ok bool) {
	if partSize <= 0 {
		return
	}

	if centerX <= 0 || centerY <= 0 {
		return
	}

	// not square
	if centerX != centerY {
		return
	}

	beginPoint := centerX - partSize/2
	if beginPoint >= 0 {
		adjustBeginPoint = beginPoint
		adjustPartSize = partSize
		ok = true
		return
	}

	adjustBeginPoint = 0
	adjustPartSize = centerX * 2
	ok = true

	return
}
