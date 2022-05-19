package main

import "fmt"

func main(){
	// TODO: RECURSIVE FUNCTION

	// menggunakan for loops
	blocklist := func (value int)int  {
		result := 1
		for i := value; i > 0; i--{
			result *= i
		}
		return result
	}

	fmt.Println("without recursive")
	fmt.Println(blocklist(5))

	// Recursive
	fmt.Println("recursive")
	fmt.Println(factorialRecursive(5))
}

// jika terlalu aonajng bisa menggunakan alias
type Blocklist func(string) bool

func registerUser(name string, blocklist Blocklist){
	if blocklist(name) {
		fmt.Println("Yo're blocked",name)
	}else {
		fmt.Println("Welcome",name)
	}
}

	func factorialRecursive(value int)int  {
		if(value == 1){
			return 1
		}else {
			return value * factorialRecursive(value-1)
		}
		
	}