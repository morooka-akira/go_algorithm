package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func stdInSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return strings.Split(stringInput, " ")
}

func stdOutList(list *[]string) {
	str := fmt.Sprintf("%v", *list)
	str = strings.Trim(str, "[]")
	fmt.Println(str)
}

func BubbleSort(A []string, N int) {
	i := 0
	s := 0
	for i < N {
		for j := N - 1; j >= 1; j-- {
			if getNum(A[j]) < getNum(A[j-1]) {
				tmp := A[j]
				A[j] = A[j-1]
				A[j-1] = tmp
				s++
			}
		}
		i++
	}
}

func getNum(str string) int {
	n, _ := strconv.Atoi(str[1:2])
	return n
}

func getSuit(str string) string {
	return str[0:1]
}

func SelectionSort(A []string, N int) {
	i := 0
	s := 0
	for i < N-1 {
		min := i
		for j := i; j < N; j++ {
			if getNum(A[j]) < getNum(A[min]) {
				min = j
			}
		}
		if getNum(A[min]) != getNum(A[i]) {
			tmp := A[min]
			A[min] = A[i]
			A[i] = tmp
			s++
		}
		i++
	}
}

func isStable(A1 []string, A2 []string, N int) bool {
	for i := 0; i < N; i++ {
		if getSuit(A1[i]) != getSuit(A2[i]) {
			return false
		}
	}
	return true
}

func main() {
	N := stdInInt()
	A := stdInSlice()
	A2 := make([]string, N)
	copy(A2, A)
	BubbleSort(A, N)
	stdOutList(&A)
	fmt.Println("Stable")
	SelectionSort(A2, N)
	stdOutList(&A2)
	if isStable(A, A2, N) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}
