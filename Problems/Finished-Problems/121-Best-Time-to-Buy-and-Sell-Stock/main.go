package main

import "fmt"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	highestProfit, lowestPrice := 0, prices[0]
	for _, val := range prices {
		if val-lowestPrice > highestProfit {
			highestProfit = val - lowestPrice
		}
		if val < lowestPrice {
			lowestPrice = val
		}
	}
	return highestProfit
}
