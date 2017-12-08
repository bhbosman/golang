package solutions

import "testing"
import "fmt"
import "strconv"

import (
	bhb_math "github.com/bhbosman/code/golang/libs/math"
)

//
// https://projecteuler.net/problem=30
//

func TestSolution30_01(t *testing.T) {

	number := 5
	highNumber := bhb_math.Pow(9, number) * (number + 1)
	createData := func() [10]int {
		data := [10]int{
			bhb_math.Pow(0, number),
			bhb_math.Pow(1, number),
			bhb_math.Pow(2, number),
			bhb_math.Pow(3, number),
			bhb_math.Pow(4, number),
			bhb_math.Pow(5, number),
			bhb_math.Pow(6, number),
			bhb_math.Pow(7, number),
			bhb_math.Pow(8, number),
			bhb_math.Pow(9, number)}
		return data
	}
	data := createData()
	fmt.Println(data, highNumber)
	grandTotal := 0
	for i := 2; i <= highNumber; i++ {
		s := strconv.Itoa(i)
		total := 0
		for _, c := range s {
			total += data[c-'0']
		}
		if total == i {
			fmt.Println(total)
			grandTotal += total
		}
	}
	fmt.Println(grandTotal)
}
