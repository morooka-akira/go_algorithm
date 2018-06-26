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

func StdInSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return strings.Split(stringInput, " ")
}

const Max int = 1000000

var T = make([]int, Max, Max)

func main() {
	n := StdInInt()
	var res = make([]string, 0)
	for i := 0; i < n; i++ {
		in := StdInSlice()
		cmd := in[0]
		str := in[1]
		if cmd == "insert" {
			Insert(T, GetKey(str))
		} else {
			if Search(T, GetKey(str)) {
				res = append(res, "yes")
			} else {
				res = append(res, "no")
			}
		}
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func GetKey(str string) int {
	buf := []rune("")
	for _, c := range str {
		switch string(c) {
		case "A":
			buf = append(buf, []rune("1")...)
		case "C":
			buf = append(buf, []rune("2")...)
		case "G":
			buf = append(buf, []rune("3")...)
		case "T":
			buf = append(buf, []rune("4")...)
		}
	}
	num, _ := strconv.Atoi(string(buf))
	return num
}

func H1(key int) int {
	return key % Max
}

func H2(key int) int {
	return 1 + (key % (Max - 1))
}

func H(key int, i int) int {
	return (H1(key) + i*H2(key)) % Max
}

func Insert(T []int, key int) {
	i := 0
	for i < Max {
		j := H(key, i)
		if T[j] == 0 {
			T[j] = key
			return
		} else {
			i++
		}
	}
}

func Search(T []int, key int) bool {
	i := 0
	for {
		j := H(key, i)
		if T[j] == key {
			return true
		} else if T[j] == 0 || i >= Max {
			return false
		} else {
			i++
		}
	}
}
