package main

import "fmt"

func main() {
	var inputArray = []int{2, 3, 5, 6, 1, 7, 3, 2}
	numberMap := make(map[int]int)
	for index, value := range inputArray {
		// fmt.Printf("Index %d value %d\n", index, value)
		fmt.Printf("current map %d %d\n", index, numberMap)
		if _, ok := numberMap[value]; ok {
			fmt.Printf("The duplicate element is %d\n", value)
			return
		} else {
			numberMap[value] = value
		}
	}
	fmt.Printf("Completed processing, no duplicate found")
}
