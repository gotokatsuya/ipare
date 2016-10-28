package util

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) At() (int, int) {
	return p.X, p.Y
}
