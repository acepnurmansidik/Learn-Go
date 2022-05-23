package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO: PACKAGE STRING

	fmt.Println(strings.Contains("Acep Nurman Sidik", "Nurman"))
	fmt.Println(strings.Split("Acep Nurman Sidik", " ")[0])
	fmt.Println(strings.ToLower("Acep Nurman Sidik"))
	fmt.Println(strings.ToUpper("Acep Nurman Sidik"))
	fmt.Println(strings.Trim("    Acep Nurman Sidik      ", " "))
	fmt.Println(strings.ReplaceAll("Acep Nurman Sidik", "Nurman","Budi"))

}
