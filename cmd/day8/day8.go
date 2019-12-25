package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/converters"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/structures"
	"log"
	"math"
)

type Layer struct {
	matrix [][]int
}

func NewLayer(matrix [][]int) Layer {
	return Layer{matrix}
}

func (l Layer) NumOf(i int) int {
	count := 0

	for _, row := range l.matrix {
		for _, d := range row {
			if d == i {
				count += 1
			}
		}
	}

	return count
}

func (l Layer) MultipleOfNumsCount(i, j int) int {
	return l.NumOf(i) * l.NumOf(j)
}

type Image struct {
	layers []Layer
	width  int
	height int
}

func NewImage(layers []Layer, width, height int) Image {
	return Image{layers, width, height}
}

func rawToImage(raw string, width, height int) Image {
	numbers := converters.StringToDigits(raw)
	matrices := structures.ArrayToMatrices(numbers, width, height)

	var layers []Layer

	for _, m := range matrices {
		layers = append(layers, NewLayer(m))
	}

	return NewImage(layers, width, height)
}

func (img Image) FewestNumLayer(i int) Layer {
	var layer Layer
	smallest := math.MaxInt64

	for _, l := range img.layers {
		c := l.NumOf(i)
		if c < smallest {
			smallest = c
			layer = l
		}
	}

	return layer
}

func main() {
	raw := files.Load("cmd/day8/input.txt", "")
	image := rawToImage(raw[0], 25, 6)

	smallestLayer := image.FewestNumLayer(0)
	result := smallestLayer.MultipleOfNumsCount(1, 2)
	log.Printf("Multiple of 1 + 2 present in Layer: %v", result)
}
