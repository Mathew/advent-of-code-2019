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

func (w Wire) WalkToCount(p Point) (int, bool) {
	steps := 0

	for x, l := range w.lines {

		// account for lines having the same end and start.
		if x > 0 {
			steps -= 1
		}

		cp := l.start
		v := NewVector(maths.ZeroOrOne(l.vector.x), maths.ZeroOrOne(l.vector.y))
		end := false

		for end == false {
			if cp == p {
				return steps, true
			}

			if cp == l.end {
				end = true
			}

			steps += 1
			cp = cp.AddVector(v)
		}
	}

	return 0, false
}

func createWires(rawA, rawB []string) (Wire, Wire) {
	wireA := NewWire()
	wireB := NewWire()
	for _, vector := range rawA {
		wireA.AddVector(NewVectorFromDirection(vector))
	}
	for _, vector := range rawB {
		wireB.AddVector(NewVectorFromDirection(vector))
	}

	return wireA, wireB
}

func getClosestIntersection(intersections []Point) int {
	smallest := math.MaxInt64
	for _, i := range intersections {
		d := i.ManhattanDistance()
		if d < smallest && d != 0 {
			smallest = d
		}
	}

	return smallest
}

func GetSmallestStepIntersection(wireA, wireB Wire, intersections []Point) int {
	smallest := math.MaxInt64

	for _, p := range intersections {

		l1, ok := wireA.WalkToCount(p)
		if !ok {
			log.Fatal("WireA doesn't intersect")
		}

		l2, ok := wireB.WalkToCount(p)
		if !ok {
			log.Fatal("WireB doesn't intersect")
		}

		l := l1 + l2
		if l < smallest && l > 0 {
			smallest = l
		}
	}

	return smallest
}

func main() {
	rawWireVectors := files.Load("cmd/day3/input.txt", "\n")
	rawWireA := strings.Split(rawWireVectors[0], ",")
	rawWireB := strings.Split(rawWireVectors[1], ",")

	wireA, wireB := createWires(rawWireA, rawWireB)
	intersections := wireA.Intersections(wireB)

	log.Printf("Closest intersection: %v \n", getClosestIntersection(intersections))
	log.Printf("Lowest step intersection: %v \n", GetSmallestStepIntersection(wireA, wireB, intersections))
}
