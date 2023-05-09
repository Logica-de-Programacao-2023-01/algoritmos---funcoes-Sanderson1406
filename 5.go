package main

import "fmt"

func slice(sli []int, posi int) int {
	if posi > len(sli) {
		return -1
	}
	return sli[posi]
}

func main() {
	var posi int
	fmt.Println("Escolha uma posição: ")
	fmt.Scan(&posi)
	sli := slice([]int{2, 4, 7, 14, 55, 8}, posi)
	fmt.Println(sli)
}
