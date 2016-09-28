package solutions

import (
	"math/big"
	"testing"
)

func TestSolution67_01(t *testing.T) {
	data := TriangeleData{
		filename:     "solution067_data_01.txt",
		size:         100,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	ans := data.do()
	required := big.NewInt(7273)

	if ans.Cmp(required) != 0 {
		t.Fatal("Fail")
	}
}
