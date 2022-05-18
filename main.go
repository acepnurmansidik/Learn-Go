package main

import "fmt"

func main(){
	// TODO: FUNCTION

	// basic function
	sayHello()

	// function with paramaeter
	sayHello2("acep", "nurman")

	// function returning value
	fmt.Println(sayHelloReturn("acep"))

	// returning multiple value
	firstName, _, lastName := getFullName()
	fmt.Println(firstName,lastName)
}

func sayHello(){
	fmt.Println("Hello World")
}

func sayHello2(firstName string, lastName string){
	fmt.Println("Hello", firstName,lastName)
}

func sayHelloReturn(name string)string{
	return "Hello " + name
}

func getFullName()(string,string,string){
	return "acep","nurman","sidik"
}