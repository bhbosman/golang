package solutions

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=26
//

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func TestSolution26_01(t *testing.T) {

	size := 10000
	repo := func(aa, bb int) ([]int, bool) {
		a := aa
		b := bb
		for {
			d := a / b
			a = (a - b*d) * 10
			if d == 0 {
				break
			}
		}
		data := make([]int, size, size)
		index := 0
		for {
			d := a / b
			r := a % b
			if r == 0 {
				return nil, false
			}
			a = (a - b*d) * 10
			if index < len(data) {
				data[index] = d
			} else {
				break
			}
			index++
		}
		return data, true
	}
	findrepo := func(data []int) (int, bool, []int) {
		for i := 1; i < size/2; i++ {
			Reciprocal := data[0:i]
			found := true
			for k := i; k < size-len(Reciprocal); k += len(Reciprocal) {
				b, e := k, k+i
				t := testEq(Reciprocal, data[b:e])
				if !t {
					found = false
					break
				}
			}
			if found {
				return i, true, Reciprocal
			}
		}
		return -1, false, nil
	}

	ans := -1

	for i := 1; i <= 1000; i++ {
		if data, ok := repo(1, i); ok {
			if l, ok, data := findrepo(data); ok {
				fmt.Println(i, l, data)
				if ans < l {

					ans = l

				}
			}
		}

	}

}
