package main

import "fmt"

func main(){
	// TODO: PNEGKONDISIAN /PERCABANGAN

	name := "acep"

	switch name{
	case "Nurman":
		fmt.Println("Hello nurmnan")
	case "Sidik":
		fmt.Println("Hello sidik")
	default:
		fmt.Println("kenalan yu")
	}
	// short statement
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Too long")
	case false:
		fmt.Println("Too short")
	}

	// swith with expression
	switch {
	case len(name) > 10:
		fmt.Println("nama terlalu panjang")
	case len(name) > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama terlalu pendek")
	}


}