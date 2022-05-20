package main

import (
	"fmt"
)

type Person struct{
	Name string
}

type Animal struct{
	Name string
}

type HasName interface{
	GetName() string
}

func sayHello(name HasName){
	fmt.Println("Hello",name.GetName())
}

func (person Person) GetName()string{
	return person.Name
}

func (animal Animal) GetName() string{
	return animal.Name
}


func main(){
	// TODO: INTERFACE

	// interface biasa digunakan untuk function2 yang general
	person := Person{Name: "acep"}
	sayHello(person)
	
	cat := Animal{Name: "Kucing"}
	sayHello(cat)
}

