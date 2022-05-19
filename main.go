package main

import "fmt"

func main(){
	// TODO: FUNCTION AS VALUE

	result := getGoodBye
	fmt.Println(result("acep"))
}

func getGoodBye(name string)string{
	return "Good bye "+name 
}
