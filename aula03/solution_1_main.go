package main

import "fmt"

type Key interface {
	Press() interface{}
}

type Keyboard struct {
	LastKeyPressed interface{}
	Keys           []Key
}

func (keyboard *Keyboard) PressAllKeys() {
	for _, key := range keyboard.Keys {
		keyboard.LastKeyPressed = key.Press()
	}
}

type KeyHelloWorld struct {
	key string
}

func NewKeyHelloWorld() KeyHelloWorld {
	return KeyHelloWorld{"Hello World"}
}

func (khw KeyHelloWorld) Press() interface{} {
	return khw.key
}

type Key100 struct{}

func (k Key100) Press() interface{} {
	return 100
}

func main() {
	// keyboard := Keyboard{
	// 	Keys: []Key{
	// 		NewKeyHelloWorld(),
	// 		Key100{},
	// 	},
	// }

	keyboard := Keyboard{}
	keyboard.Keys = append(keyboard.Keys, Key100{}, NewKeyHelloWorld())

	keyboard.PressAllKeys()
	fmt.Println(keyboard.LastKeyPressed, keyboard.Keys)

	switch v := keyboard.LastKeyPressed.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("default", v)
	}
}
