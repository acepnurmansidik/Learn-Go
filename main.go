package main

import "fmt"

func main(){
	// TODO: KONVERSI TIPE DATA

	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "acep"
	a := string(name[0])

	fmt.Println(name)
	fmt.Println(a)


}