package main

import "fmt"

func StdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func main() {
	n := StdInInt()
	fmt.Println(n * n * n)
}
