package solutions

import "testing"
import "fmt"

//
// https://projecteuler.net/problem=34
//

func TestSolution34_01(t *testing.T) {
    init := func()[]int{
        result := make([]int, 9, 9)
        for i := 
    }
    fact
	for k := 1; k <= 9; k++ {
		for x := 1; x <= 9; x++ {
			above := x*10 + k
			for y := 1; y <= 9; y++ {
				below := k*10 + y
				if above < below {
					if float64(x)/float64(y) == float64(above)/float64(below) {
						fmt.Println(x, y, above, below)
					}
				}
			}
		}
	}
}
