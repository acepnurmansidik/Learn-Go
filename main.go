package main

import "fmt"

func main(){
	// TODO: VARIADIC FUNCTION
	result := sumAll(10,10)
	fmt.Println(result)

	// slice parameter
	slice := []int{10,20}
	total := sumAll(slice...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int{
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
