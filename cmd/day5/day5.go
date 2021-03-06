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
	4: {1, intcode.Output},
}

var moreOpCodes = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.Output},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
}

func main() {
	rawIntCodes := files.Load("cmd/day5/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := intcode.NewProgramWithInputs(opCodes, intCodes, []int{1})
	prog.Execute()
	log.Print(prog.GetResult())

	prog = intcode.NewProgramWithInputs(moreOpCodes, intCodes, []int{5})
	prog.Execute()
	log.Print(prog.GetResult())
}
