package main

import "fmt"

func fibonacci(channel, exit chan int) {
	x, y := 0, 1
	for {
		select {
		case channel <- x:
			x, y = y, x+y
		case <-exit:
			fmt.Println("exit")
			return
		}
	}
}

func main() {
	channel := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			fmt.Println(<-channel)
		}
		exit <- 0
	}()
	fibonacci(channel, exit)
}
