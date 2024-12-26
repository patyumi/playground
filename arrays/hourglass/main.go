package main

import (
	"fmt"
	"time"
)

func hourglassSum(arr [][]int32) int32 {
	var (
		results []int32
		result  int32
	)

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr); j++ {

			if i+3 > len(arr) || j+3 > len(arr) {
				break
			}

			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			results = append(results, sum)

		}
	}

	result = results[0]
	for i := range results {
		if results[i] >= result {
			result = results[i]
		}
	}

	return result
}

func main() {
	fmt.Println("HackerRanks - Hourglass")
	startTime := time.Now()

	arr := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	// hourglass body
	// fmt.Println(arr[i][j], arr[i][j+1], arr[i][j+2])
	// fmt.Println(arr[i+1][j+1])
	// fmt.Println(arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])

	output := hourglassSum(arr)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Result: %d\n\nCompleted in %v", output, elapsedTime)
}
