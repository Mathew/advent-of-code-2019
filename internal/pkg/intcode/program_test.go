package intcode

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"testing"
)

var opCodes = map[int]OperationFunc{
	1: Add,
	2: Multiply,
	3: Stop,
}

func TestNewProgram(t *testing.T) {
	rawIntCodes := files.Load("../../../cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := NewProgram(opCodes, intCodes, 12, 2)
	prog.Execute()

	asserts.Equals(t, 5534943, prog.GetValueAt(0))
}

func TestNewProgramPart2(t *testing.T) {
	rawIntCodes := files.Load("../../../cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)
	found := false

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			p := NewProgram(opCodes, intCodes, noun, verb)
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
			running:  false,
		}
		p.Execute()
		asserts.Equals(t, pet.exp, p.intCodes)
	}
}
