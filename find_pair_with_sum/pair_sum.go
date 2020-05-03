package main

import (
	"fmt"
)

func main() {
	// param := os.Args

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// var array = param[1]
	// var sum = param[2]

	myArray := [7]int{1, 4, 3, 7, 8, 0, 89}
	var sum int = 7

	fmt.Println(myArray)
	fmt.Println(sum)
	// fmt.Println(reflect.TypeOf(myArray))
	// fmt.Println(reflect.TypeOf(sum))

	for index, value := range myArray {
		// fmt.Println(index, value)
		var reqNum = sum - value
		fmt.Println("required number: ", reqNum)
		// fmt.Println(reqNum)
		for i := index + 1; i < len(myArray); i++ {
			if reqNum == myArray[i] {
				fmt.Println("Two number are: ", value, myArray[i])
			}
		}
	}

}
