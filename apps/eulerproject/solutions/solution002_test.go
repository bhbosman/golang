package solutions

import (
	"fmt"
	"testing"
)

// Fibonacci ...
func Fibonacci(number int) int {
	switch number {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}

const MaxEvenValue uint64 = 4000000 * 1

//
// https://projecteuler.net/problem=1
//
func TestSolution02_01(t *testing.T) {
	sum := uint64(0)
	index := 1
	for {
		fib := uint64(Fibonacci(index))
		if fib <= MaxEvenValue {
			if fib%2 == 0 {
				sum += fib
			}
			index++
			continue
		}
		break
	}

	fmt.Println(sum)
	if sum != 4613732 {
		t.Fatal("")
	}
}

func TestSolution02_02(t *testing.T) {
	sum := uint64(2)
	val01 := uint64(3)
	val02 := uint64(5)

	for true {
		val03 := uint64(val01 + val02)
		// fmt.Print(val03, " ")
		if val03 > MaxEvenValue {
			break
		}
		sum += val03

		val01 = val03 + val02
		val02 = val03 + val01
	}

	fmt.Println(sum)
	if sum != 4613732 {
		t.Fatal("")
	}
}
