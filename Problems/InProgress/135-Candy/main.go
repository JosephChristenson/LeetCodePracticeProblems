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

type child struct {
	rating int
	candy  int
}

func candy(ratings []int) int {
	var m map[int]child
	m = make(map[int]child)
	amount, lowestamount := 0, 1

	for index, value := range ratings {

		if value < m[index-1].rating { // if the previous child is higher rank go down 1 candy
			amount = m[index-1].candy - 1
		} else { // current child is better than previous child 1 more candy
			amount = m[index-1].candy + 1
		}
		if amount < lowestamount {
			lowestamount = amount
		}
		m[index] = child{
			rating: value,
			candy:  amount,
		}
	}
	for index, value := range m{
		if m[index].rating < m[index - 1]

	}
}
