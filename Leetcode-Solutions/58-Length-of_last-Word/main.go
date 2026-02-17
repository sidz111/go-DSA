package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Siddhu qwe    "

	result := strings.TrimRight(s, " ")

	// fmt.Println(result)
	// for i := len(result) - 1; i >= 0; i-- {
	// 	fmt.Print(string(result[i]))
	// }

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(string(result[i]))
		if string(result[i]) == " " {
			// fmt.Println(i - 1)
			fmt.Println(len(result) - i - 1)
			break
		}
	}
}
