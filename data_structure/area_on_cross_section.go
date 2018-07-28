package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func StdInString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

const Max int = 1000

type Stack struct {
	Stack []int
	Top   int
}

type Area struct {
	Start int
	Area  int
}

type AreaStack struct {
	Stack []Area
	Top   int
}

func (s *Stack) IsEmpty() bool {
	return s.Top == 0
}

func (s *Stack) IsFull() bool {
	return s.Top == Max
}

func (s *Stack) Push(x int) (int, error) {
	if s.IsFull() {
		return 0, errors.New("Stack overflow")
	}
	s.Top++
	s.Stack[s.Top] = x
	return x, nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack empty")
	}
	s.Top--
	return s.Stack[s.Top+1], nil
}

func (s *AreaStack) IsEmpty() bool {
	return s.Top == 0
}

func (s *AreaStack) IsFull() bool {
	return s.Top == Max
}

func (s *AreaStack) Push(x Area) (Area, error) {
	if s.IsFull() {
		return Area{}, errors.New("Stack overflow")
	}
	s.Top++
	s.Stack[s.Top] = x
	return x, nil
}

func (s *AreaStack) Pop() (Area, error) {
	if s.IsEmpty() {
		return Area{}, errors.New("Stack empty")
	}
	s.Top--
	return s.Stack[s.Top+1], nil
}

func main() {
	in := StdInString()
	s1 := Stack{make([]int, Max), 0}
	s2 := AreaStack{make([]Area, Max), 0}
	for i, v := range in {
		if v == '\\' {
			s1.Push(i)
		}
		if v == '/' {
			s, e := s1.Pop()
			if e == nil {
				t := i - s
				for {
					sa, e := s2.Pop()
					if e != nil {
						break
					}
					if sa.Start > s {
						t += sa.Area
					} else {
						s2.Push(sa)
						break
					}
				}
				s2.Push(Area{s, t})
			}
		}
	}
	total := 0
	ans := make([]int, 0)

	for {
		sa, e := s2.Pop()
		if e != nil {
			break
		}
		total += sa.Area
		ans = append([]int{sa.Area}, ans...)
	}
	fmt.Println(total)
	fmt.Printf("%d ", len(ans))
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
}
