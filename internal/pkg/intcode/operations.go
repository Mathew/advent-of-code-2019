package intcode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	input, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Could not convert input to integer. %v. %v", text, err)
	}

	p.setValue(p.pointer+1, input)
	p.setPointer(p.pointer + 2)

	return p
}

func Output(p *Program, modes []int) *Program {
	output := p.getIntCodeValue(p.pointer+1, modes[0])
	p.setPointer(p.pointer + 2)
	log.Printf("Output: %v \n", output)

	return p
}
