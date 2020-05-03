package main

import "fmt"

func main() {
	inputArray := [10]int{2, 5, 6, 0, -1, 2, 3, 4, 5, 6}
	// for index, value := range inputArray {
	// 	fmt.Printf("index %d Value %d \n", index, value)
	// }

	var startIndex, endIndex, prevNumber, prevStartIndex, length, maxLength int
	prevNumber = inputArray[0]
	startIndex = 0
	endIndex = 0
	length = 1
	maxLength = 1
	prevStartIndex = 0
	for i := 1; i < 10; i++ {
		if prevNumber < inputArray[i] {
			prevNumber = inputArray[i]
			length++
		} else {
			if length > maxLength {
				fmt.Printf("Found new largest sub-array\n")
				maxLength = length
				startIndex = prevStartIndex
				endIndex = i - 1
			}
			prevStartIndex = i
			prevNumber = inputArray[i]
		}
	}

	if length > maxLength {
		fmt.Printf("Found new largest sub-array\n")
		maxLength = length
		startIndex = prevStartIndex
		endIndex = len(inputArray) - 1
	}

	fmt.Printf("Largest sub-Array length is %d\n", maxLength)
	fmt.Printf("start index is %d, end index is %d\n", startIndex, endIndex)
	for i := startIndex; i <= endIndex; i++ {
		fmt.Printf("%d ", inputArray[i])
	}
	fmt.Println()
}
