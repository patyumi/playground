package main

import (
	"fmt"
	SaveId "playground/saveId"
)

func main() {
	fmt.Println("Hello")

	result := SaveId.Save("cpf")
	fmt.Println("the result is ", result)
}
