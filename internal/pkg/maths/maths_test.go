package maths

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var absTests = []struct {
	input int
	exp   int
}{
	{1, 1},
	{0, 0},
	{-5, 5},
}

func TestAbs(t *testing.T) {
	for _, to := range absTests {
		asserts.Equals(t, to.exp, Abs(to.input))
	}
}

var minOfTests = []struct {
	exp   int
	input []int
}{
	{1, []int{1, 2, 3}},
	{-3, []int{-1, -2, -3}},
	{0, []int{0, 2, 3}},
}

func TestMinOf(t *testing.T) {
	for _, to := range minOfTests {
		asserts.Equals(t, to.exp, MinOf(to.input...))
	}
}

var zeroOrOneTests = []struct {
	input int
	exp   int
}{
	{1, 1},
	{2, 1},
	{-1, -1},
	{-2, -1},
	{0, 0},
}

func TestZeroOrOne(t *testing.T) {
	for _, to := range zeroOrOneTests {
		asserts.Equals(t, to.exp, ZeroOrOne(to.input))
	}
}

func TestPermutations(t *testing.T) {
	exp := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 2, 1},
		{2, 3, 1},
		{3, 1, 2},
		{1, 3, 2},
	}

	asserts.Equals(t, exp, Permutations(1, 2, 3))
}
