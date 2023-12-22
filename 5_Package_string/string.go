package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains untuk cek string yg dicari
	fmt.Println(strings.Contains("Fadly Pangestu", "Fadly"))

	// split memotong string berdasarkan separator
	fmt.Println(strings.Split("M Fadly Pangestu", ""))

	// Tolower
	fmt.Println(strings.ToLower("FADLY PANGESTU"))

	// ToUpper
	fmt.Println(strings.ToUpper("fadly pangestu"))

	// Trim memotong cutset diawalan dan akhir
	fmt.Println(strings.Trim("	fadly pangestu	", ""))

	// ReplaceAll
	fmt.Println(strings.ReplaceAll("fadly fadly", "fadly", "buday"))
}
