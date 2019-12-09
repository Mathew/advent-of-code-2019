package day5

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"log"
)

func SaveToPosition(program *intcode.Program) *intcode.Program {
	return program
}

func Output(program *intcode.Program) *intcode.Program {
	return program
}

var opCodes = map[int]intcode.OperationFunc{
	3: SaveToPosition,
	4: Output,
}

func main() {
	rawIntCodes := files.Load("cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := intcode.NewProgram(opCodes, intCodes, 12, 2)
	prog.Execute()

	log.Println(prog.GetValueAt(0))
}
