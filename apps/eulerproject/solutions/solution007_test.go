package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/code/golang/libs/math"
)

//
// https://projecteuler.net/problem=7
//

func TestSolution06(t *testing.T) {

	data := bhb_math.AtkinsSievePrime(104743)
	if len(data) >= 10001 {
		fmt.Println(data[10000])
	} else {
		fmt.Println(data)
	}
}
