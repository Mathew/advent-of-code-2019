package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
)

var opCodes = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.Output},
}

func main() {
	rawIntCodes := files.Load("cmd/day5/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := intcode.NewProgram(opCodes, intCodes)
	prog.Execute()
}
