package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	// go f("goroutine")
	// f("direct")

	time.Sleep(time.Second)
	fmt.Println("done")
}
