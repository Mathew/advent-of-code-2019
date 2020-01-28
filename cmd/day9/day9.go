package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"log"
)

var opCodes = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.OutputAndDontPause},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
	9: {2, intcode.AdjustRelativeBase},
}

func main() {
	rawIntCodes := files.Load("cmd/day9/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)
	intCodesPadded := make([]int, 7157990)

	copy(intCodesPadded, intCodes)

	prog := intcode.NewProgramWithInputs(opCodes, intCodesPadded, []int{1})
	prog.Execute()
	log.Print(prog.GetResults())

	prog = intcode.NewProgramWithInputs(opCodes, intCodesPadded, []int{2, 2})
	prog.Execute()
	log.Print(prog.GetResults())
}
