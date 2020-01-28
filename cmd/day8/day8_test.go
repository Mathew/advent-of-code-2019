package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

func TestLayer_MultipleOfNumsCount(t *testing.T) {
	l := NewLayer([][]int{
		{1, 2, 2,},
		{3, 4, 5,},
	})

	asserts.Equals(t, 2, l.MultipleOfNumsCount(1, 2))
}

func TestLayer_NumOf(t *testing.T) {
	l := NewLayer([][]int{
		{1, 2, 2,},
		{3, 4, 5,},
	})

	asserts.Equals(t, 2, l.NumOf(2))
	asserts.Equals(t, 1, l.NumOf(5))
}

func TestImage_FewestNumLayer(t *testing.T) {
	l1 := NewLayer([][]int{
		{1, 2, 0,},
		{3, 0, 5,},
	})
	l2 := NewLayer([][]int{
		{1, 0, 2,},
		{3, 4, 5,},
	})

	img := NewImage([]Layer{l1, l2}, 1, 1)
	asserts.Equals(t, l2, img.FewestNumLayer(0))
}
