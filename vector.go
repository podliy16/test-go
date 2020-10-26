package vector

import (
	"errors"
	"math"
)

type Vector []float64

var VectorsNotSameSize = errors.New("vectors are not the same size")
var ErrVectorInvalidDimension = errors.New("vectors dimensions are not of the expected size")

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Creates a vector of a specified size. All values are set to 0.
func New(size int) Vector {
	return make(Vector, size)
}

func NewFromValues(values []float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

func Add(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] + v2[i]
	}
	return result
}

func Subtract(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] - v2[i]
	}
	return result
}

func DotProduct(v1, v2 Vector) (float64, error) {
	if len(v1) != len(v2) {
		return 0.0, VectorsNotSameSize
	}

	length := len(v1)
	result := 0.0
	for i := 0; i < length; i++ {
		result += v1[i] * v2[i]
	}

	return result, nil
}

// Vectors dimensionality has to be 3.
func CrossProduct(v1, v2 Vector) (Vector, error) {
	if len(v1) != 3 || len(v2) != 3 {
		return nil, ErrVectorInvalidDimension
	}

	result := make(Vector, 3)
	result[0] = v1[1]*v2[2] - v1[2]*v2[1]
	result[1] = v1[2]*v2[0] - v1[0]*v2[2]
	result[2] = v1[0]*v2[1] - v1[1]*v2[0]

	return result, nil
}

func Scale(v Vector, value float64) Vector {
	length := len(v)
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v[i] * value
	}
	return result
}

func Magnitude(v Vector) float64 {
	result := 0.0
	for _, e := range v {
		result += e * e
	}
	return math.Sqrt(result)
}

func (v Vector) Clone() Vector {
	return NewFromValues(v)
}
