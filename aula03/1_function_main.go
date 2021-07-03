package main

import "fmt"

func rectangleArea(width, height float64) (result float64) {
	result = width * height
	return
}

func twoResults() (result1 float64, result2 float64) {
	result1 = 1.1
	result2 = 2.2
	return
}

func main() {
	fmt.Printf("rectangleArea = %.2f", rectangleArea(2, 3))
	// fmt.Println("rectangleArea =", rectangleArea(2, 3))

	res1, res2 := twoResults()
	fmt.Println("twoResult =", res1, res2)
}
