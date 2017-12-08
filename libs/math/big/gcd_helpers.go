package big

import (
	"encoding/binary"
	"math/big"
)

// UIn64ToInt helper
func UIn64ToInt(value uint64) *big.Int {
	buffer := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(buffer, value)
	result := &big.Int{}
	result.SetBytes(buffer)
	//
	return result
}

// CalculateGCD ...
func CalculateGCD(a, b uint64) *big.Int {
	bigIntA := UIn64ToInt(a)
	bigIntB := UIn64ToInt(b)
	//
	gcd := bigIntA.GCD(nil, nil, bigIntA, bigIntB)
	//
	return gcd
}

// CompareGCD ...
func CompareGCD(gcd *big.Int, b uint64) bool {
	bigIntB := UIn64ToInt(b)
	//
	return gcd.Cmp(bigIntB) == 0
}
