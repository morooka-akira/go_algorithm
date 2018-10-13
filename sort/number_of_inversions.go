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
	a := make([]int, n1+n2)
	c := 0
	for i := 0; i < n1; i++ {
		l[i] = A[left+i]
		a[c] = A[left+i]
		c++
	}
	for i := 0; i < n2; i++ {
		r[i] = A[mid+i]
		a[c] = A[mid+i]
		c++
	}
	l[n1] = math.MaxInt64
	r[n2] = math.MaxInt64
	CountInverse(a, n1+n2)
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if l[i] <= r[j] {
			A[k] = l[i]
			i++
		} else {
			A[k] = r[j]
			j++
		}
	}
}

// 反転数を求める
func CountInverse(a []int, n int) {
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[i] {
				Cnt++
			}
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
