package main

import "fmt"

func main(){
	// TODO: SLICE

	days := []string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"minggu",
	}

	var slice1 = days[5:7]

	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// slice1[0] = "sabtu lagi"
	// fmt.Println(slice1)

	// days[5] = "sabtu lagi"
	// fmt.Println(days)

	// Array vs Slice

	iniArray := [3]int{1,2,3}
	iniSlice := []int{1,2,3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}