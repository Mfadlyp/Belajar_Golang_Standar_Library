package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ParseBool mengubah string ke bool
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", result)
	}

	// Fungsi strconv.Itoa digunakan untuk mengkonversi nilai integer ke dalam bentuk string.
	resultInt := strconv.Itoa(999)
	fmt.Println("Integer to String: ", resultInt)

	// Fungsi strconv.Atoi digunakan untuk mengkonversi nilai string ke dalam bentuk integer.
	//Perhatikan bahwa Atoi mengembalikan dua nilai, yaitu nilai integer dan nilai error.
	//Pastikan untuk memeriksa kesalahan sebelum menggunakan nilai integer.

	resultStr, err := strconv.Atoi("999")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Printf("String to Integer: %d\n", resultStr)
	}
}
