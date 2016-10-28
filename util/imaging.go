package util

import (
	"image"
	"image/draw"

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
	r := src.Bounds()
	dst := image.NewNRGBA(r)
	draw.Draw(dst, r, src, r.Min, draw.Src)
	return dst
}
