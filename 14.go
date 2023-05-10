package main

import "fmt"

func juntar(eles []int, eles2 []int) []int {
	newele := []int{}
	if len(eles) <= 0 {
		return []int{}
	} else if len(eles2) <= 0 {
		return []int{}
	} else {
		for i := 0; i < len(eles); i++ {
			for s := 0; s < len(eles2); s++ {
				if eles[i] == eles2[s] {
					newele = append(newele, eles[i])
				}
			}
		}
	}
	return newele
}

func main() {
	eles := juntar([]int{1, 2, 8, 12, 15, 23, 33, 51, 45, 13}, []int{2, 8, 9, 7, 15, 13, 45})
	fmt.Println(eles)
}
