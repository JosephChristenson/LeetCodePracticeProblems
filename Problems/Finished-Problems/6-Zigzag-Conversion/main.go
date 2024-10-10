package main

import "fmt"

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
	s = "PAYPALISHIRING"
	numRows = 4
	fmt.Println(convert(s, numRows))
	s = "A"
	numRows = 1
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	result, input := []byte{}, []byte(s)

	skipFactor := numRows*2 - 1 //each row doubles back except the top and the bottom row. We account for the bottom by subtracting 1
	row, index := 0, 0
	skipPosition := 0

	gettwo := false // after the first pass we are left with peaks with the tips cut off. So we need to grab two values since they will be next to each other before continuing

	/*

		*   *   *   *
		A P L S I I G
		Y   I   R

	*/

	for row < numRows {
		if index == len(input) {
			index = 0
			skipPosition = 0
			row++
			gettwo = false
		} else {
			if skipPosition == 0 {
				if gettwo {
					skipPosition = 1
					gettwo = false
				} else {
					if row > 0 {
						gettwo = true
					}
					skipPosition = (skipFactor - row*2) - 1
				}
				result = append(result, input[index])
				input = append(input[:index], input[index+1:]...) //remove index. Don't increment index
			} else {
				index++
			}
			skipPosition--
		}
	}
	result = append(result, input...)

	return string(result)
}
