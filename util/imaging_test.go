package util

import (
	"testing"
)

func TestAdjustValuesForPixelPartAlgorithm(t *testing.T) {

	var (
		cx   int
		cy   int
		size int
	)

	{
		cx = 0
		cy = 0
		size = 0

		_, _, ok := AdjustValuesForPixelPartAlgorithm(cx, cy, size)
		if ok {
			t.Fatal("invalid args")
		}
	}

	{
		cx = 1
		cy = 0
		size = 0

		_, _, ok := AdjustValuesForPixelPartAlgorithm(cx, cy, size)
		if ok {
			t.Fatal("invalid args")
		}
	}

	{
		cx = 2
		cy = 2
		size = 6

		adjustBeginPoint, adjustPartSize, ok := AdjustValuesForPixelPartAlgorithm(cx, cy, size)
		if !ok {
			t.Fatal("invalid args")
		}
		if adjustBeginPoint != 0 {
			t.Fatal("begin should be 0")
		}
		if adjustPartSize != 4 {
			t.Fatal("partSize should be 4")
		}
	}

	{
		cx = 4
		cy = 4
		size = 2

		adjustBeginPoint, adjustPartSize, ok := AdjustValuesForPixelPartAlgorithm(cx, cy, size)
		if !ok {
			t.Fatal("invalid args")
		}
		if adjustBeginPoint != 3 {
			t.Fatal("begin should be 3")
		}
		if adjustPartSize != 2 {
			t.Fatal("partSize should be 2")
		}
	}
}
