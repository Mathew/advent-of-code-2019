package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestPassword_hasAdjancentNumbers(t *testing.T) {
	p := NewPassword(110000)
	asserts.Equals(t, true, p.hasTwoAdjacentNumbers())

	p = NewPassword(123456)
	asserts.Equals(t, false, p.hasTwoAdjacentNumbers())
}

func TestPassword_eachNumberIncreases(t *testing.T) {
	p := NewPassword(123456)
	asserts.Equals(t, true, p.eachNumberIncreases())

	p = NewPassword(113456)
	asserts.Equals(t, true, p.eachNumberIncreases())

	p = NewPassword(153456)
	asserts.Equals(t, false, p.eachNumberIncreases())
}
