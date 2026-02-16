package main

import "fmt"

func main() {

	arr := []int{7, 1, 5, 3, 6, 4}

	maxProfit := 0
	minPrice := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < minPrice {
			minPrice = arr[i]
		}
		profit := arr[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	fmt.Println("Max Profit is:", maxProfit)
}
