package main

func main() {
	nums := []int{2, 0, 1, 1, 0}
	sortColors(nums)
	nums = []int{2, 0, 1}
	sortColors(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortColors(nums []int) {
	index := 0
	end := len(nums) - 1
	if len(nums) >= 2 {
		for index < end {
			if nums[index+1] < nums[index] {
				nums[index+1], nums[index] = nums[index], nums[index+1]
				if index != 0 {
					index--
				} else {
					index++
				}
			} else {
				index++
			}
		}
	}
}
