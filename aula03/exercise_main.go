package main

type Key interface {
	Press() interface{}
}

type Keyboard struct {
	LastKeyPressed interface{}
	Keys           []Key
}

func main() {
	// Exercício: Crie um Keyboard com pelo menos duas key (de tipos distintos: string, byte, int, ...).
	// - Faça uma função para pressionar todas as Keys do Keyboard e depois imprima o valor da última Key pressionada.
}
