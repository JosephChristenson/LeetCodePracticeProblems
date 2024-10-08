package main

import (
	"fmt"
)

func main() {
	nums1, nums2 := []int{1, 3}, []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums1, nums2 = []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums1, nums2 = []int{}, []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	totalnum1, totalnum2 := len(nums1), len(nums2)
	total := totalnum1 + totalnum2

	IsOddTotal := total % 2
	MidPoint := total / 2
	index1, index2 := 0, 0

	for index1+index2 <= MidPoint-2+IsOddTotal {
		if index1 == totalnum1 {
			index2++
		} else if index2 == totalnum2 {
			index1++
		} else if nums1[index1] <= nums2[index2] {
			index1++
		} else {
			index2++
		}
	}

	if IsOddTotal == 1 {
		if index1 == totalnum1 {
			return float64(nums2[index2])
		}
		if index2 == totalnum2 {
			return float64(nums1[index1])
		}
		if nums1[index1] < nums2[index2] {
			return float64(nums1[index1])
		} else {
			return float64(nums2[index2])
		}
	} else {
		first, second := 0, 0

		if index1 == totalnum1 {
			first = nums2[index2]
			second = nums2[index2+1]
		} else if index2 == totalnum2 {
			first = nums1[index1]
			second = nums1[index1+1]
		} else {
			if nums1[index1] < nums2[index2] {
				first = nums1[index1]
				index1++
			} else {
				first = nums2[index2]
				index2++
			}

			if index1 == totalnum1 {
				second = nums2[index2]
			} else if index2 == totalnum2 {
				second = nums1[index1]
			} else if nums1[index1] < nums2[index2] {
				second = nums1[index1]
				index1++
			} else {
				second = nums2[index2]
				index2++
			}
		}
		return float64(first+second) / 2

	}

}
