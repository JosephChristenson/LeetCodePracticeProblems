package main

import (
	"fmt"
)

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

// Find and return the maximum profit you can achieve.

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	hold := -1
	total := 0

	for index := 0; index+1 < len(prices); index++ {
		if hold == -1 {
			if prices[index] < prices[index+1] {
				hold = prices[index]
			}
		} else {
			if prices[index] > prices[index+1] {
				total = total + prices[index] - hold
				hold = -1
			}
		}
	}
	if hold != -1 {
		total += prices[len(prices)-1] - hold
	}
	return total
}
