package hash

import (
	"image"

	"github.com/gotokatsuya/ipare/util"
)

type Difference struct {
	util.ResizeOption
}

func NewDifference() *Difference {
	const size = 8
	return &Difference{
		ResizeOption: util.NewResizeOption(size+1, size),
	}
}

func (d *Difference) Hash(src image.Image) uint64 {
	src = util.GrayscaleImage(d.ResizeImage(src))

	var (
		hash uint64
		one  uint64 = 1
	)

	for i := 0; i < d.Height; i++ {

		left := util.AverageGrayscaledRGBA(src.At(0, i).RGBA())

		for j := 1; j < d.Width; j++ {

			right := util.AverageGrayscaledRGBA(src.At(j, i).RGBA())

			if left > right {
				hash |= one
			}

			left = right
			one = one << 1
		}
	}
	return hash
}
