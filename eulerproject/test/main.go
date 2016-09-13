package main

import (
	"fmt"
	"testing"

	"github.com/bhbosman/golang/eulerproject/solutions"
	bhb_math "github.com/bhbosman/golang/math"
)

func dddddSolution1202(f func(number int, primes []uint64) uint64, divisions int, triangle uint64, primes []uint64) func(t *testing.B) {
	return func(ff *testing.B) {
		for i := 1; i <= ff.N; i++ {
			ans := f(divisions, primes)
			if i == 1 {
				fmt.Println(ans)
			}
		}
	}
}

func main() {
	primes := bhb_math.AtkinsSievePrime(5000000)
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors2, 50, 76576500, primes)))
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors3, 50, 76576500, primes)))
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors2, 500, 76576500, primes)))
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors3, 500, 76576500, primes)))
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors2, 5000, 76576500, primes)))
	fmt.Println(testing.Benchmark(dddddSolution1202(solutions.FindTriangleWithXDivisiors3, 5000, 76576500, primes)))

}
