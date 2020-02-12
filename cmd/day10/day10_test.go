package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"github.com/mathew/advent-of-code-2019/internal/pkg/vectors"
	"testing"
)

var testCases = []struct {
	input    []string
	expPoint vectors.Point
	expCount int
}{
	{
		[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		},
		vectors.NewPoint(3, 4),
		8,
	}, {
		[]string{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		},
		vectors.NewPoint(5, 8),
		33,
	}, {
		[]string{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		},
		vectors.NewPoint(1, 2),
		35,
	}, {
		[]string{
			".#..#..###",
			"####.###.#",
			"....###.#.",
			"..###.##.#",
			"##.##.#.#.",
			"....###..#",
			"..#.#..#.#",
			"#..#.#.###",
			".##...##.#",
			".....#.#..,",
		},
		vectors.NewPoint(6, 3),
		41,
	}, {
		[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
		vectors.NewPoint(11, 13),
		210,
	},
}

func TestMakeStarChart(t *testing.T) {
	exp := [][]int{
		{0, 0, 1, 0, 0,},
		{1, 0, 1, 0, 0,},
		{0, 0, 1, 0, 0,},
		{0, 0, 1, 0, 1,},
		{1, 0, 1, 1, 1,},
	}
	chart := NewStarChart(testCases[0].input)
	asserts.Equals(t, exp, chart.chart)
}

func TestStarChart_FindMonitoringStationLocation(t *testing.T) {
	for _, tc := range testCases {
		chart := NewStarChart(tc.input)
		astPoint, astCount := chart.FindMonitoringStationLocation()
		asserts.Equals(t, tc.expPoint, astPoint)
		asserts.Equals(t, tc.expCount, astCount)
	}
}
