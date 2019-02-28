// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320 // canvas size in pixels
	color         = "ff0000"
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func setOptions(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}

	if widthP, found := r.Form["width"]; found && len(widthP) != 0 {
		if w, err := strconv.Atoi(widthP[0]); err == nil {
			width = w
			xyscale = float64(width) / 2 / xyrange
			zscale = float64(height) * 0.4
		}
	}

	if heightP, found := r.Form["height"]; found && len(heightP) != 0 {
		if h, err := strconv.Atoi(heightP[0]); err == nil {
			height = h
			xyscale = float64(width) / 2 / xyrange
			zscale = float64(height) * 0.4
		}
	}

	if c, found := r.Form["color"]; found && len(c) != 0 {
		color = c[0]
	}

}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		setOptions(r)
		w.Header().Set("Content-Type", "image/svg+xml")
		plot(w)
	}

	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if ok1 && ok2 && ok3 && ok4 {
				fmt.Fprintf(out, "<polygon style='fill:#%s' "+
					"points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color,
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := saddle(x, y)

	if math.IsNaN(z) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
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
