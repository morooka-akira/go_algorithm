package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

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
	buf := make([]byte, 10000)
	scanner.Buffer(buf, 10000000)
	scanner.Scan()
	stringInput := scanner.Text()
	A, _ := SliceAtoi(strings.Split(stringInput, " "))
	return A
}

func CountingSort(A []int, B []int, k int, n int) {
	// 0以上k以下の値を格納する配列を作成
	var C []int = make([]int, k+1)
	// 各要素の数を数える
	for i := 0; i < n; i++ {
		C[A[i]]++
	}
	// C[i]にi以下の出現数を記録する
	for i := 1; i < k; i++ {
		C[i] = C[i] + C[i-1]
	}
	// Cが指すindexにAを挿入していく
	for i := n - 1; i >= 0; i-- {
		B[C[A[i]]-1] = A[i]
		C[A[i]]--
	}
}

func StdOutList(list *[]int) {
	str := fmt.Sprintf("%v", *list)
	str = strings.Trim(str, "[]")
	fmt.Println(str)
}

var MaxNum = 10000

func main() {
	N := StdInInt()
	A := StdInNumSlice()
	var B []int = make([]int, N)
	CountingSort(A, B, MaxNum, N)
	StdOutList(&B)
}
