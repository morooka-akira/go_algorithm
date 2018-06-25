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
	for i := 1; i <= N-1; i++ {
		str := fmt.Sprintf("%v", A)
		str = strings.Trim(str, "[]")
		fmt.Println(str)
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	stdOutList(&A)
}
