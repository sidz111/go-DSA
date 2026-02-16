package main

import "fmt"

func main() {
	m := make(map[int]int)

	nums := []int{2, 11, 7, 15}
	target := 9
	// Step 2: Loop through array
	for i, num := range nums {

		// Step 3: Find the remaining number
		remain := target - num

		// Step 4: Check if remain already in map
		if index, found := m[remain]; found {
			fmt.Println("Index", index, "and", i)
		}

		// Step 5: Store current number and index in map
		m[num] = i
	}

}
