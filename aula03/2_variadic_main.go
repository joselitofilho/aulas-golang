package main

import "fmt"

func sum(numbers ...int) (total int) {
	total = 0
	for _, num := range numbers {
		total += num
	}
	return
}

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("sum = ", sum(numbers...))
	fmt.Println("sum = ", sum(1, 2, 3, 4, 5))
}
