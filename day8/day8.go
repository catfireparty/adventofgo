package day8

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Point struct {
	x int
	y int
}

func PartOne(path string) {
	fmt.Println("Day 8 - Part 1: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)
	antennae := findAntennae(grid)
	antinodes := findAntinodes(antennae)

	uniqueLocations := 0
	for i := range antinodes {
		x := antinodes[i].x
		y := antinodes[i].y

		if y >= len(grid) || x >= len(grid[0]) || x < 0 || y < 0 {
			continue
		}

		if grid[y][x] != "#" {
			uniqueLocations++
			grid[y][x] = "#"
		}
	}

	// print out the path used
	// for i := range grid {
	// 	fmt.Println(strings.Join(grid[i], ""))
	// }

	fmt.Println("Unique locations:", uniqueLocations)
}

func createGrid(data string) [][]string {
	lines := strings.Split(data, "\n")
	grid := [][]string{}
	for i := range lines {
		if len(lines[i]) > 0 {
			grid = append(grid, strings.Split(lines[i], ""))
		}
	}
	return grid
}

func findAntennae(grid [][]string) map[string][]Point {
	antennae := make(map[string]([]Point))
	for y := range grid {
		for x := range grid[y] {
			content := grid[y][x]
			if content == "." {
				continue
			}

			antennae[content] = append(antennae[content], Point{x, y})
		}
	}
	return antennae
}

// for each pair of antennae, find the antinodes
// these occur along the line connecting them
// when the distance to one is twice that to the other
// i.e. X1---A---A---X2
//
// So, given two points, (x0, y0) and (x1, y1)
// where x1 > x0 and y1 > y0
// dx = x1 - x0
// dy = y1 - y0
// X2 = (x1 + dx, y1 + dy)
// X1 = (x0 - dx, y0 - dy)
// y1 will always be >= y0
// x1 might be less or greater so have to check
func findAntinodes(antennae map[string][]Point) []Point {
	antinodes := []Point{}
	for _, points := range antennae {
		nodes := []Point{}
		// iterate through every pair
		for i := range points {
			for j := i + 1; j < len(points); j++ {
				x0 := points[i].x
				x1 := points[j].x
				y0 := points[i].y
				y1 := points[j].y

				dx := x1 - x0
				dy := y1 - y0

				nodes = append(nodes, Point{x1 + dx, y1 + dy}, Point{x0 - dx, y0 - dy})
			}
		}

		// fmt.Println(key, nodes)
		antinodes = append(antinodes, nodes...)
	}

	return antinodes
}

func PartTwo(path string) {
	fmt.Println("Day 8 - Part 2: ")
	data := utils.ReadFile(path)
	fmt.Println(len(data))
}
