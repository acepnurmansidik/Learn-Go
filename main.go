package main

import (
	"fmt"
)

type Customer struct{
	Name string
	Address string
	Age int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is",customer.Name)	
}

func (name Customer)sayHuuu(){
	fmt.Println("Huu from",name.Name)
}

func main(){
	// TODO: STRUCT
	
	// kumpulan dari beberapa data
	var person Customer
	person.Name = "acep"
	person.Address = "Garut"
	person.Age = 23

	fmt.Println(person)

	// struct literal
	people := Customer{
		Name: "nurman",
		Address: "garut",
		Age: 23,
	}

	fmt.Println(people)

	// struct method
	person.sayHello("nurman")
	person.sayHuuu()

}

