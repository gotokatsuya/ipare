package pixel

import (
	"image"
	"sync"

	"github.com/gotokatsuya/ipare/util"
)

type Block struct {
	util.ResizeOption

	BlockWidthSize  int
	BlockHeightSize int
}

func NewBlock() *Block {
	const size = 16
	b := &Block{
		ResizeOption: util.NewResizeOption(size),
	}
	const blockSize = 8
	return b.WithBlockSquare(blockSize)
}

func (b *Block) WithBlockSquare(size int) *Block {
	b.BlockWidthSize = size
	b.BlockHeightSize = size
	return b
}

func (b *Block) Compare(src1, src2 image.Image) int {
	src1 = b.ResizeImage(src1)
	src2 = b.ResizeImage(src2)

	var sumDifference int

	var w sync.WaitGroup
	var m sync.Mutex

	for y := 0; y < b.Height; y = y + b.BlockHeightSize {
		for x := 0; x < b.Width; x = x + b.BlockWidthSize {
			begin := image.Point{X: x, Y: y}
			end := image.Point{X: x + b.BlockWidthSize, Y: y + b.BlockHeightSize}
			w.Add(1)
			go func(src1, src2 image.Image, begin, end image.Point) {
				defer w.Done()

				difference := differentInBlock(src1, src2, begin, end)

				m.Lock()
				sumDifference += difference
				m.Unlock()

			}(src1, src2, begin, end)
		}
	}
	w.Wait()

	return sumDifference
}

func differentInBlock(src1, src2 image.Image, begin, end image.Point) int {

	var difference int

	for y := begin.Y; y < end.Y; y++ {
		for x := begin.X; x < end.X; x++ {

			pixel1 := util.GrayscaleRGBA(src1.At(y, x).RGBA())
			pixel2 := util.GrayscaleRGBA(src2.At(y, x).RGBA())

			if pixel1 != pixel2 {
				difference++
			}

		}
	}

	return difference
}
