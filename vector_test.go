package vector

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestNew(t *testing.T) {
	v := New(4)
	v[0] = 1.0
	v[1] = 2.0
	v[2] = 3.0

	assertEqual(t, v[0], 1.0)
	assertEqual(t, v[1], 2.0)
	assertEqual(t, v[2], 3.0)
	assertEqual(t, v[3], 0.0)
}

func TestNewWithValues(t *testing.T) {
	v := NewFromValues([]float64{1.0, 2.0, 3.0})
	assertEqual(t, 1.0, v[0] )
	assertEqual(t, 2.0, v[1])
	assertEqual(t, 3.0, v[2])
}

func TestAdd(t *testing.T) {
	v1 := NewFromValues([]float64{1.0, 2.0, 3.0, 4.0})
	v2 := NewFromValues([]float64{1.0, 2.0, 3.0, 4.0})
	result := Add(v1, v2)
	assertEqual(t, 2.0, result[0])
	assertEqual(t, 4.0, result[1])
	assertEqual(t, 6.0, result[2])
	assertEqual(t, 8.0, result[3])
}

func TestSubtract(t *testing.T) {
	v1 := NewFromValues([]float64{1.0, 2.0, 3.0, 4.0})
	v2 := NewFromValues([]float64{1.0, 2.0, 3.0, 4.0})
	result := Subtract(v1, v2)
	assertEqual(t, 0.0, result[0])
	assertEqual(t, 0.0, result[1])
	assertEqual(t, 0.0, result[2])
	assertEqual(t, 0.0, result[3])
}

func TestDot(t *testing.T) {
	v1 := NewFromValues([]float64{1.0, 2.0, 3.0})
	v2 := NewFromValues([]float64{1.0, 2.0, 3.0})
	result, err := DotProduct(v1, v2)
	assertEqual(t, 14.0, result)
	if err != nil {
		t.Fatal("err is not nil")
	}
}

func TestCross(t *testing.T) {
	v1 := NewFromValues([]float64{0.0, 1.0, 2.0})
	v2 := NewFromValues([]float64{1.0, 2.0, 3.0})
	result, err := CrossProduct(v1, v2)
	assertEqual(t, -1.0, result[0])
	assertEqual(t, 2.0, result[1])
	assertEqual(t, -1.0, result[2])
	if err != nil {
		t.Fatal("err is not nil")
	}
}

func TestScale(t *testing.T) {
	v := NewFromValues([]float64{1.0, 2.0, 3.0})
	result := Scale(v, 3)
	assertEqual(t, 3.0, result[0])
	assertEqual(t, 6.0, result[1])
	assertEqual(t, 9.0, result[2])
}

func TestMagnitude(t *testing.T) {
	v := NewFromValues([]float64{3.0, 4.0})
	result := Magnitude(v)
	assertEqual(t, 5.0, result)
}