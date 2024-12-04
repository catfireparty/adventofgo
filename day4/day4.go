package day4

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

// 8 search possibilities with 3 coords each
var searchCoords = [8][3][2]int{
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
			for i := range searchCoords {
				found := true
				for j := range searchCoords[i] {
					x0 := x + searchCoords[i][j][0]
					y0 := y + searchCoords[i][j][1]
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

func PartTwo(path string) {
	fmt.Println("Day 4 - Part 2: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)
	fmt.Println(len(grid))
}
