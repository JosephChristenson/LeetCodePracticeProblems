package main

import (
	"fmt"
	"slices"
)

// We are given an array asteroids of integers representing asteroids in a row.

// For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

// Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

func main() {
	asteroids := []int{5, 10, -5}
	fmt.Println(asteroidCollision(asteroids))
	asteroids = []int{-1, 5, -3, 10, -5, 3}
	fmt.Println(asteroidCollision(asteroids))
	asteroids = []int{8, -8}
	fmt.Println(asteroidCollision(asteroids))
	asteroids = []int{10, 2, -5}
	fmt.Println(asteroidCollision(asteroids))
}

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func asteroidCollision(asteroids []int) []int {
	results := make([]int, 0, len(asteroids))
	goingLeft := Stack{}
	goingRight := Stack{}

	for index := 0; index < len(asteroids); index++ {
		if asteroids[index] < 0 {
			if goingRight.IsEmpty() { // it won't collide so we can add it to results
				results = append(results, asteroids[index])
			} else {
				goingLeft.Push(asteroids[index])
				goingLeft, goingRight = collisions(goingLeft, goingRight)
			}
		} else if asteroids[index] > 0 {
			goingRight.Push(asteroids[index])
			goingLeft, goingRight = collisions(goingLeft, goingRight)
		}
	}
	stackUnload := make([]int, 0, len(asteroids))
	for !goingRight.IsEmpty() {
		val, _ := goingRight.Top()
		stackUnload = append(stackUnload, val)
		goingRight.Pop()
	}
	for !goingLeft.IsEmpty() {
		val, _ := goingLeft.Top()
		stackUnload = append(stackUnload, val)
		goingLeft.Pop()
	}
	slices.Reverse(stackUnload) // comes out of stack backward need to reverse it
	results = append(results, stackUnload...)
	return results
}

func collisions(goingLeft Stack, goingRight Stack) (Stack, Stack) {
	for !goingLeft.IsEmpty() && !goingRight.IsEmpty() {
		left, _ := goingLeft.Top()
		right, _ := goingRight.Top()
		if -1*left > right {
			goingRight.Pop()
		} else if -1*left == right {
			goingLeft.Pop()
			goingRight.Pop()
		} else {
			goingLeft.Pop()
		}
	}
	return goingLeft, goingRight
}
