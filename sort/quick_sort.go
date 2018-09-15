package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Mark string
	Num  int
}

func StdInInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	n, _ := strconv.Atoi(input)
	return n
}

func StdInCard() Card {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.Split(scanner.Text(), " ")
	num, _ := strconv.Atoi(a[1])
	return Card{a[0], num}
}

func Partision(A []Card, p int, r int) int {
	x := A[r]  // 最後の要素を基準値とする
	i := p - 1 // 一番左の最小値添字
	for j := p; j < r; j++ {
		if A[j].Num <= x.Num {
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

func QuickSort(A []Card, p int, r int) {
	if p >= r {
		return
	}
	c := Partision(A, p, r)
	QuickSort(A, p, c-1)
	QuickSort(A, c+1, r)
}

func Merge(A []Card, left int, mid int, right int) {
	// 要素数を出す
	n1 := mid - left
	n2 := right - mid
	l := make([]Card, n1+1)
	r := make([]Card, n2+1)
	for i := 0; i < n1; i++ {
		l[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = A[mid+i]
	}
	l[n1].Num = math.MaxInt64
	r[n2].Num = math.MaxInt64
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if l[i].Num <= r[j].Num {
			A[k] = l[i]
			i++
		} else {
			A[k] = r[j]
			j++
		}
	}
}

func MergeSort(A []Card, left int, right int) {
	// 1つだとマージできないので最低2個以上でマージ
	if left+1 < right {
		mid := (left + right) / 2
		MergeSort(A, left, mid)
		MergeSort(A, mid, right)
		Merge(A, left, mid, right)
	}
}

func Display(A []Card) {
	for i := 0; i < len(A); i++ {
		fmt.Printf("%s %d\n", A[i].Mark, A[i].Num)
	}
}

func JudgeStable(Merge []Card, Quick []Card) {
	for i := 0; i < len(Merge); i++ {
		if Merge[i].Mark != Quick[i].Mark {
			println("Not stable")
			return
		}
	}
	println("Stable")
}

func main() {
	N := StdInInt()
	quickA := make([]Card, N)
	mergeA := make([]Card, N)
	for i := 0; i < N; i++ {
		c := StdInCard()
		quickA[i] = c
		mergeA[i] = c
	}
	MergeSort(mergeA, 0, N)
	QuickSort(quickA, 0, N-1)
	JudgeStable(mergeA, quickA)
	Display(quickA)
}
