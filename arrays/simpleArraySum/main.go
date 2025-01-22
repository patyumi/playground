package main

import "fmt"

func simpleArraySum(ar []int32) int32 {

	var sum int32

	for _, v := range ar {
		sum += v
	}

	return sum

}

func main() {
	a := []int32{1, 2, 3}

	res := simpleArraySum(a)
	fmt.Println(res)
}
