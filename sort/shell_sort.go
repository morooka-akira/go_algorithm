package main

import (
	"fmt"
	"strings"
)

func StdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func StdOutList(list *[]int) {
	str := fmt.Sprintf("%v", *list)
	str = strings.Trim(str, "[]")
	fmt.Println(str)
}

var num int = 0

func main() {
	N := StdInInt()
	A := make([]int, 0)
	for i := 0; i < N; i++ {
		A = append(A, StdInInt())
	}
	ShellSort(A, N)
	fmt.Println(num)
	for i := 0; i < N; i++ {
		fmt.Println(A[i])
	}
}

func ShellSort(A []int, N int) {
	G := []int{}
	for i := 1; i <= N; i = i*3 + 1 {
		G = append([]int{i}, G...)
	}
	m := len(G)
	for i := 0; i < m; i++ {
		InsertSort(A, N, G[i])
	}
	fmt.Println(m)
	StdOutList(&G)
}

func InsertSort(A []int, N int, g int) {
	for i := g; i <= N-1; i++ {
		v := A[i]
		j := i - g
		for j >= 0 && A[j] > v {
			A[j+g] = A[j]
			j -= g
			num++
		}
		A[j+g] = v
	}
}
