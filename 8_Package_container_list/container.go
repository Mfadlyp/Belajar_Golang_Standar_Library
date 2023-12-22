package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("M")
	data.PushBack("Fadly")
	data.PushBack("Pangestu")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
