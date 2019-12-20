package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"github.com/mathew/advent-of-code-2019/internal/pkg/maths"
	"log"
)

var opCodes = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.Output},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
}

func AmplifierPhaseSettings(intCodes []int, permutations [][]int) (int, string) {
	highestSignal := 0
	var highestSetting string

	for _, p := range permutations {
		output := 0

		for _, phase := range p{
			prog := intcode.NewProgramWithInputs(opCodes, intCodes, []int{phase, output})
			prog.Execute()
			output = prog.GetResult()
		}

		if output > highestSignal {
			highestSignal = output
			highestSetting = converters.IntsToString("", p...)
		}
	}

	return highestSignal, highestSetting
}

func main() {
	instrs := files.Load("cmd/day7/input.txt", ",")
	intCodes := converters.StringsToInts(instrs...)
	inputs := []int{0, 1, 2, 3, 4}
	ps := maths.Permutations(inputs...)

	highest, phaseSettings := AmplifierPhaseSettings(intCodes, ps)
	log.Printf("Highest output: %v setting: %v", highest, phaseSettings)
}
