package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var programExecutionTests = []struct {
	initial []int
	exp     []int
}{
	{
		[]int{1, 0, 0, 0, 99},
		[]int{2, 0, 0, 0, 99},
	},
	{
		[]int{2, 3, 0, 3, 99},
		[]int{2, 3, 0, 6, 99},
	},
	{
		[]int{2, 4, 4, 5, 99, 0},
		[]int{2, 4, 4, 5, 99, 9801},
	},
	{
		[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
}

func TestProgramExecution(t *testing.T) {
	for _, pet := range programExecutionTests {
		p := IntCodeProgram{
			intCodes: pet.initial,
			pointer:  0,
			running:  false,
		}
		p.Execute()
		asserts.Equals(t, pet.exp, p.intCodes)
	}
}
