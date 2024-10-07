package main

// import (
// 	"fmt"
// )

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	nums2 := []int{2, 5, 6}
// 	m := 3
// 	n := 3
// 	merge(nums1, m, nums2, n)
// 	fmt.Println(nums1)

// 	nums1 = []int{1}
// 	nums2 = []int{}
// 	m = 1
// 	n = 0
// 	merge(nums1, m, nums2, n)
// 	fmt.Println(nums1)

// 	nums1 = []int{0}
// 	nums2 = []int{1}
// 	m = 0
// 	n = 1
// 	merge(nums1, m, nums2, n)
// 	fmt.Println(nums1)
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {

// 	for n+m > 0 {
// 		if n == 0 {
// 			nums1[n+m-1] = nums1[m-1]
// 			m--
// 		} else if m == 0 {
// 			nums1[n+m-1] = nums2[n-1]
// 			n--
// 		} else {
// 			if nums1[m-1] >= nums2[n-1] {
// 				nums1[n+m-1] = nums1[m-1]
// 				m--
// 			} else {
// 				nums1[n+m-1] = nums2[n-1]
// 				n--
// 			}
// 		}
// 	}

// }
