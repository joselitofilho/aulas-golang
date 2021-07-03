package main

import (
	"fmt"
	"time"
)

var name string

func main() {
	//   fat(0) = 1
	//   fat(n) = 1 * 2 * ... * (n-1)
	num := 5
	if num == 0 {
		fmt.Println("1")
	} else {
		var result int = 1
		for i := 1; i <= num; i++ {
			result *= i

		}
		fmt.Println(result)
	}

	// fibb(0) = 0
	// fibb(1) = 1
	// fibb(n) = fibb(n-1) + fibb(n-2) - if n > 1
	var a, b, c uint
	a = 0
	b = 1
	i := 2
	num = 6
	if num == 0 {
		fmt.Println(b)
	} else if num == 1 {
		fmt.Println(b)
	} else {
		if num == 2 {
			fmt.Println(b)
		} else {
			for i <= num {
				c = a + b
				a = b
				b = c
				i++
			}
			fmt.Println(c)
		}
	}

	// switch
	value := 10
	switch value {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 10:
		fmt.Println(10)
		fmt.Println(10)
		fmt.Println(10)
		fmt.Println(10)
	default:
		fmt.Println("default")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("manhã")
	case t.Hour() < 18:
		fmt.Println("tarde")
	default:
		fmt.Println("noite")
	}
}
