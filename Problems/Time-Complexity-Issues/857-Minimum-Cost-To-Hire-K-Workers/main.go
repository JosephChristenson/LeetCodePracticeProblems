package main

import (
	"container/heap"
	"fmt"
)

// There are n workers. You are given two integer arrays quality and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

// We want to hire exactly k workers to form a paid group. To hire a group of k workers, we must pay them according to the following rules:

// Every worker in the paid group must be paid at least their minimum wage expectation.
// In the group, each worker's pay must be directly proportional to their quality. This means if a worker’s quality is double that of another worker in the group, then they must be paid twice as much as the other worker.
// Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within 10-5 of the actual answer will be accepted.

func main() {
	quality, wage := []int{10, 20, 5}, []int{70, 50, 30}
	k := 2
	fmt.Println(mincostToHireWorkers(quality, wage, k))
	quality, wage = []int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}
	k = 3
	fmt.Println(mincostToHireWorkers(quality, wage, k))
	quality, wage = []int{25, 68, 35, 62, 52, 57, 35, 83, 40, 51}, []int{147, 97, 251, 129, 438, 443, 120, 366, 362, 343}
	k = 6
	fmt.Println(mincostToHireWorkers(quality, wage, k))
	quality, wage = []int{14, 56, 59, 89, 39, 26, 86, 76, 3, 36}, []int{90, 217, 301, 202, 294, 445, 473, 245, 415, 487}
	k = 2
	fmt.Println(mincostToHireWorkers(quality, wage, k))
	quality, wage = []int{37, 32, 14, 14, 23, 31, 82, 96, 81, 96, 22, 17, 68, 3, 88, 59, 54, 23, 22, 77, 61, 16, 46, 22, 94, 50, 29, 46, 7, 33, 22, 99, 31, 99, 75, 67, 95, 54, 31, 48, 44, 96, 99, 20, 51, 54, 18, 85, 25, 84}, []int{453, 236, 199, 359, 107, 45, 150, 433, 32, 192, 433, 94, 113, 200, 293, 31, 48, 27, 15, 32, 295, 97, 199, 427, 90, 215, 390, 412, 475, 131, 122, 398, 479, 142, 103, 243, 86, 309, 498, 210, 173, 363, 449, 135, 353, 397, 105, 165, 165, 62}
	k = 20
	fmt.Println(mincostToHireWorkers(quality, wage, k))
}
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {

	workers := &FloatHeap{}
	storage := &FloatHeap{}
	ratio := &RatioHeap{}

	for index := range quality {
		heap.Push(workers, []float64{float64(quality[index]), float64(wage[index])})
		heap.Push(ratio, float64(wage[index])/float64(quality[index]))
	}

	lowestTotal := float64(0)

	for workers.Len() >= k && ratio.Len() > 0 {
		currentRatio := heap.Pop(ratio).(float64)
		acceptedCount := 0
		total := float64(0)
		for workers.Len() > 0 {
			compare := heap.Pop(workers).([]float64)
			if compare[0]*currentRatio >= compare[1] {
				if acceptedCount < k {
					acceptedCount++
					total = total + compare[0]*currentRatio
				}
				heap.Push(storage, compare)
			}
		}

		if acceptedCount == k {
			if lowestTotal == 0 {
				lowestTotal = total
			} else if lowestTotal > total {
				lowestTotal = total
			}
		} else {
			break
		}
		workers = storage
		storage = &FloatHeap{}

	}
	return lowestTotal
}

type FloatHeap [][]float64

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h FloatHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h *FloatHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *FloatHeap) Push(x any) { *h = append(*h, x.([]float64)) }

type RatioHeap []float64

func (h RatioHeap) Len() int           { return len(h) }
func (h RatioHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h RatioHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *RatioHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *RatioHeap) Push(x any) { *h = append(*h, x.(float64)) }

// func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
// 	workers := &FloatHeap{}
// 	storage := &FloatHeap{}
// 	for index := range quality {
// 		heap.Push(workers, []float64{float64(quality[index]), float64(wage[index]), float64(wage[index]) / float64(quality[index])})
// 	}
// 	heap.Init(workers)
// 	lowestCost := float64(0)
// 	alternate := false
// 	worker := []float64{0, 0, 0}

// 	for workers.Len() >= k || storage.Len() >= k {
// 		costFinder := &SumHeap{}
// 		if alternate {
// 			worker = heap.Pop(storage).([]float64)
// 		} else {
// 			worker = heap.Pop(workers).([]float64)
// 		}
// 		heap.Push(costFinder, worker[1])
// 		if alternate {
// 			for storage.Len() > 0 {
// 				comparison := heap.Pop(storage).([]float64)
// 				costToHire := worker[2] * comparison[0]
// 				if costToHire >= comparison[1] {
// 					heap.Push(costFinder, costToHire)
// 					heap.Push(workers, comparison)
// 				}
// 			}
// 		} else {
// 			for workers.Len() > 0 {
// 				comparison := heap.Pop(workers).([]float64)
// 				costToHire := worker[2] * comparison[0]
// 				if costToHire >= comparison[1] {
// 					heap.Push(costFinder, costToHire)
// 					heap.Push(storage, comparison)
// 				}
// 			}
// 		}
// 		totalCosts := float64(0)
// 		if costFinder.Len() >= k {
// 			loop := 0
// 			for loop < k {
// 				loop++
// 				totalCosts = totalCosts + heap.Pop(costFinder).(float64)
// 			}
// 			if lowestCost == 0 {
// 				lowestCost = totalCosts
// 			} else if lowestCost > totalCosts {
// 				lowestCost = totalCosts
// 			}
// 		}
// 		alternate = !alternate
// 	}
// 	return lowestCost
// }

// type FloatHeap [][]float64

// func (h FloatHeap) Len() int           { return len(h) }
// func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h FloatHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
// func (h *FloatHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
// func (h *FloatHeap) Push(x any) { *h = append(*h, x.([]float64)) }

// type SumHeap []float64

// func (h SumHeap) Len() int           { return len(h) }
// func (h SumHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h SumHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h *SumHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
// func (h *SumHeap) Push(x any) { *h = append(*h, x.(float64)) }
