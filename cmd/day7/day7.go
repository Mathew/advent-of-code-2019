package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/intcode"
	"github.com/mathew/advent-of-code-2019/internal/pkg/maths"
	"log"
)

var opCodesPartOne = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.Output},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
}

var opCodesTwo = map[int]intcode.OperationDesc{
	1: {3, intcode.Add},
	2: {3, intcode.Multiply},
	3: {1, intcode.SaveToPosition},
	4: {1, intcode.OutputAndPause},
	5: {2, intcode.JumpIfTrue},
	6: {2, intcode.JumpIfFalse},
	7: {3, intcode.LessThan},
	8: {3, intcode.Equals},
}

func AmplifierPhaseSettings(opCodes map[int]intcode.OperationDesc, intCodes []int, permutations [][]int) (int, string) {
	highestSignal := 0
	var highestSetting string

	for _, p := range permutations {
		var programs []*intcode.Program
		output := 0

		for _, phase := range p {
			prog := intcode.NewProgramWithInputs(opCodes, intCodes, []int{phase, output})
			programs = append(programs, &prog)
			programs[len(programs)-1].Execute()
			output = programs[len(programs)-1].GetResult()
		}

		for programs[len(programs)-1].GetState() == intcode.PAUSED {
			for _, p := range programs {
				p.AddInput(output)
				p.Execute()
				output = p.GetResult()
			}
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

	highest, phaseSettings := AmplifierPhaseSettings(opCodesPartOne, intCodes, ps)
	log.Printf("Highest output: %v setting: %v", highest, phaseSettings)

	inputs = []int{5, 6, 7, 8, 9}
	ps = maths.Permutations(inputs...)
	highest, phaseSettings = AmplifierPhaseSettings(opCodesTwo, intCodes, ps)
	log.Printf("Part Two, Highest output: %v setting: %v", highest, phaseSettings)
}
