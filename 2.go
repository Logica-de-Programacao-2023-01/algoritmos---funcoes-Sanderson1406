package main

import (
	"fmt"
	"strings"
)

func conatr(str string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	vowelCount := conatr(str)
	fmt.Printf("A quantidade de vogais na string é: %d\n", vowelCount)
}
