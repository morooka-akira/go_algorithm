package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func StdInNumSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	A, _ := SliceAtoi(strings.Split(stringInput, " "))
	return A
}

func StdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func Search(A []int, n int, key int) bool {
	A = append(A, key)
	i := 0
	for A[i] != key {
		i++
	}
	return i != n
}

func main() {
	n := StdInInt()
	S := StdInNumSlice()
	q := StdInInt()
	T := StdInNumSlice()
	h := 0
	for i := 0; i < q; i++ {
		if Search(S, n, T[i]) {
			h++
		}
	}
	fmt.Fprintln(os.Stdout, h)
}
