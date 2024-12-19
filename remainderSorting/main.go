package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("hello")

	type complete struct {
		Word      string
		Remainder int
	}

	var preresult []complete

	// declara array
	str := []string{"Colorado", "Utah", "Wisconsin", "Oregon"}
	str2 := []int{1, 2, 0, 0}

	// pega o tamanho de cada palavra
	for _, x := range str {

		preresult = append(preresult, complete{
			Word:      x,
			Remainder: len(x) % 3,
		})

		//fmt.Println(preresult)

	}
	fmt.Println(preresult)

	sort.Slice(preresult, func(i, j int) bool {

		if preresult[i].Remainder == preresult[j].Remainder {
			return preresult[i].Word < preresult[j].Word
		}

		return preresult[i].Remainder < preresult[j].Remainder
	})

	fmt.Println(preresult)

	var newarr []string

	for _, x := range preresult {
		newarr = append(newarr, x.Word)
	}

	fmt.Println(newarr)
	// ordena o array
	slices.Sort(str)
	slices.Sort(str2)
	fmt.Println(str)
	fmt.Println(str2)

	// por tamanho crescente
	// por ordem alfabética

}
