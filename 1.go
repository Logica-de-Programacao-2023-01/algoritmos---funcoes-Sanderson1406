package main

import "fmt"

func media(num []int) int {
	var soma int
	for i := 0; i < len(num); i++ {
		soma += num[i]
	}
	res := soma / int(len(num))
	return res
}

func main() {
	num := media([]int{5, 13, 4, 1, 6, 8})
	fmt.Println(num)
}
