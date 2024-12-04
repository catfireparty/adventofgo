package day4

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

// 8 search possibilities with 3 coords each
var XmasCoords = [8][3][2]int{
	{{1, 0}, {2, 0}, {3, 0}},
	{{1, 1}, {2, 2}, {3, 3}},
	{{0, 1}, {0, 2}, {0, 3}},
	{{-1, 1}, {-2, 2}, {-3, 3}},
	{{-1, 0}, {-2, 0}, {-3, 0}},
	{{-1, -1}, {-2, -2}, {-3, -3}},
	{{0, -1}, {0, -2}, {0, -3}},
	{{1, -1}, {2, -2}, {3, -3}},
}

const toFind string = "MAS"

func PartOne(path string) {
	fmt.Println("Day 4 - Part 1: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)

	xmasCount := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != "X" {
				continue
			}

			// check search coords
			for i := range XmasCoords {
				found := true
				for j := range XmasCoords[i] {
					x0 := x + XmasCoords[i][j][0]
					y0 := y + XmasCoords[i][j][1]
					if x0 < 0 || y0 < 0 || x0 >= len(grid) || y0 >= len(grid[x]) {
						found = false
						break
					}

					if grid[x0][y0] != string(toFind[j]) {
						found = false
						break
					}
				}
				if found {
					xmasCount++
				}
			}
		}
	}

	fmt.Println("Found: ", xmasCount)
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

var crossmasCoords = [][][]int{
	{{1, 1}, {-1, -1}},
	{{-1, 1}, {1, -1}},
}

func PartTwo(path string) {
	fmt.Println("Day 4 - Part 2: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)

	xmasCount := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != "A" {
				continue
			}

			if checkCoords(grid, x, y, crossmasCoords) {
				xmasCount++
			}
		}
	}

	fmt.Println("Found: ", xmasCount)
}

func checkCoords(grid [][]string, x int, y int, xmasCoords [][][]int) bool {
	for i := range xmasCoords {
		x0 := x + xmasCoords[i][0][0]
		y0 := y + xmasCoords[i][0][1]

		x1 := x + xmasCoords[i][1][0]
		y1 := y + xmasCoords[i][1][1]

		if x0 < 0 || y0 < 0 || x0 >= len(grid) || y0 >= len(grid[x]) {
			return false
		}

		if x1 < 0 || y1 < 0 || x1 >= len(grid) || y1 >= len(grid[x]) {
			return false
		}

		maybeMas := strings.Join([]string{grid[x0][y0], "A", grid[x1][y1]}, "")
		if maybeMas == "MAS" || maybeMas == "SAM" {
			continue
		}

		return false
	}
	return true
}
