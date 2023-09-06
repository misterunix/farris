package main

import (
	"fmt"
	"math"
	"math/rand"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

func main() {

	fmt.Println("Starting")

	width := 2048
	height := 2048

	t1 := -5.0
	t2 := 21.0
	t3 := 2.0
	t4 := 3.0

	//is := "(0 to 2pi)"
	//formulaX := fmt.Sprintf("math.Cos(i) + math.Cos(%f*%s)/2.0 + math.Sin(%f*%s)/3.0", v1, is, v2, is)
	//formulaY := fmt.Sprintf("math.Sin(i) + math.Sin(%f*%s)/2.0 + math.Cos(%f*%s)/3.0", v1, is, v2, is)

	w1 := float64(width) / 2.0
	h1 := float64(height) / 2.0
	w2 := w1 / 2.0
	h2 := h1 / 2.0

	img := gd.CreateTrueColor(int(width), int(height))
	defer img.Destroy()

	//tc := img.ColorAllocate(0, 0, 255)
	//check := img.StringFT(tc, "CallingCode-Regular.ttf", 18, 0, 10, 20, formulaX)
	//fmt.Println(check)
	hueOffset := rand.Float64() * 360.0

	img.Fill(0, 0, img.ColorAllocate(255, 255, 255))

	//fo := gd.GetFont(gd.FONTGIANT)

	//img.String(fo, 10, 10, "Fuck Hello World", img.ColorAllocate(0, 0, 0))

	//hh := img.StringFT(img.ColorAllocate(0, 0, 0), "./CallingCode-Regular.ttf", 12, 0, 10, 20, "Hello World")
	//for uu, jj := range hh {
	//	fmt.Println(uu, jj)
	//}

	for i := 0.0; i < 2.0*math.Pi; i += 0.00001 {
		x := math.Cos(i) + math.Cos(t1*i)/t3 + math.Sin(t2*i)/t4
		y := math.Sin(i) + math.Sin(t1*i)/t3 + math.Cos(t2*i)/t4

		hue := math.Mod(i*57.295779513+hueOffset, 360) // 0 to 360

		r, g, b := hsl.HSLtoRGB(hue, 1.0, 0.5)

		c := img.ColorAllocate(int(r), int(g), int(b))

		//img.Imagefilledrectangle(0, 0, 299, 99, c);
		img.SetPixel(int(x*w2+w1), int(y*h2+h1), c)
		img.FilledEllipse(int(x*w2+w1), int(y*h2+h1), 4, 4, c)
	}
	//x := math.Cos(t) + math.Cos(6.0*t)/2.0 + math.Sin(14.0*t)/3.0
	//y := math.Sin(t) + math.Sin(6.0*t)/2.0 + math.Cos(14.0*t)/3.0

	img.Png("out2.png")
}
