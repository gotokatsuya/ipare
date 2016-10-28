package ipare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotokatsuya/ipare/algorithm/hash"
	"github.com/gotokatsuya/ipare/util"
)

func TestSameHash(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm HashAlgorithm
		img1      string
		img2      string
	}{
		{"Average", hash.NewAverage(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
		{"Difference", hash.NewDifference(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := util.Open(tt.img1)
		assert.NoError(err, target)
		img2, err := util.Open(tt.img2)
		assert.NoError(err, target)

		distance := hash.Compare(img1, img2)
		assert.Equal(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
	}
}

func TestUnSameHash(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm HashAlgorithm
		img1      string
		img2      string
	}{
		{"Average", hash.NewAverage(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
		{"Difference", hash.NewDifference(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := util.Open(tt.img1)
		assert.NoError(err, target)
		img2, err := util.Open(tt.img2)
		assert.NoError(err, target)

		distance := hash.Compare(img1, img2)
		assert.NotEqual(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
	}
}

func TestSimilarHash(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name        string
		algorithm   HashAlgorithm
		img1        string
		img2        string
		maxDistance int
	}{
		{"Average", hash.NewAverage(), "testdata/sample1.jpg", "testdata/sample1-small.jpg", 5},
		{"Difference", hash.NewDifference(), "testdata/sample1.jpg", "testdata/sample1-small.jpg", 5},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)
		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := util.Open(tt.img1)
		assert.NoError(err, target)
		img2, err := util.Open(tt.img2)
		assert.NoError(err, target)

		distance := hash.Compare(img1, img2)
		if tt.maxDistance < distance {
			t.Fatal("Not similar.", target)
		} else {
			t.Log(distance, target)
		}
	}
}
