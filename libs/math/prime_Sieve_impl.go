package math

import (
	"math"

	"github.com/Workiva/go-datastructures/bitarray"
)

type bitArray struct {
	bitarray.BitArray
}

func flipBit(ba bitarray.BitArray, index uint64) {
	b, _ := ba.GetBit(index)
	if b {
		ba.ClearBit(index)
	} else {
		ba.SetBit(index)
	}
}

// AtkinsSievePrime ...
func AtkinsSievePrime(number uint64) []uint64 {
	isPrime := bitarray.NewBitArray(uint64(number))
	isPrime.SetBit(2)
	isPrime.SetBit(3)

	nsqrt := uint64(math.Sqrt(float64(number)))
	var x uint64
	for x = 1; x <= nsqrt; x++ {
		var y uint64
		for y = 1; y <= nsqrt; y++ {
			n := (4*(x*x) + y*y)
			if n <= number && (n%12 == 1 || n%12 == 5) {
				flipBit(isPrime, n)
			}
			n = (3*(x*x) + y*y)
			if n <= number && n%12 == 7 {
				flipBit(isPrime, n)
			}
			n = (3*(x*x) - y*y)
			if x > y && n <= number && n%12 == 11 {
				flipBit(isPrime, n)
			}
		}
	}

	var i uint64
	for i = 5; i <= nsqrt; i++ {
		val, _ := isPrime.GetBit(uint64(i))
		if val {
			var y uint64
			for y = i * i; y < number; y += i * i {
				isPrime.ClearBit(y)
			}
		}
	}
	return isPrime.ToNums()
}
