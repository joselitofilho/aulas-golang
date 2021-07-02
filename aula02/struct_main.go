package main

import "fmt"

type User struct {
	name   *string
	age    int
	values map[string]string
	list   []string
}

func main() {
	// user1 := User{
	// 	name: "Joselito",
	// 	age:  33,
	// }

	user1 := &User{}
	name := "Joselito"
	user1.name = &name
	*user1.name = "Joselito0000"
	user1.age = 33

	fmt.Println("   name: ", name)
	fmt.Println("address: ", user1.name)
	fmt.Println("content: ", *user1.name)
	fmt.Println("   user: ", user1)
}
