package main

import (
	"errors"
	"fmt"
)

func square(x int) int {
	return x * x
}

func applyFunction(numbers []int, f func(int) int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("slice vazio")
	}
	sum := 0
	for _, num := range numbers {
		result := f(num)
		sum += result
	}
	return sum, nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result, err := applyFunction(numbers, square)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
