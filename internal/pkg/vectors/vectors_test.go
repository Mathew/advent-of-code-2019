package vectors

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestVector_Slope(t *testing.T) {
	v := NewVector(3, 2)
	asserts.Equals(t, 0.6666666666666666, v.Slope())
}
