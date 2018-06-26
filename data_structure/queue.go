package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	name string
	time int
}

const Max int = 1000000

var queue = make([]*Process, Max, Max)
var head int = 0
var tail int = 0

func StdInSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return strings.Split(stringInput, " ")
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	Initialize()
	info := StdInSlice()
	n, _ := strconv.Atoi(info[0])
	q, _ := strconv.Atoi(info[1])
	for i := 0; i < n; i++ {
		p := StdInSlice()
		name := p[0]
		time, _ := strconv.Atoi(p[1])
		_, e := Enqueue(&Process{name, time})
		if e != nil {
			panic("Fail Enqueue.")
		}
	}
	time := 0
	for false == IsEmpty() {
		p, e := Dequeue()
		if e != nil {
			panic("Fail Dequeue.")
		}
		c := Min(q, (*p).time)
		(*p).time -= c
		time += c
		if (*p).time > 0 {
			Enqueue(p)
		} else {
			fmt.Printf("%v %v\n", (*p).name, time)
		}
	}
}

func Initialize() {
	head = 0
	tail = 0
}

func IsEmpty() bool {
	return head == tail
}

func IsFull() bool {
	return head == (tail+1)%Max
}

func Enqueue(x *Process) (*Process, error) {
	if IsFull() {
		return nil, errors.New("Overflow")
	}
	queue[tail] = x
	tail = (tail + 1) % Max
	return x, nil
}

func Dequeue() (*Process, error) {
	if IsEmpty() {
		return nil, errors.New("Underflow")
	}
	v := queue[head]
	head = (head + 1) % Max
	return v, nil
}
