package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"log"
)

type IntCodeProgram struct {
	intCodes []int
	pointer  int
	running  bool
}

func NewIntCodeProgram(intCodes []int, noun int, verb int) IntCodeProgram {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	arr[1] = noun
	arr[2] = verb

	return IntCodeProgram{
		intCodes: arr,
		pointer:  0,
		running:  false,
	}
}

func (p IntCodeProgram) getIntCodeValue(pos int) int {
	return p.intCodes[p.intCodes[pos]]
}

func (p *IntCodeProgram) add() {
	r := p.getIntCodeValue(p.pointer+1) + p.getIntCodeValue(p.pointer+2)
	p.intCodes[p.intCodes[p.pointer+3]] = r

	p.pointer += 4
}

func (p *IntCodeProgram) multiply() {
	r := p.getIntCodeValue(p.pointer+1) * p.getIntCodeValue(p.pointer+2)
	p.intCodes[p.intCodes[p.pointer+3]] = r

	p.pointer += 4
}

func (p *IntCodeProgram) stop() {
	p.running = false
}

func (p *IntCodeProgram) Execute() {
	p.running = true

	for ok := true; ok; ok = p.running {
		op := p.intCodes[p.pointer]

		switch op {
		case 1:
			p.add()
		case 2:
			p.multiply()
		case 3:
			p.stop()
		default:
			p.stop()
		}
	}
}

func partTwo(intCodes []int) {
	for noun :=0; noun < 100; noun++ {
		for verb :=0; verb < 100; verb++ {
			p := NewIntCodeProgram(intCodes, noun, verb)
			p.Execute()

			if p.intCodes[0] == 19690720 {
				log.Printf("Noun: %v Verb: %v", noun, verb)
				break
			}
		}
	}
}

func main() {
	rawIntCodes := files.Load("cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	prog := NewIntCodeProgram(intCodes, 12, 2)
	prog.Execute()

	log.Println(prog.intCodes[0])

	partTwo(intCodes)
}
