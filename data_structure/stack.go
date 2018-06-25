package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Max int = 1000

var stack = make([]int, Max)
var top int = 0

func StdInSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return strings.Split(stringInput, " ")
}

func main() {
	in := StdInSlice()
	for _, v := range in {
		i, e := strconv.Atoi(v)
		if e == nil {
			Push(i)
		} else {
			v1, _ := Pop()
			v2, _ := Pop()
			if v == "*" {
				Push(v1 * v2)
			}
			if v == "+" {
				Push(v1 + v2)
			}
			if v == "-" {
				Push(v2 - v1)
			}
		}
	}
	r, _ := Pop()
	fmt.Println(r)
}

func Initialize() {
	top = 0
}

func IsEmpty() bool {
	return top == 0
}

func IsFull() bool {
	return top == Max
}

func Push(x int) (int, error) {
	if IsFull() {
		return 0, errors.New("Stack overflow")
	}
	top++
	stack[top] = x
	return x, nil
}

func Pop() (int, error) {
	if IsEmpty() {
		return 0, errors.New("Stack empty")
	}
	top--
	return stack[top+1], nil
}
