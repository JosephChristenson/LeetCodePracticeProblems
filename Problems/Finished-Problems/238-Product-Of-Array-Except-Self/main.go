package main

import "fmt"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
func productExceptSelf(nums []int) []int {
	total := 1
	numOfZeros := 0
	for _, val := range nums {
		if val != 0 {
			total = total * val
		} else {
			numOfZeros++
		}
	}
	result := make([]int, 0, len(nums))
	if numOfZeros <= 1 {
		for _, val := range nums {
			if numOfZeros == 1 {
				if val != 0 {
					result = append(result, 0)
				} else {
					result = append(result, total)
				}
			} else {
				result = append(result, total/val)
			}
		}
	} else {
		result = make([]int, len(nums))
	}
	return result
}
