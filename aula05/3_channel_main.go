package main

import "fmt"

func sum(values []int, c chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	// ch <- value // Inserir
	// <-ch        // Obter

	// messages := make(chan string)
	// var value int
	// go func(msg chan string, v *int) {
	// 	time.Sleep(3 * time.Second)
	// 	msg <- "buffered"
	// 	*v = 10
	// }(messages, &value)
	// fmt.Println(<-messages, value)

	// --------------------------------------------------------------------------------

	values := []int{10, 20, 30, -20, 0, 10}

	c := make(chan int)
	go sum(values[len(values)/2:], c)
	go sum(values[:len(values)/2], c)
	x, y := <-c, <-c // receive from channel c

	fmt.Println(x, y, x+y)
}
