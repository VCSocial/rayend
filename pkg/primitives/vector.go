package Primitives

import (
	"math"
)

type Vector interface {
	Sub(v Vector)
	Add(v Vector)
	Mul(v Vector)
	Div(v Vector)
	LenSqr(v Vector) float64
	Len(v Vector) float64
	ScaleDiv(s float64)
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Point = Vec3
type Color = Vec3

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3) Mul(b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vec3) Div(b Vec3) Vec3 {
	return Vec3{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func (a Vec3) Dot(b Vec3) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func (v Vec3) LenSqr() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSqr())
}

func (v Vec3) AddScalar(s float64) Vec3 {
	return Vec3{v.X + s, v.Y + s, v.Z + s}
}

func (v Vec3) SubScalar(s float64) Vec3 {
	return Vec3{v.X - s, v.Y - s, v.Z - s}
}

func (v Vec3) ScaleAdd(s float64) Vec3 {
	return Vec3{v.X + s, v.Y + s, v.Z + s}
}

func (v Vec3) MulScalar(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) DivScalar(s float64) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

func (v Vec3) Unit() Vec3 {
	t := 1.0 / v.Len()
	return v.MulScalar(t)
}
