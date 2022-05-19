package main

import "fmt"

func main(){
	// TODO: FUNCTION AS PARAMETER

	sayHelloWithFilter("Anjing",spamFilter)

}

// jika terlalu aonajng bisa menggunakan alias
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	nameFilter := filter(name)
	fmt.Println("Hello",nameFilter)
}

func spamFilter(name string)string{
	if (name == "Anjing"){
		return "..."
	}else{
		return name
	}
}
