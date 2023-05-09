package main

import "fmt"

func contane(ss []string) string {
	var st string = ""
	if ss == nil {
		return ""
	} else {
		for i, s := range ss {
			if i != len(ss)-1 {
				st = st + s + ", "
			} else {
				st = st + s
			}
		}
	}
	return st
}

func main() {
	ss := contane([]string{"Oi", "Tchau", "Como está", "Estou bem", "Foi um prazer"})
	fmt.Println(ss)
}
