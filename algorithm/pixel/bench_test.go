package pixel

import (
	"testing"

	"github.com/gotokatsuya/ipare/util"
)

func BenchmarkPart(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img1, err := util.Open("../../testdata/sample1.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img1 == nil {
			b.Fatal("img1 == nil.")
		}
		img2, err := util.Open("../../testdata/sample2.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img2 == nil {
			b.Fatal("img2 == nil.")
		}
		part := NewPart()
		part.Compare(img1, img2)
	}
}

func BenchmarkBlock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img1, err := util.Open("../../testdata/sample1.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img1 == nil {
			b.Fatal("img1 == nil.")
		}
		img2, err := util.Open("../../testdata/sample2.jpg")
		if err != nil {
			b.Fatal(err)
		}
		if img2 == nil {
			b.Fatal("img2 == nil.")
		}
		block := NewBlock()
		block.Compare(img1, img2)
	}
}
