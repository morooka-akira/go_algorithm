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

func main() {
	A := StdInNumSlice()
	a := A[0] * A[1]
	s := A[0]*2 + A[1]*2
	fmt.Printf("%v %v\n", a, s)
}
