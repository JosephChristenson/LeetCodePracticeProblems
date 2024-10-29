package main

import "fmt"

// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 && n == 1 {
			return true
		} else {
			return false
		}
	}

	index := 0
	for n > 0 && index < len(flowerbed) {
		if index > 0 {
			if index+1 == len(flowerbed) {
				if flowerbed[index-1] == 0 && flowerbed[index] == 0 {
					flowerbed[index] = 1
					n--
				}
			} else {
				if flowerbed[index-1] == 0 && flowerbed[index] == 0 && flowerbed[index+1] == 0 {
					flowerbed[index] = 1
					n--
				}
			}
		} else {
			if flowerbed[1] == 0 && flowerbed[0] == 0 {
				flowerbed[0] = 1
				n--
			}
		}
		index++
	}
	if n > 0 {
		return false
	} else {
		return true
	}
}
