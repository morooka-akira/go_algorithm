package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	p int
	l int
	r int
}

var NIL int = -1

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

func StdInInt() int {
	var N int
	fmt.Scan(&N)
	return N
}

func StdInNumSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	A, _ := SliceAtoi(strings.Split(stringInput, " "))
	return A
}

func SetDepth(n int, d int, depth []int, tree []Node) {
	depth[n] = d
	if tree[n].r != NIL {
		SetDepth(tree[n].r, d, depth, tree)
	} else if tree[n].l != NIL {
		SetDepth(tree[n].l, d+1, depth, tree)
	}
}

func DisplayChildren(n int, tree []Node) {
	c := tree[n].l
	a := []string{}
	for c != NIL {
		a = append(a, strconv.Itoa(c))
		c = tree[c].r
	}
	fmt.Print(strings.Join(a, ", "))
}

func Display(N int, D []int, tree []Node) {
	for i := 0; i < N; i++ {
		fmt.Printf("node %d: ", i)
		fmt.Printf("parent = %d, ", tree[i].p)
		fmt.Printf("depth = %d, ", D[i])
		if tree[i].p == NIL {
			fmt.Print("root, ")
		} else if tree[i].l == NIL {
			fmt.Print("leaf, ")
		} else {
			fmt.Print("internal node, ")
		}
		fmt.Print("[")
		DisplayChildren(i, tree)
		fmt.Print("]")
		fmt.Print("\n")
	}
}

func main() {
	N := StdInInt()
	tree := make([]Node, N, N)
	D := make([]int, N, N)
	for i := 0; i < N; i++ {
		tree[i].p = NIL
		tree[i].l = NIL
		tree[i].r = NIL
	}
	for i := 0; i < N; i++ {
		node := StdInNumSlice()
		id := node[0]
		cnt := node[1]
		// 子の処理
		l := 0
		for j := 0; j < cnt; j++ {
			cid := node[2+j]
			if j == 0 { // 一番左の子をセット
				tree[id].l = cid
			} else {
				tree[l].r = cid
			}
			l = cid
			// 親をセット
			tree[cid].p = id
		}
	}
	SetDepth(0, 0, D, tree)
	Display(N, D, tree)
}
