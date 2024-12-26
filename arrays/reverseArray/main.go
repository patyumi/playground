package main

import (
	"fmt"
	"time"
)

func reverseArray(a []int32) []int32 {
	var output []int32

	for i := len(a) - 1; i >= 0; i-- {
		output = append(output, a[i])
	}

	return output
}

func main() {
	fmt.Println("HackerRanks - Reverse Array")
	startTime := time.Now()

	input := []int32{1, 4, 3, 2}
	output := reverseArray(input)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Array successfully reversed from:\n%v\nto:\n%v\n\nCompleted in: %d", input, output, elapsedTime)
	fmt.Println(output)
}
