package math

// FindAllDivisorsAndSum ...
// Param:
//		data: lis for all the primes that make makes up a number
func FindAllDivisorsAndSum(data []int) []int {
	divisors := make(map[int]int)
	divisors[1] = 0
	for _, v := range data {
		temp := make(map[int]int)
		checkAndAdd := func(value int) {
			_, ok := divisors[int(value)]
			if !ok {
				temp[value] = 0
			}
		}

		for k := range divisors {
			checkAndAdd(k * v)
		}
		checkAndAdd(v)
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
	return result
}
