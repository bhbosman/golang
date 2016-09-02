package sol009

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=9
//

func TestSolution06_01(t *testing.T) {
	sumOfLength := 1000
	minLenghtOfA := int((sumOfLength - 3) / 3)
	finished := false

	beginValue := 1
	for aa, a, jumpaa := beginValue*beginValue, beginValue, 2*beginValue+1; !finished && a <= minLenghtOfA; aa, a, jumpaa = aa+jumpaa, a+1, jumpaa+2 {
		for bb, b, jumpbb := aa+jumpaa, a+1, 2*(a+1)+1; !finished && a+b+b+1 <= sumOfLength; bb, b, jumpbb = bb+jumpbb, b+1, jumpbb+2 {
			c := sumOfLength - b - a
			cc := c * c
			if aa+bb == cc {
				// finished = true
				fmt.Println(a, b, c, a*b*c)
			}
		}
	}
}
