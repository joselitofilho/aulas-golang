package main

import "fmt"

func main() {
	area := func(width, height float32) float64 {
		return float64(width) * float64(height)
	}
	fmt.Println(area(2, 3))

	areaResult := func(width, height float64) float64 {
		return width * height
	}(2, 3)
	fmt.Println(areaResult)
}
