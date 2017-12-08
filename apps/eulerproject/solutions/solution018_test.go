package solutions

import (
	"math/big"
	"testing"
)

//
// https://projecteuler.net/problem=18
//

func TestSolution18_01(t *testing.T) {
	data := TriangeleData{
		filename:     "solution018_data_01.txt",
		size:         15,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	ans := data.do()
	required := big.NewInt(1074)
	if ans.Cmp(required) != 0 {
		t.Fatal("Fail")
	}
}

func TestSolution18_02(t *testing.T) {
	data := TriangeleData{
		filename:     "solution018_data_02.txt",
		size:         4,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	ans := data.do()
	required := big.NewInt(23)

	if ans.Cmp(required) != 0 {
		t.Fatal("Fail")
	}
}
