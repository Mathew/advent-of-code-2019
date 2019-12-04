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

func calculateModuleFuel(f int) int {
	tf := 0

	for ok := true; ok; ok = f > 0 {
		f = calculateFuel(f)

		if f > 0 {
			tf += f
		}
	}

	return tf
}

func main() {
	rawMasses := files.Load("cmd/day1/input1.txt", "\n")
	var modules []int
	var totalModules []int

	for _, rm := range rawMasses {
		m := converters.StringToInt(rm)
		modules = append(modules, calculateFuel(m))
		totalModules = append(totalModules, calculateModuleFuel(m))
	}

	log.Printf("Total fuel (part 1): %v", sum(modules...))
	log.Printf("Total module fuel (part 2): %v", sum(totalModules...))
}
