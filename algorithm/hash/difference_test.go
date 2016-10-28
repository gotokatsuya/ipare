package hash

import (
	"testing"

	"github.com/gotokatsuya/ipare/util"
)

func TestDifference(t *testing.T) {
	difference := NewDifference()
	img, err := util.DecodeImageByPath("../../testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("img == nil.")
	}
	res := difference.Hash(img)
	if res == 0 {
		t.Fatal(res)
	} else {
		t.Log(res)
	}
}
