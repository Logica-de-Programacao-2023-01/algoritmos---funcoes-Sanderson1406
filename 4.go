package main

import (
	"fmt"
	"sort"
)

func maiorvalor(num []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	return num[1]
}

func main() {
	num := maiorvalor([]int{2, 4, 7, 56, 55, 8})
	fmt.Println(num)
}
