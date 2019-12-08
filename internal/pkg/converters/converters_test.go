package converters

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestStringToInt(t *testing.T) {
	asserts.Equals(t, 6, StringToInt("6"))
}

func TestStringsToInts(t *testing.T) {
	asserts.Equals(t, []int{6, 5, 3}, StringsToInts("6", "5", "3"))
}

func TestIntToDigits(t *testing.T) {
	asserts.Equals(t, []int{1}, IntToDigits(1))
	asserts.Equals(t, []int{6, 5, 4}, IntToDigits(654))
	asserts.Equals(t, []int{6, 5, 4}, IntToDigits(654))
}

func TestReverse(t *testing.T) {
	asserts.Equals(t, []int{6, 5, 4}, Reverse([]int{4, 5, 6}))
}