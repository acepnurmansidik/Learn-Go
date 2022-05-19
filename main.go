package main

import "fmt"

func main(){
	// TODO: CLOSURE
	name := "acep"
	counter := 0

	increment := func ()  {
		name := "nurman"
		fmt.Println("increment")
		fmt.Println(name)
		counter++
	}

	increment()
	fmt.Println(name)
	fmt.Println(counter)
}

