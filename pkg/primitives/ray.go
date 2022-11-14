package Primitives

import "math"

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(s float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(s))
}

func (r Ray) Color() Color {
	y := r.HitSphere(Point{0, 0, -1}, 0.5)
	if y > 0 {
		n := r.At(y).Sub(Vec3{0, 0, -1}).Unit()
		res := Vec3{n.X + 1, n.Y + 1, n.Z + 1}.MulScalar(0.5)
		return Color{res.X, res.Y, res.Z}
	}

	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)

	white := Vec3{1, 1, 1}.MulScalar(1.0 - t)
	blue := Vec3{0.5, 0.7, 1}.MulScalar(t)

	return white.Add(blue)
}

func (r Ray) HitSphere(center Point, radius float64) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := oc.MulScalar(2.0).Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}
