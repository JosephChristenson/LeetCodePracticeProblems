package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := float64(0)
	for i := 0; i < k; i++ {
		max = max + float64(nums[i])
	}
	max = max / float64(k)
	previous := max
	for index := 1; index+k <= len(nums); index++ {
		previous = ((previous * float64(k)) - float64(nums[index-1]) + float64(nums[index+k-1])) / float64(k)
		if previous > max {
			max = previous
		}
	}
	return max
}
