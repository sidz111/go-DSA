package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	pro := 1
	result := make([]int, len(nums))
	for i := range nums {
		pro *= nums[i]
	}
	for i := range nums {
		result[i] = pro / nums[i]
	}
	fmt.Println("answer is", result)

}
