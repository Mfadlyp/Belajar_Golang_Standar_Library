package main

import (
	"container/ring"
	"strconv"
)

func main() {
	// membuat ring dengan panjang 5
	r := ring.New(5)

	// membuat value yg akan diisi di ring
	for i := 0; i < r.Len(); i++ {
		r.Value = "Value" + strconv.Itoa(i)
		r = r.Next()
	}

}
