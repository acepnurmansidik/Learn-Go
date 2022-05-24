package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool{
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}


func main() {
	// TODO: PACKAGE REFLECT | Validation Library
	// structTag bisa digunakan untuk validasi

	sample := Sample{"Acep"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(isValid(sample))
	sample.Name = ""
	fmt.Println(isValid(sample))
}
