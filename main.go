package main

import "fmt"

func main(){
	// TODO: VARIABEL
	var name string

	name = "Acep"
	fmt.Println(name)
	name = "Acep Nurman Sidik"
	fmt.Println(name)

	// =============
	var age uint8 = 24
	fmt.Println(age)

	// deklasrasi variabel tanpa var (hanya diawal)
	address:="Luzern"
	fmt.Println(address)
	address = "Florida"
	fmt.Println(address)

	// Deklarasi multiple variable
	var (
		firstname="acep nurman "
		lastname="sidik"
	)

	fmt.Print(firstname)
	fmt.Println(lastname)

}