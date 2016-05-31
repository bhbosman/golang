package igprotocol

import (
	"fmt"
	"math"
	"testing"
	// "github.com/go-gl/mathgl/mgl64"
)

// MyTestingT ...
type MyTestingT struct {
	*testing.T
}


// CheckIntWithMessage ...
func (t *MyTestingT) CheckIntWithMessage(actual int, expected int, s string) {
	if actual != expected {
		t.Fatalf("%s. Expected: %.d, Actual: %d", s, expected, actual)
	}
}

// CheckErrorWithMessage ...
func (t *MyTestingT) CheckErrorWithMessage(err error, s string) {
	if err != nil {
		t.Fatalf("%s: Error Message: %s", s, err)
	}
}


// CheckBool ...
func (t *MyTestingT) CheckBool(b bool, s string) {
	if !b {
		t.Fatalf(s)
	}
}

// CheckFloat64 ...
func (t *MyTestingT) CheckFloat64(expected, actual float64, s string) {
	if !(math.Abs(expected-actual) < 1.0e-14) {
		errorMessage := fmt.Sprintf(
			"Error: %s. Expected: %.14f, Actual: %.14f",
			s,
			expected,
			actual)
		t.CheckBool(false, errorMessage)
	}
}
