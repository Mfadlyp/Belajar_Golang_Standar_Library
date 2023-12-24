package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("fadly"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("BuDi"))
}
