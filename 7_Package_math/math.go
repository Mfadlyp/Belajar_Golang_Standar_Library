package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Round(float64) membulatkan float 64 keatas atau kebawah sesuai paling dekat
	fmt.Println(math.Round(1.40))

	// math.Floor(float64) membulatkan float64 kebawah
	fmt.Println(math.Floor(-2.35)) // -3

	// math.Ceil(float64) membulatkan float64 keatas
	fmt.Println(math.Ceil(2.78)) // 3

	// math.Max(float64, float64) mengembalikan nilai float64 keatas
	fmt.Println(math.Max(9, 3)) // 9

	// math.Min(float64, float64) mengembalikan nilai float64 kebawah
	fmt.Println(math.Min(9, 3)) // 3
}
