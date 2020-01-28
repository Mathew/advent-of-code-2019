package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"strconv"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{
		[]int{104, 1125899906842624, 99},
		1125899906842624,
	},
}

var opCodesTest = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.OutputAndPause},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
	9: {2, intcode.AdjustRelativeBase},
}

func TestProgram(t *testing.T) {
	for _, tc := range tests {
		p := intcode.NewProgram(opCodesTest, tc.input)
		p.Execute()
		asserts.Equals(t, tc.output, p.GetResult())
	}
}

func TestProgramOutputLength(t *testing.T) {
	p := intcode.NewProgram(opCodesTest, []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0})
	p.Execute()
	asserts.Equals(t, 16, len(strconv.Itoa(p.GetResult())))
}

func TestProgramPause(t *testing.T) {
	input := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	output := []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	intCodes := make([]int, 200)
	copy(intCodes, input)

	// Use don't pause operation.
	p := intcode.NewProgram(opCodes, intCodes)
	p.Execute()

	asserts.Equals(t, output, p.GetResults())
}
