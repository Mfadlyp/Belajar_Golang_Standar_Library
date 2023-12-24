package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	//			struct tag
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name   string `required:"true" max:"10"`
	Addres string `required:"true" max:"10"`
	Age    string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Nilai value", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)

		fmt.Println(structField.Name, "with tpye", structField.Type)

		// stelah akses field, bisa akses tag nya
		fmt.Println("Hasil struct field tag required : ", structField.Tag.Get("required"))
		fmt.Println("Hasil struct field tag max : ", structField.Tag.Get("max"))
	}

}

func isValid(value any) (result bool) {
	result = true
	datavalid := reflect.TypeOf(value)
	for i := 0; i < datavalid.NumField(); i++ {
		field := datavalid.Field(i)

		// lalu cek
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	// sample := Sample{"Fadly"}
	// sampleType := reflect.TypeOf(sample)
	// structField := sampleType.Field(0)
	// fmt.Println(structField.Name)

	readField(Sample{"fadly"})

	// akses isValid
	person := Person{
		Name:   "Fadly",
		Addres: "sehat",
		Age:    "20",
	}
	fmt.Println(isValid(person))
}
