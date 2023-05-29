package main

import "fmt"

func primo(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primotr(n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("Número insuficiente")
	}
	nms := make([]int, 0)
	for i := 2; i <= n; i++ {
		if primo(i) {
			nms = append(nms, i)
		}
	}
	return nms, nil
}

func main() {
	var n int
	fmt.Println("Digite um numero: ")
	fmt.Scanln(&n)
	s, err := primotr(n)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(s)
}
