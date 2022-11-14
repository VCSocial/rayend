package Primitives

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) Point(s float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(s))
}

func (r Ray) Color() Color {
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)

	white := Vec3{1, 1, 1}.MulScalar(1.0 - t)
	blue := Vec3{0.5, 0.7, 1}.MulScalar(t)

	res := white.Add(blue)
	return Color{res.X, res.Y, res.Z}
}
