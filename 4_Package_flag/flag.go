package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host name")
	username := flag.String("username", "root", "Put your username")
	password := flag.String("password", "root", "Put your password")

	flag.Parse()
	fmt.Println(*host, *username, *password)
}
