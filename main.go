package main

import "fmt"

func main(){
	// TODO: ARAAY

	var names [3]string

	names[0] = "Acep"
	names[1] = "Nurman"
	names[2] = "Sidik"

	fmt.Println(names[0])

	// deklarasi array langsung di value
	var values = [3]int{
		89,
		87,
		25,
	}
	fmt.Println(values)
	fmt.Println(values[2])

	// function array
	fmt.Println(len(values))
}