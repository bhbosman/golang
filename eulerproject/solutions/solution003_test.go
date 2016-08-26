package solutions

import "testing"

// https://projecteuler.net/problem=3
func TestSolution03_PrimesUnder_10(t *testing.T) {
	// Number := 600851475143
	Number := 10

	data := []int{2, 3, 5, 7}
	primes := AtkinsSievePrimeBool(Number)
	if len(primes) != len(data) {
		t.Fatal("Prime count does not match")
	}

	for i := range primes {
		if primes[i] != data[i] {
			t.Fatal("index value not match")
		}
	}
}

func TestSolution03_PrimesUnder_100(t *testing.T) {
	// Number := 600851475143
	Number := 100

	data := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	primes := AtkinsSievePrimeBool(Number)
	if len(primes) != len(data) {
		t.Fatal("Prime count does not match")
	}
	for i := range primes {
		if primes[i] != data[i] {
			t.Fatal("index value not match")
		}
	}
}

func TestSolution03_PrimesFactors_For_13195(t *testing.T) {
	// Number := 600851475143
	Number := 13195

	data := []int{5, 7, 13, 29}
	primes := AtkinsSievePrimeBool(Number)
	factors := make([]int, 0, len(primes))
	for _, value := range primes {
		if Number%value == 0 {
			factors = append(factors, value)
		}
	}

	if len(factors) != len(data) {
		t.Fatal("Prime count does not match")
	}

	for i, value := range factors {
		if value != data[i] {
			t.Fatalf("index value not match. index: %d. Values: %d, %d", i, value, data[i])
		}
	}
}

// func TestSolution04_PrimesFactors_For_13195(t *testing.T) {
// 	// Number := 600851475143
// 	Number := 13195

// 	data := []int{5, 7, 13, 29}
// 	primes := AtkinsSievePrime_Bit(Number)
// 	factors := make([]int, 0, len(primes))
// 	for _, value := range primes {
// 		if Number%value == 0 {
// 			factors = append(factors, value)
// 		}
// 	}

// 	if len(factors) != len(data) {
// 		t.Fatal("Prime count does not match")
// 	}

// 	for i, value := range factors {
// 		if value != data[i] {
// 			t.Fatalf("index value not match. index: %d. Values: %d, %d", i, value, data[i])
// 		}
// 	}
// }
