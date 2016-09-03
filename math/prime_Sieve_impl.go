package math

import (
	"math"

	"github.com/Workiva/go-datastructures/bitarray"
)

type bitArray struct {
	data bitarray.BitArray
}

func newBitArray(number uint64) bitArray {
	return bitArray{
		data: bitarray.NewBitArray(number)}
}

func (ba *bitArray) flip(index uint64) {
	b, _ := ba.data.GetBit(index)
	if b {
		ba.data.ClearBit(index)
	} else {
		ba.data.SetBit(index)
	}
}

func (ba *bitArray) setBit(k uint64) error {
	return ba.data.SetBit(k)
}

func (ba *bitArray) getBit(k uint64) (bool, error) {
	return ba.data.GetBit(k)
}

func (ba *bitArray) clearBit(k uint64) error {
	return ba.data.ClearBit(k)
}

func (ba *bitArray) toNums() []uint64 {
	return ba.data.ToNums()
}

// AtkinsSievePrime ...
func AtkinsSievePrime(number uint64) []uint64 {

	isPrime := newBitArray(uint64(number))
	isPrime.setBit(2)
	isPrime.setBit(3)

	nsqrt := uint64(math.Sqrt(float64(number)))
	var x uint64
	for x = 1; x <= nsqrt; x++ {
		// fmt.Println(x)
		var y uint64
		for y = 1; y <= nsqrt; y++ {
			n := (4*(x*x) + y*y)
			if n <= number && (n%12 == 1 || n%12 == 5) {
				isPrime.flip(n)
			}
			n = (3*(x*x) + y*y)
			if n <= number && n%12 == 7 {
				isPrime.flip(n)
			}
			n = (3*(x*x) - y*y)
			if x > y && n <= number && n%12 == 11 {
				isPrime.flip(n)
			}
		}
	}

	var i uint64
	for i = 5; i <= nsqrt; i++ {
		val, _ := isPrime.getBit(uint64(i))
		if val {
			var y uint64
			for y = i * i; y < number; y += i * i {
				isPrime.clearBit(y)
			}
		}
	}

	return isPrime.toNums()
}
