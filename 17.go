package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func sortStrings(strSlice []string) (string, error) {
	if len(strSlice) == 0 {
		return "", errors.New("slice vazio")
	}
	sort.Strings(strSlice)
	result := strings.Join(strSlice, " ")
	return result, nil
}

func main() {
	strSlice := []string{"banana", "abacaxi", "laranja", "uva"}
	result, err := sortStrings(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
