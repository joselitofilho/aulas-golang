package main

import "fmt"

var name string

func main() {
	//   fat(0) = 1
	//   fat(n) = 1 * 2 * ... * (n-1)
	num := 5
	if num == 0 {
		fmt.Println("1")
		return
	} else {
		var result int = 1
		for i := 1; i <= num; i++ {
			result = result * i
		}
		fmt.Println(result)
		return
	}

	// fibb(0) = 0
	// fibb(1) = 1
	// fibb(n) = fibb(n-1) + fibb(n-2) - if n > 1
}
