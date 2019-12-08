package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/maths"
	"log"
	"math"
	"strings"
)

type Vector struct {
	x int
	y int
	w int
}

func NewVector(x, y int) Vector {
	return Vector{x, y, 0}
}

func NewVectorFromDirection(direction string) Vector {
	x := 0
	y := 0

	dir := string(direction[0])

	switch dir {
	case "R":
		x = converters.StringToInt(direction[1:])
	case "L":
		x = -converters.StringToInt(direction[1:])
	case "U":
		y = converters.StringToInt(direction[1:])
	case "D":
		y = -converters.StringToInt(direction[1:])
	}

	return Vector{
		x, y, 0,
	}
}

type Point struct {
	x int
	y int
	w int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
		w: 1,
	}
}

func (p Point) AddVector(v Vector) Point {
	return NewPoint(p.x+v.x, p.y+v.y)
}

func (p Point) ManhattanDistance() int {
	d := maths.Abs(p.x) + maths.Abs(p.y)
	if d < 0 {
		return -d
	}
	return d
}

type Line struct {
	start  Point
	end    Point
	vector Vector
}

func NewLine(start, end Point, vector Vector) Line {
	return Line{start, end, vector}
}

func (l Line) Intersects(o Line) (Point, bool) {
	// l1
	x1 := l.start.x
	y1 := l.start.y
	x2 := l.end.x
	y2 := l.end.y

	//l2
	x3 := o.start.x
	y3 := o.start.y
	x4 := o.end.x
	y4 := o.end.y

	det := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if det == 0 {
		return Point{}, false
	}

	circuit := map[Point]int{}

	for _, cl := range []Line{l, o} {
		v := NewVector(maths.ZeroOrOne(cl.vector.x), maths.ZeroOrOne(cl.vector.y))
		p := cl.start
		done := false
		for !done {
			circuit[p] += 1

			if circuit[p] == 2 {
				return p, true
			}

			if p == cl.end {
				done = true
			}

			p = p.AddVector(v)
		}
	}

	return Point{}, false
}

type Wire struct {
	lines []Line
	ptr   Point
}

func NewWire() Wire {
	return Wire{
		lines: []Line{},
		ptr:   Point{},
	}
}

func (w *Wire) AddVector(v Vector) {
	end := w.ptr.AddVector(v)
	w.lines = append(w.lines, NewLine(w.ptr, end, v))
	w.ptr = end
}

func (w Wire) Intersections(o Wire) []Point {
	var is []Point
	ic := 0
	for _, l := range w.lines {
		for _, l2 := range o.lines {
			if p, ok := l.Intersects(l2); ok {
				is = append(is, p)
				ic += 1
			}
		}
	}

	return is
}

func findDistanceToClosestIntersection(wireAVectors, wireBVectors []string) int {
	wireA := NewWire()
	wireB := NewWire()
	for _, vector := range wireAVectors {
		wireA.AddVector(NewVectorFromDirection(vector))
	}
	for _, vector := range wireBVectors {
		wireB.AddVector(NewVectorFromDirection(vector))
	}

	intersections := wireA.Intersections(wireB)

	smallest := math.MaxInt64
	for _, i := range intersections {
		d := i.ManhattanDistance()
		if d < smallest && d != 0 {
			smallest = d
		}
	}

	return smallest
}

func main() {
	rawWireVectors := files.Load("cmd/day3/input.txt", "\n")
	WireAVectors := strings.Split(rawWireVectors[0], ",")
	WireBVectors := strings.Split(rawWireVectors[1], ",")

	smallest := findDistanceToClosestIntersection(WireAVectors, WireBVectors)

	log.Println(smallest)
}
