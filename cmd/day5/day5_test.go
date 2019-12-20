package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"testing"
)

var programExecutionTests = []struct {
	initial []int
	exp     []int
}{
	{
		[]int{1002, 4, 3, 4, 33},
		[]int{1002, 4, 3, 4, 99,},
	},
	{
		[]int{1101, 100, -1, 4, 0},
		[]int{1101, 100, -1, 4, 99},
	},
}

func TestProgramExecution(t *testing.T) {
	for _, pet := range programExecutionTests {
		p := intcode.NewProgram(opCodes, pet.initial)
		p.Execute()

		asserts.Equals(t, pet.exp, p.GetIntCodes())
	}
}
