package mgl

import (
	"fmt"
	"math"
)

type Vector2f struct {
	X float64
	Y float64
}

type Vector4f struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Vector3f struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3f) ScalarMul(scalar float64) Vector3f {
	return Vector3f{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func (v Vector3f) ScalarDiv(scalar float64) Vector3f {
	if scalar == 0 {
		panic("cannot divide a vector by scalar 0")
	}

	inverse := 1.0 / scalar
	return v.ScalarMul(inverse)
}

func (v Vector3f) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3f) Neg() Vector3f {
	return Vector3f{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vector3f) Norm() Vector3f {
	return v.ScalarDiv(v.Mag())
}

func (v Vector3f) Add(other Vector3f) Vector3f {
	return Vector3f{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vector3f) Sub(other Vector3f) Vector3f {
	return Vector3f{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vector3f) Dot(b Vector3f) float64 {
	return (v.X*b.X + v.Y*b.Y + v.Z*b.Z)
}

func (v Vector3f) String() string {
	return fmt.Sprintf("v(%f, %f, %f)", v.X, v.Y, v.Z)
}
