package solutions

import "testing"
import "fmt"
import "strconv"

//
// https://projecteuler.net/problem=30
//

func TestSolution30_01(t *testing.T) {

	number := 5
	highNumber := Pow(9, number) * (number + 1)
	createData := func() [10]int {
		data := [10]int{
			Pow(0, number),
			Pow(1, number),
			Pow(2, number),
			Pow(3, number),
			Pow(4, number),
			Pow(5, number),
			Pow(6, number),
			Pow(7, number),
			Pow(8, number),
			Pow(9, number)}
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
