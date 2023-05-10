package main

import (
	"errors"
	"fmt"
	"strings"
)

func set(s string) (string, error) {
	if s == "" {
		return "", errors.New("string vazia")
	}
	subst := strings.NewReplacer(
		"a", "*", "e", "*", "i", "*", "o", "*", "u", "*",
		"A", "*", "E", "*", "I", "*", "O", "*", "U", "*",
	)
	return subst.Replace(s), nil
}

func main() {
	s, err := set("OI, como vai?")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
