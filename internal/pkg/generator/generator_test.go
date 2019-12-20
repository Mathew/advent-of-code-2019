package generator

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestCreateIntGenerator(t *testing.T) {
	gen := CreateIntGenerator(1, 2, 3)
	r, ok := gen()
	asserts.Equals(t, 1, r)
	asserts.Equals(t, true, ok)

	r, ok = gen()
	asserts.Equals(t, 2, r)
	asserts.Equals(t, true, ok)

	r, ok = gen()
	asserts.Equals(t, 3, r)
	asserts.Equals(t, true, ok)

	r, ok = gen()
	asserts.Equals(t, 0, r)
	asserts.Equals(t, false, ok)
}
