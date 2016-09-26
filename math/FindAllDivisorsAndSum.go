package math

func FindAllDivisorsAndSum(number int, data []int) ([]int, int) {
	divisors := make(map[int]int)
	divisors[1] = 0

	checkAndAdd := func(value int, temp map[int]int) {
		if value != number {
			_, ok := divisors[int(value)]
			if !ok {
				temp[value] = 0
			}
		}
	}
	for _, v := range data {
		temp := make(map[int]int)

		for k := range divisors {
			checkAndAdd(k*v, temp)
		}
		checkAndAdd(v, temp)
		for k := range temp {
			divisors[k] = 0
		}
	}
	result := make([]int, 0, len(divisors))
	sum := 0
	for k := range divisors {
		result = append(result, k)
		sum += k
	}
	return result, sum
}
