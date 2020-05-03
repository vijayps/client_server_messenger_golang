package main

import (
	"fmt"
)

func main() {
	binArray := []int{0, 1, 1, 1, 0, 0, 0, 1, 1, 0}
	binarySort(binArray)
}

func binarySort(binArray []int) {
	var i, j int
	i = 0
	j = len(binArray) - 1

	fmt.Println("i index", i)
	fmt.Println("last index ", j)

	for i <= j {
		for i <= j {
			if binArray[i] == 0 {
				i++
			}
			break
		}
		fmt.Println("found one 1")

		for j >= i {
			if binArray[j] == 1 {
				j--
			}
			break
		}
		fmt.Println("found one 0")

		// swap
		binArray[i] = 0
		binArray[j] = 1
		i++
		j--
	}
	fmt.Println("sorted array ", binArray)
}
