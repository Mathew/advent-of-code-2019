package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var wirePaths = []struct {
	first  []string
	second []string
	exp    int
}{
	{
		[]string{"R8", "U5", "L5", "D3"},
		[]string{"U7", "R6", "D4", "L4"},
		6,
	},
	{
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
		159,
	},
	{
		[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
		135,
	},
}

func TestFindDistance(t *testing.T) {
	for _, wP := range wirePaths {
		d := findDistanceToClosestIntersection(wP.first, wP.second)
		asserts.Equals(t, wP.exp, d)
	}
}

var vectorAddition = []struct {
	p   Point
	v   Vector
	exp Point
}{
	{
		NewPoint(1, 5),
		NewVector(6, 0),
		NewPoint(7, 5),
	},
	{
		NewPoint(-1, -5),
		NewVector(-6, 0),
		NewPoint(-7, -5),
	},
	{
		NewPoint(1, 5),
		NewVector(0, 6),
		NewPoint(1, 11),
	},
	{
		NewPoint(-1, -5),
		NewVector(0, -5),
		NewPoint(-1, -10),
	},
}

func TestPoint_AddVector(t *testing.T) {
	for _, va := range vectorAddition {
		asserts.Equals(t, va.exp, va.p.AddVector(va.v))
	}
}

var manhattanDistance = []struct {
	p   Point
	exp int
}{
	{
		NewPoint(1, 5),
		6,
	},
	{
		NewPoint(-1, -5),
		6,
	},
	{
		NewPoint(1, 5),
		6,
	},
	{
		NewPoint(-1, 5),
		6,
	},
}

func TestPoint_ManhattanDistanceFrom(t *testing.T) {
	for _, md := range manhattanDistance {
		asserts.Equals(t, md.exp, md.p.ManhattanDistance())
	}
}

var newVectorFromDirectionTests = []struct {
	direction string
	exp       Vector
}{
	{
		"R8",
		Vector{8, 0, 0},
	},
	{
		"U5",
		Vector{0, 5, 0},
	},
	{
		"L5",
		Vector{-5, 0, 0},
	},
	{
		"D3",
		Vector{0, -3, 0},
	},
}

func TestNewVectorFromDirection(t *testing.T) {
	for _, dt := range newVectorFromDirectionTests {
		asserts.Equals(t, dt.exp, NewVectorFromDirection(dt.direction))
	}

}

var intersectTests = []struct {
	name  int
	line1 Line
	line2 Line
	ok    bool
	point Point
}{
	{
		1,
		NewLine(NewPoint(1, 2), NewPoint(2, 1), NewVector(1, -1)),
		NewLine(NewPoint(2, 2), NewPoint(2, 2), NewVector(0, 0)),
		false,
		NewPoint(0, 0),
	},
	{
		2,
		NewLine(NewPoint(3, 1), NewPoint(3, 3), NewVector(0, 2)),
		NewLine(NewPoint(4, 1), NewPoint(4, 3), NewVector(0, 2)),
		false,
		NewPoint(0, 0),
	},
	{
		3,
		NewLine(NewPoint(0, 0), NewPoint(8, 0), NewVector(8, 0)),
		NewLine(NewPoint(6, 7), NewPoint(6, 3), NewVector(0, -4)),
		false,
		NewPoint(0, 0),
	},
	{
		4,
		NewLine(NewPoint(1, 0), NewPoint(8, 0), NewVector(7, 0)),
		NewLine(NewPoint(1, 7), NewPoint(1, 0), NewVector(0, -7)),
		true,
		NewPoint(1, 0),
	},
	{
		6,
		NewLine(NewPoint(-1, 0), NewPoint(-8, 0), NewVector(-7, 0)),
		NewLine(NewPoint(-1, -7), NewPoint(-1, 0), NewVector(0, 7)),
		true,
		NewPoint(-1, 0),
	},
	{
		5,
		NewLine(NewPoint(0, 0), NewPoint(8, 0), NewVector(8, 0)),
		NewLine(NewPoint(1, 7), NewPoint(1, -1), NewVector(0, -8)),
		true,
		NewPoint(1, 0),
	},
}

func TestLine_Intersects(t *testing.T) {
	for _, ni := range intersectTests {
		p, ok := ni.line1.Intersects(ni.line2)
		asserts.Assert(t, ni.ok == ok, "Name: %v, Exp %v got %v.  l1: %v l2: %v intersect p: %v \n", ni.name, ni.ok, ok, ni.line1, ni.line2, p)

		if ok {
			asserts.Assert(t, ni.point == p, "Name: %v, Exp %v got $v", ni.name, ni.point, p)
		}
	}
}
