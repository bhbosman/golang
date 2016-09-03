package math

import "sort"

func MinumumDivisibleNumber(numbers []int) int {
	sort.Sort(sort.IntSlice(numbers))
	result := 1
	for index := 0; index < len(numbers); index++ {
		v := numbers[index]
		result = result * v
		for k := index; k < len(numbers); k++ {
			if numbers[k]%v == 0 {
				numbers[k] = numbers[k] / v
			}
		}
	}
	return result
}
