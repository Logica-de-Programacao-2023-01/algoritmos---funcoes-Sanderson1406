package main

import "fmt"

func sring(ss []string) string {
	st := ""
	for _, s := range ss {
		st += s
	}
	return st
}

func main() {
	ss := sring([]string{"Bom dia, ", "Boa tarde, ", "Boa noite"})
	fmt.Println(ss)
}
