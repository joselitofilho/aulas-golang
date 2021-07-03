package main

import "fmt"

func init() {
	fmt.Println("Executed First")
}

func init() {
	fmt.Println("Me too! In order")
}

func main() {
	fmt.Println("Hello World")
}
