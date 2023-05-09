package main

import (
	"fmt"
	"sort"
)

func order(eles []int) []int {
	sort.Ints(eles)
	return eles
}

func main() {
	eles := order([]int{1, 2, 8, 12, 15, 23, 33, 51, 45, 13})
	fmt.Println(eles)
}
