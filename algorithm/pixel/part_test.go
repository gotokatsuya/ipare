package pixel

import (
	"testing"

	"github.com/gotokatsuya/ipare/util"
)

func TestPart(t *testing.T) {
	part := NewPart()
	img1, err := util.DecodeImageByPath("../../testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img1 == nil {
		t.Fatal("img1 == nil.")
	}
	img2, err := util.DecodeImageByPath("../../testdata/sample2.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img2 == nil {
		t.Fatal("img2 == nil.")
	}
	res := part.Compare(img1, img2)
	t.Log(res)
}
