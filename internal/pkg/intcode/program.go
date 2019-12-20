package intcode

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/generator"
	"log"
)

type
(
	Program struct {
		intCodes []int
		opCodes  map[int]OperationDesc
		pointer  int
		running  bool
		inputs   generator.Generator
		output   int
	}
	OperationFunc func(program *Program, modes []int) *Program
	OperationDesc struct {
		NumArguments int
		Operation    OperationFunc
	}
)

const (
	POSITION_MODE  = iota
	IMMEDIATE_MODE = iota
)

func NewProgramWithNounAndVerb(opCodes map[int]OperationDesc, intCodes []int, noun int, verb int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	arr[1] = noun
	arr[2] = verb

	return Program{
		opCodes:  opCodes,
		intCodes: arr,
		pointer:  0,
		running:  false,
	}
}

func NewProgramWithInputs(opCodes map[int]OperationDesc, intCodes []int, inputs []int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	return Program{
		opCodes:  opCodes,
		intCodes: arr,
		pointer:  0,
		running:  false,
		inputs:   generator.CreateIntGenerator(inputs...),
	}
}

func NewProgram(opCodes map[int]OperationDesc, intCodes []int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	return Program{
		opCodes:  opCodes,
		intCodes: arr,
		pointer:  0,
		running:  false,
	}
}

func (p *Program) setValue(pos, value int) {
	p.intCodes[p.intCodes[pos]] = value
}

func (p *Program) setPointer(pos int) {
	p.pointer = pos
}

func (p Program) getIntCodeValue(pos int, mode int) int {
	if mode == POSITION_MODE {
		return p.intCodes[p.intCodes[pos]]
	} else if mode == IMMEDIATE_MODE {
		return p.intCodes[pos]
	} else {
		log.Fatalf("Unrecognised mode: %v", mode)
	}

	return 0
}

func (p Program) GetValueAt(pos int) int {
	return p.intCodes[pos]
}

func (p Program) getOpCodeValue(opCode int) int {
	if opCode > 9 {
		digits := converters.IntToDigits(opCode)

		return digits[len(digits)-1]
	}

	return opCode
}

func (p Program) getOpCodeModes(opCode, numParams int) []int {
	var modes []int
	if opCode > 9 {
		digits := converters.IntToDigits(opCode)
		modes = converters.Reverse(digits[:len(digits)-2])
	}

	if l := numParams - len(modes); l > 0 {
		for x := 0; x < l; x++ {
			modes = append(modes, POSITION_MODE)
		}
	}

	return modes
}

func (p *Program) Execute() {
	p.running = true

	for p.running {
		opCode := p.intCodes[p.pointer]
		code := p.getOpCodeValue(opCode)

		if op, ok := p.opCodes[code]; ok {
			modes := p.getOpCodeModes(opCode, op.NumArguments)
			op.Operation(p, modes)
		} else {
			Stop(p, []int{})
		}
	}
}

func (p Program) GetIntCodes() []int {
	return p.intCodes
}

func (p Program) GetResult() int {
	return p.output
}

func (p Program) GetInput() (int, bool) {
	return p.inputs()
}
