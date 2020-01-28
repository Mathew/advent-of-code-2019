package intcode

import "log"

func Add(p *Program, modes []int) *Program {
	log.Print("Add")
	r := p.getIntCodeValue(p.pointer+1, modes[0]) + p.getIntCodeValue(p.pointer+2, modes[1])
	p.setValue(p.pointer+3, r, modes[2])
	p.setPointer(p.pointer + 4)

	return p
}

func Multiply(p *Program, modes []int) *Program {
	log.Print("Multiply")
	r := p.getIntCodeValue(p.pointer+1, modes[0]) * p.getIntCodeValue(p.pointer+2, modes[1])
	p.setValue(p.pointer+3, r, modes[2])
	p.setPointer(p.pointer + 4)

	return p
}

func Stop(p *Program, _ []int) *Program {
	log.Print("Stop")
	p.state = STOPPED

	return p
}

func SaveToPosition(p *Program, modes []int) *Program {
	log.Print("SaveToPosition")
	v, ok := p.GetInput()
	if !ok {
		log.Fatal("SaveToPosition: Input generator bust")
	}
	p.setValue(p.pointer+1, v, modes[0])
	p.setPointer(p.pointer + 2)

	return p
}

func Output(p *Program, modes []int) *Program {
	log.Print("Output")
	p.AddOutput(p.getIntCodeValue(p.pointer+1, modes[0]))
	p.setPointer(p.pointer + 2)

	return p
}

func OutputAndPause(p *Program, modes []int) *Program {
	log.Print("OutputAndPause")
	p.state = PAUSED
	p.AddOutput(p.getIntCodeValue(p.pointer+1, modes[0]))
	p.setPointer(p.pointer + 2)

	return p
}

func OutputAndDontPause(p *Program, modes []int) *Program {
	log.Print("OutputAndDontPause")
	p.AddOutput(p.getIntCodeValue(p.pointer+1, modes[0]))
	p.setPointer(p.pointer + 2)

	return p
}

func JumpIfTrue(p *Program, modes []int) *Program {
	log.Print("JumpIfTrue")
	if param := p.getIntCodeValue(p.pointer+1, modes[0]); param != 0 {
		p.setPointer(p.getIntCodeValue(p.pointer+2, modes[1]))
	} else {
		p.setPointer(p.pointer + 3)
	}

	return p
}

func JumpIfFalse(p *Program, modes []int) *Program {
	log.Print("JumpIfFalse")
	if param := p.getIntCodeValue(p.pointer+1, modes[0]); param == 0 {
		p.setPointer(p.getIntCodeValue(p.pointer+2, modes[1]))
	} else {
		p.setPointer(p.pointer + 3)
	}

	return p
}

func LessThan(p *Program, modes []int) *Program {
	log.Print("LessThan")
	if p.getIntCodeValue(p.pointer+1, modes[0]) < p.getIntCodeValue(p.pointer+2, modes[1]) {
		p.setValue(p.pointer+3, 1, modes[2])
	} else {
		p.setValue(p.pointer+3, 0, modes[2])
	}

	p.setPointer(p.pointer + 4)

	return p
}

func Equals(p *Program, modes []int) *Program {
	log.Print("Equals")
	if p.getIntCodeValue(p.pointer+1, modes[0]) == p.getIntCodeValue(p.pointer+2, modes[1]) {
		p.setValue(p.pointer+3, 1, modes[2])
	} else {
		p.setValue(p.pointer+3, 0, modes[2])
	}

	p.setPointer(p.pointer + 4)

	return p
}

func AdjustRelativeBase(p *Program, modes []int) *Program {
	log.Print("Adjust base")
	change := p.getIntCodeValue(p.pointer+1, modes[0])
	p.SetRelativePointer(p.relativeBasePointer + change)
	p.setPointer(p.pointer + 2)

	return p
}
