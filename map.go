package main

import (
	"fmt"
	"strings"
)

func WordCounts(s string) map[string]int {
	res := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		res[word]++
	}
	return res
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCounts(s)
	fmt.Println(w)
}
