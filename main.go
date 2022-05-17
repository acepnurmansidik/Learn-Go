package main

import (
	"fmt"
)

func main(){
	// TODO: PERULANGAN

	// for biasa
	fmt.Println("####### Perulangan For biasa ####")
	counter := 1
	for counter <= 3{
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for dengan statement
	fmt.Println("")
	fmt.Println("####### Perulangan for dengan statement ####")
	for counter1 := 1; counter1 <= 3; counter1++{
		fmt.Println("Perulangan ke", counter1)
	}

	// for range
	fmt.Println("")
	fmt.Println("####### Perulangan for dengan statement ####")
	// array / slice
	// tanda underscore berfungsi jika key nya tidak digunakan
	cars := []string{"Tesla", "Honda", "Toyoya"}
	for _, car := range cars {
		fmt.Println(car)
	}
	fmt.Println("")
	person := make(map[string]string)
	person["name"] = "acep"
	person["address"] = "garut"

	for key, value := range person{
		fmt.Println(key,"=", value)
	}
}