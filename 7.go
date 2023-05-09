package main

import "fmt"

func mapper(arr []int, f func(int) int) []int {
	result := []int{}
	for _, ele := range arr {
		result = append(result, f(ele))
	}
	return result
}

func main() {
	result := mapper([]int{1, 2, 3}, func(i int) int {
		return i * 2
	})
	fmt.Println(result)
}
