package main

import (
	"fmt"
)

func main(){
	// TODO: DEFER, PANIC & RECOVER

	// defer
	// akan menjalankan function ketika sudah selesai programnya meskipun ada error
	logging := func ()  {
		fmt.Println("Selesai memangil function")
	}

	runApplication := func(value int){
		defer logging()
		fmt.Println("Run Application")
		result := 10 /value
		fmt.Println(result)
	}

	fmt.Println("## D E F E R ##################")
	runApplication(10)
	
	// Panic
	endApp := func ()  {
		fmt.Println("Aplikais selesai")
	}

	rundApp := func (error bool)  {
		defer endApp()
		if error {
			panic("APLIKAIS ERROR!!!")
		}
		fmt.Println("Aplikasi berjalan")
	}
	fmt.Println("")
	fmt.Println("## P A N I C ##################")
	rundApp(false)

	fmt.Println("")
	fmt.Println("## R E C O V E R ##################")
	// recover akan menangkap response dari panik
	// recover hanya disimpan di defer
	// recover akan terus berjalan meskipun ada error
	akhirApp := func ()  {
		message := recover()
		if message != nil{
			fmt.Println("Error denga message: ",message)
		}
		fmt.Println("Aplikasi selesai")
	}
	App := func (error bool)  {
		defer akhirApp()
		if error {
			panic("APLIKASI ERROR!")
		}
	}

	App(true)
}

