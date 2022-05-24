package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {
	// TODO: PACKAGE REFLECT | StructTag

	sample := Sample{"Acep"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
}
