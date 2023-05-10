package main

import "fmt"

func par(eles []int) []int {
	newele := []int{}
	if len(eles) <= 0 {
		return []int{}
	} else {
		for i := 1; i < len(eles); i++ {
			if eles[i]%2 == 0 {
				newele = append(newele, eles[i])
			}
		}
	}
	return newele
}

func main() {
	eles := par([]int{1, 2, 8, 12, 15, 23, 33, 51, 45, 13})
	fmt.Println(eles)
}
