package main

import (
	"fmt"
	"math/cmplx"
	"image/png"
	"image/color"
	"os"
	"image"
	"sync"
)

const IMAGE_HEIGHT int = 3240
const IMAGE_WIDTH int = 5760
const C_INCREMENT float64 = 0.0005
const GRAPH_RANGE float64 = 2.0
const C_SEED complex128 = 0.0099

type Point struct {
	X, Y float64
}

func iterateAndDraw(x_min float64, x_max float64, y_min float64, y_max float64, img *image.RGBA) {
	for x := x_min; x < x_max; x += C_INCREMENT {
		for y := y_min; y < y_max; y += C_INCREMENT {
			inf, colour := isTendToInf(C_SEED, complex(y, x))
			if inf {
				img.Set(IMAGE_WIDTH/2 + int(x/GRAPH_RANGE * float64(IMAGE_WIDTH/2)), 
					IMAGE_HEIGHT/2 + int(y/GRAPH_RANGE * float64(IMAGE_HEIGHT/2)), 
					colour)
			}
		}
	}
}

func main() {
	//tests()
	img := image.NewRGBA(image.Rect(0, 0, IMAGE_WIDTH, IMAGE_HEIGHT))
	var waitGroup sync.WaitGroup	
	waitGroup.Add(3)

	// 1st Quadrant
	go func() {
		defer waitGroup.Done()
		iterateAndDraw(0, GRAPH_RANGE, -1 * GRAPH_RANGE, 0, img)
	} ()

	// 2nd Quadrant
	go func() {
		defer waitGroup.Done()
		iterateAndDraw(-1 * GRAPH_RANGE, 0, -1 * GRAPH_RANGE, 0, img)
	} ()
	// 3rd Quadrant
	go func() {
		defer waitGroup.Done()
		iterateAndDraw(-1 * GRAPH_RANGE, 0, 0, GRAPH_RANGE, img)
	} ()

	// 4th Quadrant
	iterateAndDraw(0, GRAPH_RANGE, 0, GRAPH_RANGE, img)

	waitGroup.Wait()

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func computeNext(z complex128, c complex128) complex128 {
	return z * z + c
}

func isTendToInf(z complex128, c complex128) (bool, color.RGBA) {
	for i:= 0; i < 100; i++ {
		z = computeNext(z, c)
		if cmplx.IsInf(z) {
			return true, getColour(i)
		}
	}
	return false, color.RGBA{0, 0, 0, 255}
}

func getColour(iterationsInf int) color.RGBA {
	switch {
	case iterationsInf < 25:
		return color.RGBA{0, 0, 255, 255}
	case iterationsInf < 50:
		return color.RGBA{0, 255, 0, 255}
	case iterationsInf < 100:
		return color.RGBA{255, 0, 0, 255}
	default:
		return color.RGBA{0, 0, 255, 255}
	}
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
			result, _ := isTendToInf(0, v.test)
			fmt.Printf("| %v       |       %v |\n", v.test, result == v.result)
		}
		fmt.Println()
	}
