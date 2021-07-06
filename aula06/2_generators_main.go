package main

import "fmt"

func sequency() <-chan int {
	channel := make(chan int)
	go func() {
		i := 0
		for {
			channel <- i
			i++
		}
	}()
	return channel
}

func main() {
	channel := sequency()
	for i := 0; i < 100; i++ {
		fmt.Println(<-channel)
	}
}
