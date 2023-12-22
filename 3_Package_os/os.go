package main

import (
	"fmt"
	"os"
)

func main() {
	// args untuk mengambil data argumen yg ketika aplikasi dijalankan
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// contoh lain

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

}
