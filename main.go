package main

import (
	"fmt"
	"strconv"
)

func main() {
	// TODO: PACKAGE STRICONV

	// parse to bool
	result, err := strconv.ParseBool("true")
	if err == nil{
		fmt.Println(result)
	}else {
		fmt.Println("Error ", err.Error())
	}

	// parse to int
	number, err := strconv.ParseInt("1024", 10,32)
	if err == nil{
		fmt.Println(number)
	}else {
		fmt.Println("Error ", err.Error())
	}

	value := strconv.FormatInt(10000, 16)
	fmt.Println(value)
}
