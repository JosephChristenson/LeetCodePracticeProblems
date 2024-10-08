package main

import "fmt"

// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

// Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.

func main() {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	gas, cost = []int{2, 3, 4}, []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	score, start, highestcost := 0, 0, 0
	for index, val := range gas {
		if index+1 == len(gas) {
			score = score + val - cost[0]
		} else {
			score = score + val - cost[index+1]
		}
		if score <= highestcost {
			highestcost = score
			start = index + 1 // wanted index is the one after the worst cost.
		}

		fmt.Println(score)
	}

	if score >= 0 {
		if start == len(gas) {
			start = 0
		}
		return start + 1 // correct for base 0 index
	}
	return -1
}
