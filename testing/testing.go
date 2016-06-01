package testing

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

// MyTestingT ...
type MyTestingT struct {
	*testing.T
}

// CheckInt32WithMessage ...
func (t *MyTestingT) CheckInt32WithMessage(actual int32, expected int32, s string) {
	if actual != expected {
		t.Fail()
	}
}

// CheckInt32 ...
func (t *MyTestingT) CheckInt32(expected, actual int32) {
	t.CheckInt32WithMessage(actual, expected, "")
}

// Checkaa ...
func (t *MyTestingT) Checkaa(expected interface{}, actual interface{}) {
	typeActual := reflect.TypeOf(actual)
	typeExpected := reflect.TypeOf(expected)
	if typeActual == typeExpected {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected differs from actual")
		}
	} else {
		t.Errorf("Error in types between actual and expected. %s %s", typeActual, typeExpected)
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

// CheckVector ...
func (t *MyTestingT) CheckVector(expected, actual mgl64.Vec3, s string) {
	b := expected.ApproxEqual(actual)
	if !b {
		t.Fatalf("%s, Vectors are not equal. Expected: (%.14f, %.14f, %.14f), Actual: (%.14f, %.14f, %.14f)",
			s,
			expected[0],
			expected[1],
			expected[2],
			actual[0],
			actual[1],
			actual[2])

	}
	t.CheckFloat64(expected.X(), actual.X(), fmt.Sprintf("%s Vector's X differ", s))
	t.CheckFloat64(expected.Y(), actual.Y(), fmt.Sprintf("%s Vector's Y differ", s))
	t.CheckFloat64(expected.Z(), actual.Z(), fmt.Sprintf("%s Vector's Z differ", s))
}

// CheckErrorWithMessage ...
func (t *MyTestingT) CheckErrorWithMessage(err error, s string) {
	if err != nil {
		t.Fatalf("%s: Error Message: %s", s, err)
	}
}

// CheckIntWithMessage ...
func (t *MyTestingT) CheckIntWithMessage(actual int, expected int, s string) {
	if actual != expected {
		t.Fatalf("%s. Expected: %.d, Actual: %d", s, expected, actual)
	}
}
