package main

import "fmt"

type Reader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write(data []byte) (int, error)
}

func main() {
	fmt.Println("hello")
}
