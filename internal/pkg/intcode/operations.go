package intcode

func Add(p *Program) *Program {
	r := p.getIntCodeValue(p.pointer+1) + p.getIntCodeValue(p.pointer+2)
	p.setValue(p.pointer+3, r)
	p.setPointer(p.pointer + 4)

	return p
}

func Multiply(p *Program) *Program {
	r := p.getIntCodeValue(p.pointer+1) * p.getIntCodeValue(p.pointer+2)
	p.setValue(p.pointer+3, r)
	p.setPointer(p.pointer + 4)

	return p
}

func Stop(p *Program) *Program {
	p.running = false

	return p
}
