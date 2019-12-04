package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"log"
	"strconv"
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
	var masses []int

	for _, m := range rawMasses {
		r, err := strconv.Atoi(m)
		if err != nil {
			log.Fatalf("Cannot convert: %s", m)
		}

		masses = append(masses, calculateFuel(r))
	}

	log.Print(sum(masses...))
}
