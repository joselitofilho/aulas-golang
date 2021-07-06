package main

import "fmt"

type Key interface {
	Press() interface{}
}

type Keyboard struct {
	LastKeyPressed interface{}
	Keys           []Key
}

type KeyA struct{}

func (k KeyA) Press() interface{} {
	return "A"
}

type KeyB struct{}

func (k KeyB) Press() interface{} {
	return "B"
}

type Key1 struct{}

func (k *Key1) Press() interface{} {
	return 1
}

func (keyboard *Keyboard) PressAllKeys() {
	for _, k := range keyboard.Keys {
		keyboard.LastKeyPressed = k.Press()
		fmt.Print(keyboard.LastKeyPressed)
	}
}

func main() {
	keyboard := &Keyboard{Keys: []Key{&KeyA{}, &KeyB{}, &Key1{}}}
	keyboard.PressAllKeys()
	fmt.Println("\n\nLastKeyPressed =", keyboard.LastKeyPressed)

	switch value := keyboard.LastKeyPressed.(type) {
	case int:
		fmt.Println("LastKeyPressed is an integer = ", value)
	case string:
		fmt.Println("LastKeyPressed is a string = ", value)
	}
}
