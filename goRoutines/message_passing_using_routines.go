package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	go userOne(10, &wg)
	wg.Add(1)

	go userTwo("Vijay", &wg)
	wg.Add(1)

	// var anyKey, firstKey string
	// fmt.Println("Enter your name:")
	// fmt.Scanf("%s%s", &anyKey, &firstKey)
	// // fmt.Scanln(&anyKey)
	// fmt.Println("Key pressed ", anyKey, firstKey)
	wg.Wait()
	fmt.Println("Done with routines")
}

func userOne(noOfTimes int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < noOfTimes; i++ {
		fmt.Println("Index call ", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func userTwo(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Name call ", i, name)
		time.Sleep(1 * time.Millisecond)
	}
}
