package main

import "fmt"

func sequency() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := sequency()
	fmt.Println("    newInt =", nextInt())
	fmt.Println("    newInt =", nextInt())
	fmt.Println("    newInt =", nextInt())

	newInteger := sequency()
	fmt.Println("newInteger =", newInteger())

	fmt.Println("    newInt =", nextInt())

	fmt.Println("  sequency =", sequency()())
	fmt.Println("  sequency =", sequency()())
	fmt.Println("  sequency =", sequency()())
}
