// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zmax := 0.0
	zmin := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := corner(i, j)
			if !math.IsNaN(z) {
				zmax = math.Max(zmax, z)
				zmin = math.Min(zmin, z)
			}
		}
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1 := corner(i+1, j)
			bx, by, z2 := corner(i, j)
			cx, cy, z3 := corner(i, j+1)
			dx, dy, z4 := corner(i+1, j+1)
			fmt.Printf("<polygon style='fill:%s' "+
				"points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				zcolor([4]float64{z1, z2, z3, z4}, zmin, zmax),
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func zcolor(zs [4]float64, zmin, zmax float64) string {
	n := 0.0
	sum := 0.0
	for _, v := range zs {
		if !math.IsNaN(v) {
			n++
			sum += v
		}
	}
	z := sum / n
	z = ((z-zmin)/(zmax-zmin))*2 - 1
	var r, b uint8
	r = uint8(z * 0xff)
	b = uint8(math.Abs(z * 0xff))
	if z > 0 {
		b = 0
	} else {
		r = 0
	}
	return fmt.Sprintf("#%02x%02x%02x", r, 0, b)
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) * 0.1
}

func saddle(x, y float64) float64 {
	return (math.Pow(x, 2) - math.Pow(y, 2)) * 3e-3
}

//!-
