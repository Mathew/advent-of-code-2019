package structures

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestArrayToMatrices(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	exp := [][][]int{
		{
			{1, 2, 3,},
			{4, 5, 6,},
		}, {
			{7, 8, 9,},
			{0, 1, 2,},
		},
	}
	asserts.Equals(t, exp, ArrayToMatrices(arr, 3, 2))
}
