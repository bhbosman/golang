package solutions

import (
	"math"
)

// Storage ...
type Storage interface {
	Flip(index int)
	SetData(index int, value bool)
	GetData(index int) bool
	GetCount() int
	FillInPrimes() []int
}

func atkinsSievePrimeInternal(number int, storage Storage) {
	storage.SetData(2, true)
	storage.SetData(3, true)

	nsqrt := int(math.Sqrt(float64(number)))
	for x := 1; x <= nsqrt; x++ {
		for y := 1; y <= nsqrt; y++ {
			n := 4*(x*x) + y*y
			if n <= number && (n%12 == 1 || n%12 == 5) {
				storage.Flip(n)
			}
			n = 3*(x*x) + y*y
			if n <= number && n%12 == 7 {
				storage.Flip(n)
			}
			n = 3*(x*x) - y*y
			if x > y && n <= number && n%12 == 11 {
				storage.Flip(n)
			}
		}
	}

	for i := 5; i <= nsqrt; i++ {
		if storage.GetData(i) {
			for y := i * i; y < number; y += i * i {
				storage.SetData(y, false)
			}
		}
	}
}

// AtkinsSievePrimeBool ...
func AtkinsSievePrimeBool(number int) []int {
	isprime := &StorageBoolImpl{
		count: 0,
		data:  make([]bool, number, number)}
	atkinsSievePrimeInternal(number, isprime)
	result := isprime.FillInPrimes()
	return result
}

// func AtkinsSievePrime_Bit(number int) []int {
// 	isprime := NewStorageBoolImpl(number)
// 	AtkinsSievePrimeInternal(number, isprime)
// 	result := isprime.FillInPrimes()
// 	return result
// }
