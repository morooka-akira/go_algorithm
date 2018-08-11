package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func StdInInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	n, _ := strconv.Atoi(input)
	return n
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) PrintPoint() {
	fmt.Printf("%.8f %.8f\n", p.X, p.Y)
}

// dはdepth
func Koch(d int, p1 *Point, p2 *Point) {
	if d == 0 {
		return
	}
	s := Point{float64(0.0), float64(0.0)}
	s.X = (2*p1.X + 1*p2.X) / 3
	s.Y = (2*p1.Y + 1*p2.Y) / 3
	t := Point{float64(0.0), float64(0.0)}
	t.X = (1*p1.X + 2*p2.X) / 3
	t.Y = (1*p1.Y + 2*p2.Y) / 3
	u := Point{}
	u.X = (t.X-s.X)*math.Cos(float64(60.0*math.Pi/180)) - (t.Y-s.Y)*math.Sin(float64(60.0*math.Pi/180)) + s.X
	u.Y = (t.X-s.X)*math.Sin(float64(60.0*math.Pi/180)) + (t.Y-s.Y)*math.Cos(float64(60.0*math.Pi/180)) + s.Y

	Koch(d-1, p1, &s)
	s.PrintPoint()
	Koch(d-1, &s, &u)
	u.PrintPoint()
	Koch(d-1, &u, &t)
	t.PrintPoint()
	Koch(d-1, &t, p2)
}

func main() {
	d := StdInInt()
	p1 := &Point{float64(0.0), float64(0.0)}
	p2 := &Point{float64(100.0), float64(0.0)}
	p1.PrintPoint()
	Koch(d, p1, p2)
	p2.PrintPoint()
}
