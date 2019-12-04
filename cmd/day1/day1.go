package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"log"
)

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func sum(fuels ...int) int {
	fs := 0
	for _, f := range fuels {
		fs += f
	}

	return fs
}

func main() {
	rawMasses := files.Load("cmd/day1/input1.txt")
	var modules []int

	for _, rm := range rawMasses {
		m := converters.StringToInt(rm)
		modules = append(modules, calculateFuel(m))
	}

	log.Print(sum(modules...))
}
