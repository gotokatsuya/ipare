package hash

import (
	"testing"

	"github.com/gotokatsuya/ipare/util"
)

func BenchmarkAverage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img1, err := util.DecodeImageByPath("../../testdata/sample1.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img1 == nil {
			b.Fatal("img1 == nil.")
		}
		img2, err := util.DecodeImageByPath("../../testdata/sample2.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img2 == nil {
			b.Fatal("img2 == nil.")
		}
		average := NewAverage()
		average.Hash(img1)
		average.Hash(img2)
	}
}

func BenchmarkDifference(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img1, err := util.DecodeImageByPath("../../testdata/sample1.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img1 == nil {
			b.Fatal("img1 == nil.")
		}
		img2, err := util.DecodeImageByPath("../../testdata/sample2.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img2 == nil {
			b.Fatal("img2 == nil.")
		}
		difference := NewDifference()
		difference.Hash(img1)
		difference.Hash(img1)
	}
}
