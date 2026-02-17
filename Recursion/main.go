package main

import "fmt"

func naturalNumberSum(n int) int {
	if n == 1 {
		return 1
	}
	return n + naturalNumberSum(n-1)
}
func main() {

	fmt.Println("Natural Number sum of 6 is", naturalNumberSum(6))
}
