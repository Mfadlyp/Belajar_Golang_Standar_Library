package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("hello/fmt.go"))               // hello
	fmt.Println(path.Base("hello/fmt.go"))              // fmt.go
	fmt.Println(path.Ext("hello/fmt.go"))               // .go
	fmt.Println(path.Join("hello", "world", "main.go")) // hello/world/main.go

	// filepath
	fmt.Println(filepath.Dir("hello/fmt.go"))
	fmt.Println(filepath.Base("hello/fmt.go"))
	fmt.Println(filepath.Ext("hello/fmt.go"))
	fmt.Println(filepath.IsAbs("hello/fmt.go"))
	fmt.Println(filepath.IsLocal("hello/fmt.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))

}
