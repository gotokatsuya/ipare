package pixel

import (
	"testing"

	"github.com/gotokatsuya/ipare/util"
)

func TestPart(t *testing.T) {
	part := NewPart()
	img1, err := util.Open("../../testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img1 == nil {
		t.Fatal("img1 == nil.")
	}
	img2, err := util.Open("../../testdata/sample2.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img2 == nil {
		t.Fatal("img2 == nil.")
	}
	resizeOpt := util.NewResizeOption(100)
	img1 = resizeOpt.ResizeImage(img1)
	img2 = resizeOpt.ResizeImage(img2)
	res := part.WithBeginPoint(util.CenterPoint(img1)).Compare(img1, img2)
	t.Log(res)
}
