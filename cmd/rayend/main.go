package main

import (
	"fmt"
	"os"

	prm "github.com/vcsocial/rayend/pkg/primitives"
)

func render() {
	aspectRatio := 16.0 / 9.0
	w := 640
	h := int64(float64(w) / aspectRatio)
	fmt.Printf("P3\n%d %d\n255\n", w, h)

	vh := 2.0
	vw := aspectRatio * vh
	focalLen := 1.0

	origin := prm.Vec3{X: 0, Y: 0, Z: 0}
	horizontal := prm.Vec3{X: vw, Y: 0, Z: 0}
	vertical := prm.Vec3{X: 0, Y: vh, Z: 0}
	y := horizontal.DivScalar(2)
	z := vertical.DivScalar(2)

	lowerLeftCorner := origin.Sub(y).Sub(z).Sub(prm.Vec3{X: 0, Y: 0, Z: focalLen})

	for j := h - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d\n", j)
		for i := 0; i < w; i++ {
			u := float64(i) / (float64(w) - 1)
			v := float64(j) / (float64(h) - 1)

			a := horizontal.MulScalar(u)
			b := vertical.MulScalar(v)

			r := prm.Ray{
				Origin:    origin,
				Direction: lowerLeftCorner.Add(a).Add(b).Sub(origin),
			}
			writeColor(r.Color())
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}

func writeColor(v prm.Color) {
	const modifier = 255.999
	fmt.Printf("%d %d %d \n", int(modifier*v.X), int(modifier*v.Y), int(modifier*v.Z))
}

func main() {
	render()
}
