package intcode

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"log"
)

const (
	STOPPED = iota
	RUNNING = iota
	PAUSED  = iota
)

type
(
	Program struct {
		intCodes            []int
		opCodes             map[int]OperationDesc
		pointer             int
		relativeBasePointer int
		state               int
		inputs              []int
		outputs             []int
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
	RELATIVE_MODE  = iota
)

func NewProgramWithNounAndVerb(opCodes map[int]OperationDesc, intCodes []int, noun int, verb int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	arr[1] = noun
	arr[2] = verb

	return Program{
		opCodes:             opCodes,
		intCodes:            arr,
		pointer:             0,
		relativeBasePointer: 0,
		state:               STOPPED,
	}
}

func NewProgramWithInputs(opCodes map[int]OperationDesc, intCodes []int, inputs []int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	return Program{
		opCodes:             opCodes,
		intCodes:            arr,
		pointer:             0,
		relativeBasePointer: 0,
		state:               STOPPED,
		inputs:              inputs,
	}
}

func NewProgram(opCodes map[int]OperationDesc, intCodes []int) Program {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	return Program{
		opCodes:             opCodes,
		intCodes:            arr,
		pointer:             0,
		relativeBasePointer: 0,
		state:               STOPPED,
	}
}

func (p *Program) setValue(pos, value int, mode int) {
	if mode == POSITION_MODE {
		p.intCodes[p.intCodes[pos]] = value
	} else if mode == RELATIVE_MODE {
		p.intCodes[p.relativeBasePointer+p.intCodes[pos]] = value
	} else {
		log.Fatalf("Unsupported write mode: %v", mode)
	}
}

func (p *Program) setPointer(pos int) {
	p.pointer = pos
}

func (p *Program) SetRelativePointer(i int) {
	p.relativeBasePointer = i
}

func (p Program) getIntCodeValue(pos int, mode int) int {
	if mode == POSITION_MODE {
		return p.intCodes[p.intCodes[pos]]
	} else if mode == IMMEDIATE_MODE {
		return p.intCodes[pos]
	} else if mode == RELATIVE_MODE {
		return p.intCodes[p.relativeBasePointer+p.intCodes[pos]]
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
	p.state = RUNNING

	for p.state == RUNNING {
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
	return p.outputs[len(p.outputs)-1]
}

func (p Program) GetResults() []int {
	return p.outputs
}

func (p *Program) GetInput() (int, bool) {
	if len(p.inputs) < 1 {
		return 0, false
	}
	v := p.inputs[0]
	p.inputs = p.inputs[1:]

	return v, true
}

func (p *Program) AddInput(i int) {
	p.inputs = append(p.inputs, i)
}

func (p Program) GetState() int {
	return p.state
}

func (p *Program) AddOutput(i int) {
	p.outputs = append(p.outputs, i)
}
