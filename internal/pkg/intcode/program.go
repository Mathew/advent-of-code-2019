package intcode

type Program struct {
	intCodes []int
	opCodes  map[int]OperationFunc
	pointer  int
	running  bool
}

type OperationFunc func(program *Program) *Program

func NewProgram(opCodes map[int]OperationFunc, intCodes []int, noun int, verb int) Program {
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

func (p *Program) setValue(pos, value int) {
	p.intCodes[p.intCodes[pos]] = value
}

func (p *Program) setPointer(pos int) {
	p.pointer = pos
}

func (p Program) getIntCodeValue(pos int) int {
	return p.intCodes[p.intCodes[pos]]
}

func (p Program) GetValueAt(pos int) int {
	return p.intCodes[pos]
}

func (p *Program) Execute() {
	p.running = true

	//for ok := true; ok; ok = p.running {
	for p.running {
		opCode := p.intCodes[p.pointer]

		if op, ok := p.opCodes[opCode]; ok {
			op(p)
		} else {
			Stop(p)
		}
	}
}
