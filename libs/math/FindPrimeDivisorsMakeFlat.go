package math

func FindPrimeDivisorsMakeFlat(data map[int]int) []int {
	result := make([]int, 0, 100)
	for k, v := range data {
		for index := uint64(1); index <= uint64(v); index++ {
			result = append(result, k)
		}
	}
	return result
}
