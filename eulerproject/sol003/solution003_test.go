package sol003

import "testing"

//
// https://projecteuler.net/problem=3
//
func TestSolution03_PrimesUnder_10(t *testing.T) {
	var Number uint64 = 10

	data := []uint64{2, 3, 5, 7}

	primes := AtkinsSievePrime(Number)
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
	var Number uint64 = 100

	data := []uint64{
		2, 3, 5, 7, 11, 13, 17,
		19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67,
		71, 73, 79, 83, 89, 97}

	primes := AtkinsSievePrime(Number)
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
	var Number uint64 = 13195

	data := []uint64{
		5, 7, 13, 29}
	primes := AtkinsSievePrime(Number)
	factors := make([]uint64, 0, len(primes))
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

func TestSolution03_FindLargestPrimeFactor_For_600851475143(t *testing.T) {

	d := LargestPrimeFactor(600851475143)

	if d != 6857 {
		t.Fatal("Prime count does not match")
	}
}
