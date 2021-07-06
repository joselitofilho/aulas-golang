package main

import "fmt"

type reader interface {
	read() (int, error)
}

type writer interface {
	write() (int, error)
}

type readWriter interface {
	reader
	writer
}

type printer struct {
	name string
}

func (p printer) read() (int, error) {
	fmt.Println("--------------------------------------")
	fmt.Println("Read()", p.name)
	fmt.Println("--------------------------------------")
	return 0, nil
}

func (p printer) write() (int, error) {
	fmt.Println("--------------------------------------")
	fmt.Println("Write()", p.name)
	fmt.Println("--------------------------------------")
	return 0, nil
}

func main() {
	p := printer{"Printer"}
	var rw readWriter = p
	// rw.read()
	// rw.write()

	if w, ok := rw.(writer); ok {
		w.write()
	} else {
		panic("w não é writer")
	}
}
