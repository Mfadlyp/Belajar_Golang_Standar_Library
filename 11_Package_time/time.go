package main

import (
	"fmt"
	"time"
)

func main() {
	// memnagil waktu saat ini
	now := time.Now()
	fmt.Println("Time : ", now.Local())

	// untuk membuat waktu
	createTimeUTC := time.Date(2023, time.December, 23, 23, 0, 0, 0, time.UTC)
	fmt.Println("Membuat date", createTimeUTC.Local())

	// parse tanggal dari string
	// parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	// fmt.Println(parse)

	// parse tanggal yang lain
	formatter := "2006-01-02 15:04:05"

	value := "2023-12-23 10:10:10"

	parse, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(parse)
	}

	// cara membuat duration
	duration_1 := time.Second * 100
	duration_2 := time.Millisecond * 10
	duration_3 := duration_1 - duration_2

	fmt.Println(duration_3)
}
