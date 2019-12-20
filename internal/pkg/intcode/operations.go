package intcode

import "log"

func Add(p *Program, modes []int) *Program {
	r := p.getIntCodeValue(p.pointer+1, modes[0]) + p.getIntCodeValue(p.pointer+2, modes[1])
	p.setValue(p.pointer+3, r)
	p.setPointer(p.pointer + 4)

	return p
}

func Multiply(p *Program, modes []int) *Program {
	r := p.getIntCodeValue(p.pointer+1, modes[0]) * p.getIntCodeValue(p.pointer+2, modes[1])
	p.setValue(p.pointer+3, r)
	p.setPointer(p.pointer + 4)

	return p
}

func Stop(p *Program, _ []int) *Program {
	p.running = false

	return p
}

func SaveToPosition(p *Program, modes []int) *Program {
	v, ok := p.GetInput()
	if !ok {
		log.Fatal("SaveToPosition: Input generator bust")
	}
	p.setValue(p.pointer+1, v)
	p.setPointer(p.pointer + 2)

	return p
}

func Output(p *Program, modes []int) *Program {
	p.output = p.getIntCodeValue(p.pointer+1, modes[0])
	p.setPointer(p.pointer + 2)

	return p
}

func JumpIfTrue(p *Program, modes []int) *Program {
	if param := p.getIntCodeValue(p.pointer+1, modes[0]); param != 0 {
		p.setPointer(p.getIntCodeValue(p.pointer+2, modes[1]))
	} else {
		p.setPointer(p.pointer + 3)
	}

	return p
}

func JumpIfFalse(p *Program, modes []int) *Program {
	if param := p.getIntCodeValue(p.pointer+1, modes[0]); param == 0 {
		p.setPointer(p.getIntCodeValue(p.pointer+2, modes[1]))
	} else {
		p.setPointer(p.pointer + 3)
	}

	return p
}

func LessThan(p *Program, modes []int) *Program {
	if p.getIntCodeValue(p.pointer+1, modes[0]) < p.getIntCodeValue(p.pointer+2, modes[1]) {
		p.setValue(p.pointer+3, 1)
	} else {
		p.setValue(p.pointer+3, 0)
	}

	p.setPointer(p.pointer + 4)

	return p
}

func Equals(p *Program, modes []int) *Program {
	if p.getIntCodeValue(p.pointer+1, modes[0]) == p.getIntCodeValue(p.pointer+2, modes[1]) {
		p.setValue(p.pointer+3, 1)
	} else {
		p.setValue(p.pointer+3, 0)
	}

	p.setPointer(p.pointer + 4)

	return p
}
