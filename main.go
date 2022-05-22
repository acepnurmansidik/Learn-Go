package main

import (
	"fmt"
	"golang/database"
)

func main() {
	// TODO: PACKAGE INITIALIZATION
	// init berguna untuk mejalankan untuk pertama kali dieksekusi

	result := database.GetDatabase()
	fmt.Println(result)
}
