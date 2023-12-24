package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// encoding sebuah data ke string
	encodeData := base64.StdEncoding.EncodeToString([]byte("M Fadly Pangestu"))
	fmt.Println(encodeData)

	// decoding data
	decodeData, err := base64.StdEncoding.DecodeString(encodeData)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decodeData))
	}

	
}
