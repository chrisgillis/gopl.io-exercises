package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var zoom float64 = 1.0
var xmin, ymin, xmax, ymax float64 = -zoom, -zoom, +zoom, +zoom

func main() {
	const (
		width, height = 1024, 1024
	)

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Err: %v", err)
			return
		}

		for k, v := range r.Form {
			if k == "zoom" {
				zoom, _ = strconv.ParseFloat(v[0], 64)
				xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom

			}
		}

		w.Header().Set("Content-Type", "image/png")

		img := image.NewRGBA(image.Rect(0, 0, width, height))

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, color.RGBA{uint8(mandelbrot(z)), 0, 0, 255})
			}
		}
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) float64 {
	const iterations = 255
	const h = 1e-6
	const eps = 1e-3

	var n uint8 = 0
	for ; n < iterations; n++ {
		dz := (f(z+complex(h, h)) - f(z)) / complex(h, h)
		z0 := z - f(z)/dz
		if cmplx.Abs(z0-z) < eps {
			return float64(n * 15)
		}
		z = z0
	}
	return 0.0
}

func f(z complex128) complex128 {
	return (z * z * z / cmplx.Cos(z)) - 1.0
}
