package intcode

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"testing"
)

var opCodes = map[int]OperationDesc{
	1: {2, Add},
	2: {2, Multiply},
	3: {0, Stop},
}

func TestNewProgram(t *testing.T) {
	rawIntCodes := files.Load("../../../cmd/day2/inputs.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := NewProgramWithNounAndVerb(opCodes, intCodes, 12, 2)
	prog.Execute()

	asserts.Equals(t, 5534943, prog.GetValueAt(0))
}

func TestNewProgramPart2(t *testing.T) {
	rawIntCodes := files.Load("../../../cmd/day2/inputs.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)
	found := false

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			p := NewProgramWithNounAndVerb(opCodes, intCodes, noun, verb)
			p.Execute()

			if p.GetValueAt(0) == 19690720 {
				found = true
				asserts.Equals(t, 76, noun)
				asserts.Equals(t, 3, verb)
				break
			}
		}
	}

	asserts.Equals(t, true, found)
}

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
		p := Program{
			opCodes:  opCodes,
			intCodes: pet.initial,
			pointer:  0,
			state:    false,
		}
		p.Execute()
		asserts.Equals(t, pet.exp, p.intCodes)
	}
}

func TestProgram_getOpCodeValue(t *testing.T) {
	p := NewProgram(opCodes, []int{2, 3, 0, 1})
	asserts.Equals(t, 2, p.getOpCodeValue(2))

	p = NewProgram(opCodes, []int{102, 3, 0, 1})
	asserts.Equals(t, 2, p.getOpCodeValue(102))
}

func TestProgram_getOpCodeModes(t *testing.T) {
	p := NewProgram(opCodes, []int{2, 3, 0, 1})
	asserts.Equals(t, []int{POSITION_MODE, POSITION_MODE}, p.getOpCodeModes(2, 2))

	p = NewProgram(opCodes, []int{102, 3, 0, 1})
	asserts.Equals(t, []int{IMMEDIATE_MODE, POSITION_MODE}, p.getOpCodeModes(102, 2))
}
