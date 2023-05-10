package main

import "fmt"

func veri(ele int) bool {
	primo := true
	for i := 2; i*i <= ele; i++ {
		if ele <= 0 {
			fmt.Println("Negativo")
			break
		} else if ele%i == 0 {
			primo = false
		} else {
			primo = true
		}
	}
	return primo
}

func main() {
	var ele int
	fmt.Println("Informe um número: ")
	fmt.Scan(&ele)
	resul := veri(ele)
	fmt.Println(resul)
}
