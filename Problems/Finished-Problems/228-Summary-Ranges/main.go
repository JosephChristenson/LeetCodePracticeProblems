package main

import (
	"fmt"
	"strconv"
)

// You are given a sorted unique integer array nums.

// A range [a,b] is the set of all integers from a to b (inclusive).

// Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

// Each range [a,b] in the list should be output as:

// "a->b" if a != b
// "a" if a == b

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))

}

func summaryRanges(nums []int) []string {

	resultNum := []int{}
	resultString := []string{}

	for _, val := range nums {
		if len(resultNum) == 0 {
			resultNum = append(resultNum, val)
		} else if len(resultNum) == 1 {
			if val == resultNum[0]+1 {
				resultNum = append(resultNum, val)
			} else {
				resultString = append(resultString, strconv.Itoa(resultNum[0]))
				resultNum = []int{val}
			}
		} else if resultNum[1] == val-1 {
			resultNum[1] = val
		} else {
			temp := strconv.Itoa(resultNum[0]) + "->" + strconv.Itoa(resultNum[1])
			resultString = append(resultString, temp)
			resultNum = []int{val}
		}
	}

	if len(resultNum) == 1 {
		resultString = append(resultString, strconv.Itoa(resultNum[0]))
	} else if len(resultNum) == 2 {
		temp := strconv.Itoa(resultNum[0]) + "->" + strconv.Itoa(resultNum[1])
		resultString = append(resultString, temp)
	}
	return resultString
}
