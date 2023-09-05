package main

import (
	"fmt"
	"math"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

func main() {

	fmt.Println("Starting")

	width := 512
	height := 512

	v1 := 3.0
	v2 := 13.0

	is := "(0 to 2pi)"
	formulaX := fmt.Sprintf("math.Cos(i) + math.Cos(%f*%s)/2.0 + math.Sin(%f*%s)/3.0", v1, is, v2, is)
	//formulaY := fmt.Sprintf("math.Sin(i) + math.Sin(%f*%s)/2.0 + math.Cos(%f*%s)/3.0", v1, is, v2, is)

	w1 := float64(width) / 2.0
	h1 := float64(height) / 2.0
	w2 := w1 / 2.0
	h2 := h1 / 2.0

	img := gd.CreateTrueColor(int(width), int(height))
	defer img.Destroy()

	tc := img.ColorAllocate(0, 0, 255)
	check := img.StringFT(tc, "CallingCode-Regular.ttf", 18, 0, 10, 20, formulaX)
	fmt.Println(check)

	img.Fill(0, 0, img.ColorAllocate(255, 255, 255))

	for i := 0.0; i < 2.0*math.Pi; i += 0.00001 {
		x := math.Cos(i) + math.Cos(v1*i)/2.0 + math.Sin(v2*i)/3.0
		y := math.Sin(i) + math.Sin(v1*i)/2.0 + math.Cos(v2*i)/3.0

		r, g, b := hsl.HSLtoRGB(i*57.295779513, 1.0, 0.5)
		c := img.ColorAllocate(int(r), int(g), int(b))
		img.SetPixel(int(x*w2+w1), int(y*h2+h1), c)
	}
	//x := math.Cos(t) + math.Cos(6.0*t)/2.0 + math.Sin(14.0*t)/3.0
	//y := math.Sin(t) + math.Sin(6.0*t)/2.0 + math.Cos(14.0*t)/3.0

	img.Png("out.png")
}
