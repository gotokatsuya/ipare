package ipare

import (
	"image"

	"github.com/gotokatsuya/ipare/algorithm/pixel"
)

type PixelAlgorithm interface {
	Compare(image.Image, image.Image) int
}

type Pixel struct {
	Algorithm PixelAlgorithm
}

func NewPixel() *Pixel {
	return &Pixel{
		Algorithm: pixel.NewPart(),
	}
}

func (p *Pixel) Compare(src1, src2 image.Image) int {
	return p.Algorithm.Compare(src1, src2)
}
