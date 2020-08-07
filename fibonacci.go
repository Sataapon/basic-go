package main

import (
	"fmt"
)

func fibonacci() func() int {
	curr, next := 0, 1
	return func() int {
		curr, next = next, curr+next
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
