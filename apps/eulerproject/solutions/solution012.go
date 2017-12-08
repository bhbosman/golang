package solutions

import bhb_math "github.com/bhbosman/code/golang/libs/math"

// FindTriangleWithXDivisiors2 ...
func FindTriangleWithXDivisiors2(number int, primes []uint64) uint64 {
	index := 2
	divisors := 0
	saved := uint64(0)
	for divisors < number {
		divisors = 1
		triangleNumber := uint64(index * (index + 1) / 2)
		saved = triangleNumber
		for _, p := range primes {
			if p*p < triangleNumber {
				if triangleNumber%p == 0 {
					count := 1
					for triangleNumber%p == 0 {
						count++
						triangleNumber /= p
					}
					divisors *= (count)
				}
			} else {
				divisors *= 2
				break
			}
		}
		index++
	}
	return saved
}

// FindTriangleWithXDivisiors3 ...
// Impl from the dicumentation
func FindTriangleWithXDivisiors3(number int, primes []uint64) uint64 {
	n := uint64(3) //start with a prime
	Dn := 2        //number of divisors for any prime
	cnt := 0       //to insure the while loop is entered
	var Dn1, exponent int
	var n1 uint64

	for cnt < number {
		n++
		n1 = n
		if n1%2 == 0 {
			n1 = n1 / 2
		}
		Dn1 = 1
		for i := 0; i < len(primes); i++ {

			if primes[i]*primes[i] > uint64(n1) {
				Dn1 = 2 * Dn1
				break
			}

			exponent = 1
			for uint64(n1)%primes[i] == 0 {
				exponent++
				n1 = n1 / primes[i]
			}
			if exponent > 1 {
				Dn1 = Dn1 * exponent
			}
			if n1 == 1 {
				break
			}
		}
		cnt = Dn * Dn1
		Dn = Dn1
	}
	return n * (n - 1) / 2
}

// FindTriangleWithXDivisiors ...
func FindTriangleWithXDivisiors(number int) uint64 {
	primes := bhb_math.AtkinsSievePrime(10474300)
	return FindTriangleWithXDivisiors2(number, primes)
}
