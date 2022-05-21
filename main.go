package main

import (
	"fmt"
)

type Address struct{
	City, Province, Country string
}

func random() interface{} {
	return "acep"
}

func main() {
	// TODO: POINTER

	// pass by value, jika di asignment ke variabel lain valuenya kan di copy
	address1 := Address{"Subang", "West Java", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Garut"

	// jka ditambahkan pointer seolah akan membuat memory baru
	// address2 = &Address{"Subang", "West Java", "Indonesia"}

	// jika ditambahkan bintang di depannya maka akan memaksa semua variabel ke address 2
	*address2 = Address{"Subang", "West Java", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// pointer new, akan mengembalikan data kosong/tidak ada data awal
	address4 := new(Address)
	address4.City = "Surabaya"
	fmt.Println(address4)


}
