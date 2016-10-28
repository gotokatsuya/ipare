package pixel

import (
	"image"

	"github.com/gotokatsuya/ipare/util"
)

type Part struct {
	PartBeginPoint image.Point
	PartWidthSize  int
	PartHeightSize int
}

const (
	defaultPartSize = 8
)

func NewPart() *Part {
	p := &Part{
		PartBeginPoint: image.Point{X: 0, Y: 0},
	}
	return p.WithPartSquare(defaultPartSize)
}

func (p *Part) WithBeginPoint(x, y int) *Part {
	p.PartBeginPoint = image.Point{X: x, Y: y}
	return p
}

func (p *Part) WithPartSquare(size int) *Part {
	p.PartWidthSize = size
	p.PartHeightSize = size
	return p
}

func (p *Part) partEndPoint() (int, int) {
	return p.PartBeginPoint.X + p.PartWidthSize, p.PartBeginPoint.Y + p.PartHeightSize
}

func (p *Part) adjustPoint(src image.Image) {
	partEndPointX, partEndPointY := p.partEndPoint()
	srcBounds := src.Bounds()
	if partEndPointX > srcBounds.Max.X || partEndPointY > srcBounds.Max.Y {
		p.WithBeginPoint(0, 0).WithPartSquare(defaultPartSize)
	}
}

func (p *Part) preCompare(src1, src2 image.Image) {
	p.adjustPoint(src1)
	p.adjustPoint(src2)
}

func (p *Part) Compare(src1, src2 image.Image) int {

	p.preCompare(src1, src2)

	var sumDifference int

	partEndPointX, partEndPointY := p.partEndPoint()

	for y := p.PartBeginPoint.Y; y < partEndPointY; y++ {
		for x := p.PartBeginPoint.X; x < partEndPointX; x++ {

			pixel1 := util.GrayscaleRGBA(src1.At(y, x).RGBA())
			pixel2 := util.GrayscaleRGBA(src2.At(y, x).RGBA())

			if pixel1 != pixel2 {
				sumDifference++
			}

		}
	}
	return sumDifference
}
