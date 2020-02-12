package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/vectors"
	"log"
)

type StarChart struct {
	chart     [][]int
	asteroids []vectors.Point
}

func (sc StarChart) Draw() {
	for y := 0; y < len(sc.chart[0]); y++ {
		for x := 0; x < len(sc.chart); x++ {
			fmt.Printf(" %v ", sc.chart[x][y])
		}
		fmt.Println("")
	}
}

func (sc StarChart) FindMonitoringStationLocation() (vectors.Point, int) {
	highest := 0
	ast := vectors.NewPoint(0, 0)

	for _, a := range sc.asteroids {
		uniqueSlopes := map[string]map[float64]bool{
			"x+y+": {},
			"x-y+": {},
			"x+y-": {},
			"x-y-": {},
		}

		for _, b := range sc.asteroids {
			if a == b {
				continue
			}

			v := b.Subtract(a)
			x := ""
			y := ""
			if b.X < a.X {x = "x-"} else {x = "x+"}
			if b.Y < a.Y {y = "y-"} else {y = "y+"}
			q := fmt.Sprintf("%v%v", x, y)

			uniqueSlopes[q][v.Slope()] = true
		}

		total := 0
		for _, us := range uniqueSlopes {
			total += len(us)
		}

		if total > highest {
			highest = total
			ast = a
		}
	}

	return ast, highest
}

func NewStarChart(rawCharts []string) StarChart {
	ylen := len(rawCharts)
	xlen := len(rawCharts[0])

	var asteroids []vectors.Point

	chart := make([][]int, ylen)
	for i := range chart {
		chart[i] = make([]int, xlen)
	}

	for y, line := range rawCharts {
		for x, c := range line {
			if string(c) == "#" {
				chart[x][y] = 1
				asteroids = append(asteroids, vectors.NewPoint(x, y))
			}
		}
	}

	return StarChart{
		chart,
		asteroids,
	}
}

func main() {
	asteroidBeltraw := files.Load("cmd/day10/input.txt", "\n")
	chart := NewStarChart(asteroidBeltraw)
	log.Print(chart.FindMonitoringStationLocation())
}
