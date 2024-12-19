package SaveId

import (
	"fmt"
	"testing"
)

func TestSaveId(t *testing.T) {

	type saveId struct {
		input          string
		expectedResult bool
	}

	inputs := []saveId{
		{input: "cpf", expectedResult: true},
		{input: "email", expectedResult: false},
		{input: "telefone", expectedResult: true},
		{input: "carrinho", expectedResult: false},
		{input: "boneca", expectedResult: false},
	}

	for _, v := range inputs {
		fmt.Println("testing", v.input)
		result := Save(v.input)
		if result != v.expectedResult {
			t.Errorf("Expected %v but got %v", v.expectedResult, result)
		}
	}

}
