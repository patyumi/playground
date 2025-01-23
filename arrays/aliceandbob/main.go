package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {

	var (
		result []int32
		alice  int32
		bob    int32
	)

	for i := range a {

		if a[i] > b[i] {
			alice++
		}
		if a[i] < b[i] {
			bob++
		}
	}

	result = append(result, alice, bob)
	return result

}

func main() {

	alice := []int32{17, 28, 30}
	bob := []int32{99, 16, 8}

	result := compareTriplets(alice, bob)
	fmt.Println(result)
}
