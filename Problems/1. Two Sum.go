package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	fmt.Print(twoSums(nums, target))
// 	nums = []int{3, 2, 4}
// 	target = 6
// 	fmt.Print(twoSums(nums, target))
// 	nums = []int{3, 3}
// 	target = 6
// 	fmt.Print(twoSums(nums, target))
// }

// func twoSums(nums []int, target int) []int {

// 	checkedvalues := []int{}

// 	for index, value := range nums {

// 		result := contains(checkedvalues, target-value)
// 		if result != -1 {
// 			return []int{result, index}
// 		}
// 		checkedvalues = append(checkedvalues, value)
// 	}
// 	return nil
// }

// func contains(s []int, e int) int {
// 	for result, a := range s {
// 		if a == e {
// 			return result
// 		}
// 	}
// 	return -1
// }
