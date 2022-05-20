package main

import (
	"fmt"
)

func Ups(value int) interface{} {
	if value == 1 {
		return 1
	} else if value == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// TODO: INTERFACE

	// interface biasa digunakan untuk function2 yang general

	// interface kosong
	var data interface{} = Ups(5)
	fmt.Println(data)
}
