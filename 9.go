package main

import (
	"errors"
	"fmt"
	"strings"
)

func extractWords(str string) ([]string, error) {
	if len(str) == 0 {
		return nil, errors.New("string vazia")
	}
	words := strings.Fields(str)
	return words, nil
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	words, err := extractWords(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", words)
	}
}
