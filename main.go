package main

import "fmt"

func main(){
	// TODO: OERASI MATEMATIKA
	penjumlahan := 10 + 10
	pengurangan := 10 - 10
	perkalian := 10 * 10
	pembagian := 10 / 10
	modulus := 10 % 3

	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(pembagian)
	fmt.Println(modulus)
	
	fmt.Println("Augmanted assigments ##################################")

	// Augmanted assigments
	i := 10
	i += 10
	fmt.Println(i)
	i -= 8
	fmt.Println(i)
	i *= 10
	fmt.Println(i)
	i /= 10
	fmt.Println(i)
	i %= 10
	fmt.Println(i)

	fmt.Println("Unary operator ##################################")

	// Unary operator
	i++ // i = i + 1
	fmt.Println(i)
	i--
	fmt.Println(i)
	negative := -100
	fmt.Println(negative)
	positive := +100
	fmt.Println(positive)
	married := true
	fmt.Println(!married)

}