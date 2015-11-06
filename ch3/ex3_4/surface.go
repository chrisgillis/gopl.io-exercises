package main

import (
	"fmt"
	"math"
        "net/http"
        "log"
        "os"
        "strconv"
)

var width, height = 600.0, 320.0
var xyrange = 30.0
var cells = 100
var xyscale = width/2.0/xyrange
var zscale = height * 0.4
var angle = math.Pi / 8
var strokeHi = "stroke:red"

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

    handler := func(w http.ResponseWriter,r *http.Request) {
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "Err: %v", err)
            os.Exit(1)
        }

        for k,v := range r.Form {
            switch k {
                case "width":
                    width,_ = strconv.ParseFloat(v[0],64)
                    xyscale = width/2.0/xyrange
                case "height":
                    height,_ = strconv.ParseFloat(v[0],64)
                    zscale = height * 0.4
                case "stroke":
                    strokeHi = "stroke:"+v[0]
            }
        }

        w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", int(width), int(height))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if !math.IsNaN(ax) && !math.IsNaN(ay) && !math.IsNaN(bx) && !math.IsNaN(by) && !math.IsNaN(cx) && !math.IsNaN(cy) && !math.IsNaN(dx) && !math.IsNaN(dy) {
                            var stroke string
                            if az > 0 {
                                stroke = strokeHi 
                            } else {
                                stroke = "stroke:blue"
                            }

				fmt.Fprintf(w,"<polygon points='%g,%g %g,%g %g,%g %g,%g' style='"+stroke+"' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(w,"</svg>")
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

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
