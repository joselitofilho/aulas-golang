package main

import (
	"fmt"
	"time"
)

func worker(done chan int) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- 1234
}

func main() {
	done := make(chan int, 1)
	go worker(done)
	<-done
}
