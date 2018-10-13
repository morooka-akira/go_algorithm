package main

import (
	"bufio"
	"fmt"
	"math"
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

var Cnt int = 0

func merge(A []int, left int, mid int, right int) {
	// 要素数を出す
	n1 := mid - left
	n2 := right - mid
	l := make([]int, n1+1)
	r := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		l[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = A[mid+i]
	}
	l[n1] = math.MaxInt64
	r[n2] = math.MaxInt64
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if l[i] <= r[j] {
			A[k] = l[i]
			i++
		} else {
			// leftのほうが大きい場合はn1-iで反点数が求められる
			Cnt += n1 - i
			A[k] = r[j]
			j++
		}
	}
}

func mergeSort(A []int, left int, right int) {
	// 1つだとマージできないので最低2個以上でマージ
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func main() {
	N := StdInInt()
	S := StdInSlice()
	mergeSort(S, 0, N)
	fmt.Println(Cnt)
}
