package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writter := csv.NewWriter(os.Stdout)

	_ = writter.Write([]string{"eko", "buday", "wahyu"})
	_ = writter.Write([]string{"nana", "nono", "ajax"})
	_ = writter.Write([]string{"koci", "botak", "jamet"})

	writter.Flush()
}
