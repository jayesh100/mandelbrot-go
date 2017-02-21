package main

import (
	"fmt"
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
	tests()
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
		img.Set(IMAGE_WIDTH/2 + int(i.X/2 * float64(IMAGE_WIDTH)), IMAGE_HEIGHT/2 + int(i.Y/2 * float64(IMAGE_HEIGHT)),color.RGBA{0, 0, 255, 255})
	}

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

// TESTING CODE

type testResult struct {
	test complex128
	result bool
}

func tests() {
	//Test cases
	infTest := [5]testResult{ 
		testResult{ 0 + 0i, false}, 
		testResult{ 1 + 0i, true}, 
		testResult{ 0 + 1i, false},
		testResult{ 0 + 2i, true}, 
		testResult{ 1 + 2i, true}}
	fmt.Println()
	fmt.Println("Starting tests....")
	fmt.Printf("Testing \"tends to infinity\" tests [ %v Tests ]\n", len(infTest))
	fmt.Println("-----------------------------")
	fmt.Println("| TEST CASE    |     PASSED |")
	fmt.Println("-----------------------------")
	for _, v := range infTest {
		fmt.Printf("| %v       |       %v |\n", v.test, isTendToInf(0, v.test) == v.result)
	}
	fmt.Println()
	
}