package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Key  int
	Prev *Node
	Next *Node
}

var Null = Node{0, nil, nil}

func StdInSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	return strings.Split(stringInput, " ")
}

func main() {
	Initialize()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		in := StdInSlice()
		ExecCmd(in)
	}
	cur := Null.Next
	buf := []byte("")
	for cur != &Null {
		buf = append(buf, []byte(strconv.Itoa(cur.Key))...)
		buf = append(buf, []byte(" ")...)
		cur = cur.Next
	}
	fmt.Printf("%v \n", string(buf))
}

func ExecCmd(in []string) {
	cmd := in[0]
	switch cmd {
	case "insert":
		key, _ := strconv.Atoi(in[1])
		Insert(key)
	case "delete":
		key, _ := strconv.Atoi(in[1])
		DeleteKey(key)
	case "deleteFirst":
		DeleteFirst()
	case "deleteLast":
		DeleteLast()
	}
}

func Initialize() {
	Null.Next = &Null
	Null.Prev = &Null
}

func Insert(key int) {
	x := &Node{key, &Null, Null.Next}
	Null.Next.Prev = x
	Null.Next = x
	x.Prev = &Null
}

func DeleteNode(node *Node) {
	if node == &Null {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func DeleteFirst() {
	DeleteNode(Null.Next)
}

func DeleteLast() {
	DeleteNode(Null.Prev)
}

func DeleteKey(key int) {
	DeleteNode(ListSearch(key))
}

func ListSearch(key int) *Node {
	cur := Null.Next
	for cur != &Null && cur.Key != key {
		cur = cur.Next
	}
	return cur
}
