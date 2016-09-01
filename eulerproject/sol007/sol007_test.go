package sol007

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=7
//

func TestSolution06(t *testing.T) {

	data := AtkinsSievePrime(104743)
	if len(data) >= 10001 {
		fmt.Println(data[10000])
	} else {
		fmt.Println(data)
	}
}
