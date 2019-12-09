package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"log"
)

var opCodes = map[int]intcode.OperationFunc{
	1: intcode.Add,
	2: intcode.Multiply,
	3: intcode.Stop,
}

func partTwo(intCodes []int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			p := intcode.NewProgram(opCodes, intCodes, noun, verb)
			p.Execute()

			if p.GetValueAt(0) == 19690720 {
				log.Printf("Noun: %v Verb: %v", noun, verb)
				break
			}
		}
	}
}

func main() {
	rawIntCodes := files.Load("cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := intcode.NewProgram(opCodes, intCodes, 12, 2)
	prog.Execute()

	log.Println(prog.GetValueAt(0))

	partTwo(intCodes)
}
