package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	// csv reader
	// membauat string csv
	csvString := "m,fadly,pangestu\n" +
		"budi,senja,matahari\n" +
		"kolak,mak,enak"

	// membaca csv
	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
