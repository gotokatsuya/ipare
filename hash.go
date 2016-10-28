package ipare

import (
	"image"

	"github.com/gotokatsuya/ipare/algorithm/hash"
)

type HashAlgorithm interface {
	Hash(image.Image) uint64
}

type Hash struct {
	Algorithm HashAlgorithm
}

func NewHash() *Hash {
	return &Hash{
		Algorithm: hash.NewDifference(),
	}
}

func (h *Hash) Do(src image.Image) uint64 {
	return h.Algorithm.Hash(src)
}

func (h *Hash) Compare(src1, src2 image.Image) int {
	return hammingDistance(h.Algorithm.Hash(src1), h.Algorithm.Hash(src2))
}

func hammingDistance(hash1, hash2 uint64) int {
	var dist int
	for v := hash1 ^ hash2; v != 0; v &= v - 1 {
		dist++
	}
	return dist
}
