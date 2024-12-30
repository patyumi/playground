package main

import (
	"fmt"
	"time"
)

func dynamicArray(n int32, queries [][]int32) []int32 {
	lastAnswer := int32(0)
	var answers []int32
	arr := make([][]int32, n)

	for i := 0; i < len(queries); i++ {
		queryType := queries[i][0]
		x := queries[i][1]
		y := queries[i][2]
		idx := ((x ^ lastAnswer) % n)

		if queryType == 1 {
			arr[idx] = append(arr[idx], y)
		}

		if queryType == 2 {
			lastAnswer = arr[idx][y%int32(len(arr[idx]))]
			answers = append(answers, lastAnswer)
		}

	}

	return answers
}

func main() {
	module := "Dynamic Array"
	fmt.Printf("Hacker Rank - %s\n", module)

	startTime := time.Now()

	number := 2
	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	result := dynamicArray(int32(number), queries)
	fmt.Println(result)

	elapsedTime := time.Since(startTime)
	fmt.Printf("\n%s completed. Took %d.", module, elapsedTime)
}
