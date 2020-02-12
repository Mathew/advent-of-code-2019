package vectors

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/maths"
	"math"
)

type Vector struct {
	X int
	Y int
	w int
}

func NewVector(x, y int) Vector {
	return Vector{x, y, 0}
}

func (v Vector) Slope() float64 {
	if v.X == 0 {
		return math.MaxInt64
	}

	return float64(v.Y) / float64(v.X)
	//return math.Round((float64(v.Y) / float64(v.X))*100)/100
}

func (v Vector) ManhattanDistance() int {
	d := maths.Abs(v.X) + maths.Abs(v.Y)
	if d < 0 {
		return -d
	}
	return d
}

type Point struct {
	X int
	Y int
	w int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
		w: 1,
	}
}

func (p Point) AddVector(v Vector) Point {
	return NewPoint(p.X+v.X, p.Y+v.Y)
}

func (p Point) Subtract(p2 Point) Vector{
	return NewVector(p.X- p2.X, p.Y- p2.Y)
}

func (p Point) Slope(p2 Point) int {
	return (p2.Y - p.Y) / (p2.X - p.X)
}
