package main

import (
	"fmt"
)

func main(){
	// TODO: BREAK & CONTINUE

	// break
	for i := 0; i < 10; i++{
		if i == 5{
			break
		}
		fmt.Println("Perulangan ke",i)
	}

	// continue
	fmt.Println("")
	fmt.Println("CONTINUE #########################")
	for i := 0; i < 10; i++{
		if i % 2 == 0{
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}