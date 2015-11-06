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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
                    var total float64 
                    var z complex128
                    for i:= 0.0; i <= 3.0; i++ {
                        if i == 0 {
			    x := float64(px)/width*(xmax-xmin) + xmin
	                    z = complex(x, y)
                        } else {
                            y := (float64(py) + (i/4))/height*(ymax-ymin)+ymin
                            x := (float64(px) + (i/4))/width*(xmax-xmin)+xmin
                            z = complex(x,y)
                        }
                        total += mandelbrot(z)
                    }
                    total /= 4;

		    img.Set(px, py, color.RGBA{uint8(total),0,0,255})
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) float64 {
	const iterations = 50
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return float64(255 - contrast*n)
		}
	}
	return 0.0
}
