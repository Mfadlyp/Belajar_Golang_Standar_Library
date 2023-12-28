package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// bufio reader
	input := strings.NewReader("this line for string\n with number\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println("Cetak string", string(line))
	}

	// bufio writter
	write := bufio.NewWriter(os.Stdout)
	write.WriteString("Hello ")
	write.WriteString("Terima kasih")

	write.Flush()
}
