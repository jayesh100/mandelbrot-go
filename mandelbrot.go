package main

import (
	//"fmt"
	"math/cmplx"
	"image/png"
	"image/color"
	"os"
	"image"
)

const IMAGE_HEIGHT int = 1080
const IMAGE_WIDTH int = 1920


type Point struct {
	X, Y float64
}

func main() {
	var grid []Point = make([]Point, 1, 1)
	img := image.NewRGBA(image.Rect(0, 0, IMAGE_WIDTH, IMAGE_HEIGHT))

	for x := -2.00; x < 2; x += 0.001 {
		for y := -2.00; y < 2; y += 0.001 {
			if isTendToInf(0, complex(y, x)){
				grid = append(grid, Point{x, y})
			}
		}
	}	

	for _, i := range grid {
		//fmt.Printf("Drawing point @ { %v, %v }", int(i.X/2 * 1920),int(i.Y/2 * 1080))
		img.Set(int(i.X/2 * float64(IMAGE_WIDTH)), int(i.Y/2 * float64(IMAGE_HEIGHT)),color.RGBA{0, 0, 255, 255})
	}
	//img.Set(2, 3, color.RGBA{0, 0, 255, 255})

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)
}

func computeNext(z complex128, c complex128) complex128{
	return z * z + c
}

func isTendToInf(z complex128, c complex128) bool {
	for i:= 0; i < 1000; i++ {
		z = computeNext(z, c)
		if cmplx.IsInf(z) {
			return true
		}
	}
	return false
}