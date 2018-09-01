package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StdInInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	n, _ := strconv.Atoi(input)
	return n
}

func SliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, _ := strconv.Atoi(a)
		si = append(si, i)
	}
	return si
}

func StdInSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return SliceAtoi(strings.Split(stringInput, " "))
}

func Partision(A []int, p int, r int) int {
	x := A[r]  // 最後の要素を基準値とする
	i := p - 1 // 一番左の最小値添字
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			t := A[i]
			A[i] = A[j]
			A[j] = t
		}
	}
	c := i + 1
	t2 := A[c]
	A[c] = A[r]
	A[r] = t2
	return c
}

func Display(A []int, center int) {
	for i := 0; i < len(A); i++ {
		if i == center {
			fmt.Printf("[%d]", A[i])
		} else {
			fmt.Printf("%d", A[i])
		}
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	N := StdInInt()
	A := StdInSlice()
	c := Partision(A, 0, N-1)
	Display(A, c)
}
