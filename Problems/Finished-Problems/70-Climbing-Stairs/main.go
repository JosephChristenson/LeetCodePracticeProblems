package main

import (
	"fmt"
)

// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func main() {
	n := 2
	fmt.Println(climbStairs(n))
	n = 3
	fmt.Println(climbStairs(n))
	n = 4
	fmt.Println(climbStairs(n))
	n = 5
	fmt.Println(climbStairs(n))
	n = 6
	fmt.Println(climbStairs(n))
	n = 1
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {
	store1, store2 := 1, 0

	for n > 1 {
		temp := store1
		store1 = store1 + store2
		store2 = temp
		n--
	}
	return store1 + store2
}

// func climbStairs(n int) int {
// 	return possibleSteps(0, n)
// }

// I like this solution but it is too slow

// func possibleSteps(n int, goal int) int {
// 	if n+1 == goal {
// 		return 1
// 	}
// 	if n+2 == goal {
// 		return 2 // could be the last step is n + 2 or n + 1 + 1
// 	}
// 	return possibleSteps(n+1, goal) + possibleSteps(n+2, goal)
// }
