package main

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

//Brute Force answer
// func twoSums(nums []int, target int) []int {

// 	for firstindex, element := range nums {
// 		for secondindex, secondelement := range nums {
// 			if firstindex != secondindex {
// 				if (element + secondelement) == target {
// 					result := []int{firstindex, secondindex}
// 					return result
// 				}
// 			}
// 		}
// 	}

// 	result := []int{0}
// 	return result
// }
