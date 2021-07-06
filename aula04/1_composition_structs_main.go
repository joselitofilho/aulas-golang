package main

import "fmt"

type details struct {
	genre string
}

type game struct {
	name  string
	price float64
	details
}

func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
}

func (g game) show() {
	fmt.Println(" Name:", g.name)
	fmt.Println("Price:", g.price)
	g.showDetails()
}

func main() {
	newGame := game{"Counter Strike", 149.99, details{"18+"}}
	newGame.show()
}
