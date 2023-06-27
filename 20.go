package main

import (
	"errors"
	"fmt"
)

func filterStrings(strSlice []string) ([]string, error) {
	if len(strSlice) == 0 {
		return nil, errors.New("slice vazio")
	}

	var result []string
	for _, str := range strSlice {
		if len(str) > 5 {
			result = append(result, str)
		}
	}
	return result, nil
}

func main() {
	strSlice := []string{"apple", "banana", "orange", "kiwi", "grapefruit"}
	result, err := filterStrings(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
