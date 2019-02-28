// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	hx := .4 / width * (xmax - ymin)
	hy := .4 / height * (ymax - ymin)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			cs := []uint8{mandelbrot(complex(x+hx, y)),
				mandelbrot(complex(x, y+hy)),
				mandelbrot(complex(x-hx, y)),
				mandelbrot(complex(x, y+hy)),
				mandelbrot(complex(x, y))}
			c := float64(cs[0])
			for i := 1; i < 5; i++ {
				c = math.Sqrt(c * float64(cs[i]))
			}
			img.Set(px, py, color.Gray{uint8(c)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
}

//!-
