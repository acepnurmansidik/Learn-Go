package main

import (
	"fmt"
)

func random() interface{} {
	return "acep"
}

func main() {
	// TODO: TYPE ASSORTIONS

	var result interface{} = random()
	// rubah yang tadinya interface kosong menjadi string
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// type mengambil tipe dari interface kosong
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown type")
	}
}
