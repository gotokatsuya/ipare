package hash

import (
	"image"

	"github.com/gotokatsuya/ipare/util"
)

type Average struct {
	util.ResizeOption
}

func NewAverage() *Average {
	const size = 8
	return &Average{
		ResizeOption: util.NewResizeOption(size),
	}
}

func (a *Average) Hash(src image.Image) uint64 {
	src = util.GrayscaleImage(a.ResizeImage(src))

	var (
		pixels    []uint64
		sumPixels uint64
	)

	for i := 0; i < a.Height; i++ {
		for j := 0; j < a.Width; j++ {

			pixel := util.AverageGrayscaledRGBA(src.At(j, i).RGBA())

			pixels = append(pixels, pixel)

			sumPixels += pixel

		}
	}

	average := uint64(sumPixels / uint64(a.Height*a.Width))

	var (
		hash uint64
		one  uint64 = 1
	)

	for _, pixel := range pixels {
		if pixel > average {
			hash |= one
		}
		one = one << 1
	}
	return hash
}
