package main

import (
	"fmt"
	"log"
	"regexp"
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

func RemoveSymbols(s string) string {
	reg, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, " ")
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCounts(RemoveSymbols(s))
	fmt.Println(w)
}
