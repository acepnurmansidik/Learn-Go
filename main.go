package main

import "fmt"

func main(){
	// TODO: OPERASI BOOLEAN
	nilaiAkhir := 90
	absensi := 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

}