package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 8         // angle of x,y axes (30 deg)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if !math.IsNaN(ax) && !math.IsNaN(ay) && !math.IsNaN(bx) && !math.IsNaN(by) && !math.IsNaN(cx) && !math.IsNaN(cy) && !math.IsNaN(dx) && !math.IsNaN(dy) {
                            var stroke string
                            if az > 0 {
                                stroke = "stroke:red"
                            } else {
                                stroke = "stroke:blue"
                            }

				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='"+stroke+"' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2d canvas
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from 0,0
	return math.Sin(r) / r
}
