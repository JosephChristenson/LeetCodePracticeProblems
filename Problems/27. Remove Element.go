package main

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.

// import (
// 	"fmt"
// )

// func main() {
// 	nums := []int{3, 2, 2, 3}
// 	val := 3
// 	fmt.Println(removeElement(nums, val))
// 	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	val = 2
// 	fmt.Println(removeElement(nums, val))
// }

// func removeElement(nums []int, val int) int {
// 	index := len(nums) - 1

// 	for index >= 0 {
// 		if nums[index] == val {
// 			nums = remove(nums, index)
// 		}
// 		index--
// 	}

// 	return len(nums)
// }

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }
