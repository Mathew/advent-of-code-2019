package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"log"
	"strconv"
	"strings"
)

const puzzleInput = "387638-919123"

type Password struct {
	digits []int
}

func NewPassword(n int) Password {
	if n < 100000 || n > 999999 {
		log.Fatal("Invalid password")
	}

	return Password{converters.Reverse(converters.IntToDigits(n))}
}

func (p Password) hasTwoAdjacentNumbers() bool {
	for x := 1; x < len(p.digits); x++ {
		if p.digits[x-1] == p.digits[x] {
			return true
		}
	}

	return false
}

func (p Password) eachNumberIncreases() bool {
	for x := 1; x < len(p.digits); x++ {
		if p.digits[x-1] > p.digits[x] {
			return false
		}
	}

	return true
}

func (p Password) IsValid() bool {
	return p.eachNumberIncreases() && p.hasTwoAdjacentNumbers()
}

func main() {
	r := strings.Split(puzzleInput, "-")
	rawMin, rawMax := r[0], r[1]
	min, _ := strconv.Atoi(rawMin)
	max, _ := strconv.Atoi(rawMax)

	c := 0
	for x := min; x <= max; x++ {
		p := NewPassword(x)
		if p.IsValid() {
			c += 1
		}
	}

	log.Printf("Valid Password Count: %v", c)
}
