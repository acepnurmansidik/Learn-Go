package main

import "fmt"

func main(){
	// TODO: ANONYMOUS FUNCTION
	blocklist := func (name string)bool  {
		return name == "admin"
	}

	registerUser("admin", blocklist)
	registerUser("acep", blocklist)

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

