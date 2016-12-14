package solutions

import "testing"
import "fmt"

//
// https://projecteuler.net/problem=27
//

func matrixData(n int64) int64 {
	if n%2 == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return 4*(n*n) - 6*(n-1) + matrixData(n-2)
}

func TestSolution28_01(t *testing.T) {

	fmt.Println(1, matrixData(1))
	fmt.Println(3, matrixData(3))
	fmt.Println(5, matrixData(5))
	fmt.Println(7, matrixData(7))
	fmt.Println(1001, matrixData(1001))
}
