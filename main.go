package main

import (
	"fmt"
	"runtime"
)

// func main() {
// 	argsWithProg := os.Args
// 	argsWithoutProg := os.Args[1:]
// 	arg := os.Args[3]

// 	fmt.Println(argsWithProg)
// 	fmt.Println(argsWithoutProg)
// 	fmt.Println(arg)
// }

func main() {
	countGoRoutines()
}

func countGoRoutines() {
	fmt.Printf("Number of go goRoutines: %d\n", runtime.NumGoroutine())
}
