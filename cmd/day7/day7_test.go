package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/maths"
	"testing"
)

var tests = []struct {
	input     []int
	expSignal int
	expPhase  string
}{
	{
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		43210, "43210",
	},
	{
		[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		54321, "01234",
	},
	{
		[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		65210, "10432",
	},
}

func TestAmplifierPhaseSettings(t *testing.T) {
	for _, tc := range tests {
		signal, phaseSetting := AmplifierPhaseSettings(tc.input, maths.Permutations(0, 1, 2, 3, 4))
		asserts.Equals(t, tc.expSignal, signal)
		asserts.Equals(t, tc.expPhase, phaseSetting)
	}
}
