package main

import (
	"fmt"
	"time"
)


func main() {
	// TODO: PACKAGE TIME

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())

	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse(time.RFC3339, "2022-05-24")
	fmt.Println(parse)
}
