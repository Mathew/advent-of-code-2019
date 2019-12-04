package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"log"
)

type IntCodeProgram struct {
	intCodes []int
	position int
	running  bool
}

func NewIntCodeProgram(intCodes []int) IntCodeProgram {
	arr := make([]int, len(intCodes))
	copy(arr, intCodes)

	return IntCodeProgram{
		intCodes: arr,
		position: 0,
		running:  false,
	}
}

func (p IntCodeProgram) getIntCodeValue(pos int) int {
	return p.intCodes[p.intCodes[pos]]
}

func (p *IntCodeProgram) add() {
	r := p.getIntCodeValue(p.position + 1) + p.getIntCodeValue(p.position + 2)
	p.intCodes[p.intCodes[p.position + 3]] = r

	p.position += 4
}

func (p *IntCodeProgram) multiply() {
	r := p.getIntCodeValue(p.position + 1) * p.getIntCodeValue(p.position + 2)
	p.intCodes[p.intCodes[p.position + 3]] = r

	p.position += 4
}

func (p *IntCodeProgram) stop() {
	p.running = false
}

func (p *IntCodeProgram) Execute() {
	p.running = true

	for ok := true; ok; ok = p.running {
		op := p.intCodes[p.position]

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

func main() {
	rawIntCodes := files.Load("cmd/day2/input.txt", ",")
	intCodes := converters.StringsToInts(rawIntCodes...)

	intCodes[1] = 12
	intCodes[2] = 2

	prog := NewIntCodeProgram(intCodes)
	prog.Execute()

	log.Println(prog.intCodes[0])
}
