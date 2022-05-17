package main

import "fmt"

func main(){
	// TODO: Tipe data map

	person := map[string]string{
		"name": "acep",
		"address": "garut",
		"age": "23",
	}

	fmt.Println(person)
	// add map
	person["role"] = "Fulsstack Developer"
	fmt.Println(person)
	// modify map
	person["age"] = "20"
	fmt.Println(person)
	// delete map
	delete(person, "age")

}