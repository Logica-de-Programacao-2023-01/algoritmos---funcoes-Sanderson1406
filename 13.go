package main

import "fmt"

func soma(ele int) int {
	var soma int
	if ele <= 0 {
		return 0
	} else {
		for ele > 0 {
			digit := ele % 10
			soma += digit
			ele /= 10
		}
	}
	return soma
}

func main() {
	var dit int
	fmt.Println("Digite um número: ")
	fmt.Scan(&dit)
	ele := soma(dit)
	fmt.Println(ele)
}
