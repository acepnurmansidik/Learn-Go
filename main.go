package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	// TODO: PACKAGE REFLECT

	sample := Sample{"Acep"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
}
