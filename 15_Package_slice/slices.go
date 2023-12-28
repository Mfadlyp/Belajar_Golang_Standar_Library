package main

import (
	"fmt"
	"slices"
)

func main() {
	nameString := []string{"john", "paul", "udin", "samsul"}
	ageInt := []int{23, 34, 45, 32}

	fmt.Println(slices.Max(ageInt))
	fmt.Println(slices.Min(ageInt))

	// mencari nama sesuai dengan kehendak
	fmt.Println(slices.Contains(nameString, "paul"))

	// mencari index
	fmt.Println(slices.Index(nameString, "samsul"))
}
