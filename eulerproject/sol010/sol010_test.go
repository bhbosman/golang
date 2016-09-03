package sol010

import (
	"fmt"
	"math"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=9
//

//
// My solution
//
func TestSolution10_01(t *testing.T) {
	number := uint64(2000000)

	data := bhb_math.AtkinsSievePrime(number)

	b := math.MaxUint64/uint64(len(data)) > number
	mul := uint64(0)
	if b {
		for _, v := range data {
			mul += v
		}
	} else {
		t.Fatal("need big number calc")
	}
	fmt.Println(mul)

}
