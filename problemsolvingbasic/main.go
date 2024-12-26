package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("[HackerRank] Problem Solving (Basic)")

	senha := "43Ah*ck0rr0nk"

	var dSenha []string
	countNumber := 0

	for i := (len(senha) - 1); i != 0; i-- {
		var aSenha []string

		k, _ := regexp.MatchString("[a-z]", string(senha[i]))
		j, _ := regexp.MatchString("[A-Z]", string(senha[i-1]))
		number, _ := regexp.MatchString("[0-9]", string(senha[i]))
		cha, _ := regexp.MatchString("*", string(senha[i]))

		if j && k {

			aSenha = append(aSenha, string(senha[i]))
			aSenha = append(aSenha, string(senha[i-1]))

			dSenha = append(aSenha, dSenha...)

			i -= 2
			continue

		}

		if cha {
			continue
		}

		if number {
			aSenha = append(aSenha, string(senha[countNumber]))
			dSenha = append(aSenha, dSenha...)
			countNumber++
			continue
		}

		aSenha = append(aSenha, string(senha[i]))
		dSenha = append(aSenha, dSenha...)

	}

	mySenha := strings.Join(dSenha, "")
	fmt.Println(mySenha)
	// Test 1 - Password Decryption
	// h A c k 3 r r 4 n k
	// 0 1 - - - 5 6 - 8 9
	// 4 3 A h * c k 0 r r 0 n k
	// 0 1 2 3 - 5 6 7 8 9 - - 12

	// h A c k 3 r r 4 k n

	// começa com s[i] = 0
	// se s[i] é lowercase e o próximo caracter s[i+1] é uppercase, troca eles e
	// adiciona um * depois dos dois, e aí vai pra i+2
	// Ah*

	// se s[i] é um número, troca ele por 0, coloca o número original no começo
	// move para i+1

	// move o caracter atual pra i+1

	// pare se i for maior ou igual ao tamanho da string, se não, volta pro passo 2

	// Sample test - FizzBuzz

	// var number int32
	// number = 15

	// for i := 1; i <= int(number); i++ {

	// 	if (i%3 == 0) && (i%5 == 0) {
	// 		fmt.Println("FizzBuzz")
	// 	}
	// 	if (i%3 == 0) && (i%5 != 0) {
	// 		fmt.Println("Fizz")
	// 	}
	// 	if (i%5 == 0) && (i%3 != 0) {
	// 		fmt.Println("Buzz")
	// 	}
	// 	if (i%3 != 0) && (i%5 != 0) {
	// 		fmt.Println(i)
	// 	}

	// }

}
