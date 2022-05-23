package main

import (
	"container/list"
	"fmt"
)

func main() {
	// TODO: PACKAGE CONTAINER/LIST

	data := list.New()

	data.PushBack("acep")
	data.PushBack("nurman")
	data.PushBack("sidik")
	data.PushFront("budi")

	// depan ke belakang
	for element := data.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)
	}
	fmt.Println("===========================================")

	// belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev(){
		fmt.Println(element.Value)
	}
}
