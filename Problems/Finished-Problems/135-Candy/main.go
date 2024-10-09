package main

import "fmt"

// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

// You are giving candies to these children subjected to the following requirements:

// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println(candy(ratings))
	ratings = []int{1, 2, 2}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {

	results := make([]int, len(ratings))

	for index, _ := range ratings {
		if index == 0 {
			results[index] = 1
		} else {
			if ratings[index] > ratings[index-1] {
				results[index] = results[(index-1)] + 1
			} else {
				results[index] = 1
			}
		}
	}
	index := len(ratings) - 1
	total := 0

	for index >= 0 {
		temp := 0
		if index == len(ratings)-1 {
			temp = 1
		} else {
			if ratings[index] > ratings[index+1] {
				temp = results[(index+1)] + 1
			} else {
				temp = 1
			}
		}

		if results[index] < temp {
			results[index] = temp
		}
		total = total + results[index]
		index--
	}
	return total
}
