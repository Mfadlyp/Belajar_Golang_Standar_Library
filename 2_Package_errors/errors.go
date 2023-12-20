package main

import (
	"errors"
	"fmt"
)

var (
	ValidationErr = errors.New("Validation Error")
	NotFoundError = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationErr
	}

	if id == "fadly" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("")

	if err != nil {
		// menggunakan erros.Is() untuk cek jenis error
		if errors.Is(err, ValidationErr) {
			fmt.Println("ValidationErr")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("NotFoundError")
		} else {
			fmt.Println("Unkown Error")
		}
	}
}
