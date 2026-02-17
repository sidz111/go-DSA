package main

import (
	"fmt"
)

func main() {
	str := "Hello"

	//length of String
	fmt.Println("len is", len(str))

	//Reversed String
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
}
