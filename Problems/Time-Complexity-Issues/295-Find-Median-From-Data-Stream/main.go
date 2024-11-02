package main

import (
	"fmt"
	"sort"
)

// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:

// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

func main() {

	commands := []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}
	num := [][]int{{}, {1}, {2}, {}, {3}, {}}
	// runFunc(commands, num)
	// commands = []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	// num = [][]int{{}, {-1}, {}, {-2}, {}, {-3}, {}, {-4}, {}, {-5}, {}}
	// runFunc(commands, num)
	commands = []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	num = [][]int{{}, {1}, {}, {2}, {}, {1}, {}, {4}, {}, {5}, {}, {6}, {}, {7}, {}, {8}, {}, {9}, {}, {10}, {}}
	runFunc(commands, num)

}

func runFunc(input []string, values [][]int) {
	var median MedianFinder

	for i, s := range input {
		if s == "MedianFinder" {
			median = Constructor()
		} else if s == "addNum" {
			median.AddNum(values[i][0])
		} else if s == "findMedian" {
			fmt.Println(median.FindMedian())
		}
	}
}

type MedianFinder struct {
	list []int
}

func Constructor() MedianFinder {
	result := MedianFinder{}
	return result
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.list) > 0 {
		if this.list[0] > num {
			this.list = append([]int{num}, this.list...)
		} else if this.list[len(this.list)-1] < num {
			this.list = append(this.list, num)
		} else {
			this.list = append(this.list, num)
			sort.Slice(this.list, func(i, j int) bool {
				return this.list[i] < this.list[j]
			})
		}
	} else {
		this.list = append(this.list, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.list)%2 == 1 { // is it an odd length array
		return float64(this.list[len(this.list)/2])
	} else {
		return (float64(this.list[len(this.list)/2-1]) + float64(this.list[len(this.list)/2])) / 2
	}
}
