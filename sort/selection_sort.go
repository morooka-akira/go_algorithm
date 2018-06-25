package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
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

func stdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func stdInNumSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	A, _ := sliceAtoi(strings.Split(stringInput, " "))
	return A
}

func stdOutList(list *[]int) {
	str := fmt.Sprintf("%v", *list)
	str = strings.Trim(str, "[]")
	fmt.Println(str)
}

func main() {
	N := stdInInt()
	A := stdInNumSlice()
	i := 0
	s := 0
	for i < N-1 {
		min := i
		for j := i; j < N; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		if A[min] != A[i] {
			tmp := A[min]
			A[min] = A[i]
			A[i] = tmp
			s++
		}
		i++
	}
	stdOutList(&A)
	fmt.Println(s)
}
