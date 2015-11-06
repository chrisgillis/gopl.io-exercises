package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
                    x := float64(px)/width*(xmax-xmin)+xmin
                    z := complex(x,y)
		    img.Set(px, py, color.RGBA{uint8(mandelbrot(z)),0,0,255})
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) float64 {
	const iterations = 255 
        const h = 1e-6
        const eps = 1e-3

        var n uint8 = 0
	for ; n < iterations; n++ {
                dz := (f(z + complex(h,h)) - f(z)) / complex(h,h)
                z0 := z - f(z) / dz
		if cmplx.Abs(z0 - z) < eps {
			return float64(n*15)
		}
                z = z0
	}
	return 0.0
}

func f(z complex128) complex128 {
    return (z*z*z/cmplx.Cos(z)) - 1.0
}

