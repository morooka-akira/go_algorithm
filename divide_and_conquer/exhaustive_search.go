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

func Solve(i int, m int) bool {
	if m == 0 { // m=0は解決できた
		return true
	}
	if m < 0 { // mが0以下は解決できないのでそれ以上みない
		return false
	}
	if i >= n { // 配列の最後まで見たら解決できない
		return false
	}
	// 左側はA[i]を選択しなかったケース、右は選択したケース
	return Solve(i+1, m) || Solve(i+1, m-A[i])
}

var n int
var A []int

func main() {
	n = StdInInt()
	A = StdInSlice()
	q := StdInInt()
	m := StdInSlice()
	for i := 0; i < q; i++ {
		if Solve(0, m[i]) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
