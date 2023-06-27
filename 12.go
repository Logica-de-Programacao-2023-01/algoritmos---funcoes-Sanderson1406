package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filterUpperCase(strSlice []string) (string, error) {
	if len(strSlice) == 0 {
		return "", errors.New("slice vazio")
	}
	var result strings.Builder
	for _, str := range strSlice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String()), nil
}

func main() {
	strSlice := []string{"Prato", "video", "Ceub", "mençao", "Aluno"}
	result, err := filterUpperCase(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
