package main

import (
	"fmt"
)

func main() {
	prev, next := 0, 1
	i := 0
	for i < 10 {
		fmt.Println(prev)
		prev, next = next, prev+next
		i++
	}
}
