package intcode

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var testOpCodes = map[int]OperationDesc{
	1: {3, Add},
	2: {3, Multiply},
	3: {1, SaveToPosition},
	4: {1, OutputAndPause},
	5: {2, JumpIfTrue},
	6: {2, JumpIfFalse},
	7: {3, LessThan},
	8: {3, Equals},
	9: {2, AdjustRelativeBase},
}

func TestAdjustRelativeBase(t *testing.T) {
	intCodes := make([]int, 20, 20)
	copy(intCodes, []int{4, 109, 19, 204, -20})

	prog := NewProgramWithInputs(testOpCodes, intCodes, []int{1})
	prog.SetRelativePointer(1)
	prog.setPointer(1)
	prog.Execute()

	asserts.Equals(t, 20, prog.relativeBasePointer)
	asserts.Equals(t, 4, prog.GetResult())
}

var ts = []struct {
	name    string
	initial []int
	input   int
	exp     int
}{
	{"1", []int{109, -1, 4, 1, 99}, 0, -1},
	{"2", []int{109, -1, 104, 1, 99}, 0, 1},
	{"3", []int{109, -1, 204, 1, 99}, 0, 109},
	{"4", []int{109, 1, 9, 2, 204, -6, 99}, 0, 204},
	{"5", []int{109, 1, 109, 9, 204, -6, 99}, 0, 204},
	{"6", []int{109, 1, 209, -1, 204, -106, 99}, 0, 204},
	{"7", []int{109, 1, 3, 3, 204, 2, 99}, 100, 100},
	{"8", []int{109, 1, 203, 2, 204, 2, 99}, 100, 100},
}

func TestOpCodes(t *testing.T) {
	for _, ct := range ts {
		p := NewProgramWithInputs(testOpCodes, ct.initial, []int{ct.input})
		p.Execute()
		asserts.Equals(t, ct.exp, p.GetResult())
	}
}
