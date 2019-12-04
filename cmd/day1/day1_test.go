package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var calculateFuelTests = []struct {
	mass   int
	result int
}{
	{
		12,
		2,
	},
	{
		14,
		2,
	},
	{
		1969,
		654,
	},
	{
		100756,
		33583,
	},
}

func TestCalculateFuel(t *testing.T) {
	for _, ft := range calculateFuelTests {
		asserts.Equals(t, ft.result, calculateFuel(ft.mass))
	}
}

func TestSumFuel(t *testing.T) {
	asserts.Equals(t, 6, sum(1, 2, 3))
}
