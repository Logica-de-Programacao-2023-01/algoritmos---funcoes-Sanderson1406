package main

import "fmt"

func verifi(n int, eles []int) (bool, error) {
	vdd := false
	if len(eles) == 0 {
		return false, fmt.Errorf("Slice vazia")
	} else {
		for _, s := range eles {
			if s == n {
				vdd = true
				break
			} else {
				vdd = false
			}
		}
	}
	return vdd, nil
}

func main() {
	var n int
	fmt.Println("Escolha um número")
	fmt.Scan(&n)
	eles, err := verifi(n, []int{1, 2, 8, 12, 15, 23, 33, 51, 45, 13})
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Verificação:", eles)
	}
}
