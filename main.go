package main

import (
	"fmt"
)

type Man struct{
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main() {
	// TODO: POINTER DI METHOD

	// jika buat struct function/methode disarankan menggunakan pointer
	acep := Man{Name: "acep"}
	acep.Married()
	fmt.Println(acep.Name)
}
