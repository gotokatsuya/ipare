package ipare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotokatsuya/ipare/algorithm/pixel"
	"github.com/gotokatsuya/ipare/util"
)

func TestSamePixel(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm PixelAlgorithm
		img1      string
		img2      string
	}{
		{"Part", pixel.NewPart(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
		{"Block", pixel.NewBlock(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		pixel := NewPixel()
		pixel.Algorithm = tt.algorithm

		img1, err := util.Open(tt.img1)
		assert.NoError(err, target)
		img2, err := util.Open(tt.img2)
		assert.NoError(err, target)

		distance := pixel.Compare(img1, img2)
		assert.Equal(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
	}
}

func TestUnSamePixel(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm PixelAlgorithm
		img1      string
		img2      string
	}{
		{"Part", pixel.NewPart(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
		{"Block", pixel.NewBlock(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		pixel := NewPixel()
		pixel.Algorithm = tt.algorithm

		img1, err := util.Open(tt.img1)
		assert.NoError(err, target)
		img2, err := util.Open(tt.img2)
		assert.NoError(err, target)

		distance := pixel.Compare(img1, img2)
		assert.NotEqual(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
	}
}
