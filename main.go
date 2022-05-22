package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: PACKAGE OS

	// Args
	result := os.Args
	fmt.Println(result)
	
	// hostname, mengambil nama host system operasi yangi digunakan
	hostname, err := os.Hostname()
	if err == nil{
		fmt.Println(hostname)
	}else {
		fmt.Println("Error",err.Error())
	}
}
