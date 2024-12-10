package day10

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Point struct {
	x int
	y int
}

var directions = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func PartOne(path string) {
	fmt.Println("Day 10 - Part 1: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)

	totalScore := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				score := findPaths(Point{x, y}, grid, false)
				totalScore += score
			}
		}
	}

	fmt.Println("Total path score:", totalScore)
}

func getFrontierPoints(start Point, grid [][]int, next int) []Point {
	frontier := []Point{}
	for i := range directions {
		x := start.x + directions[i].x
		y := start.y + directions[i].y
		point := Point{x, y}
		if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[0])-1 {
			continue
		}
		if grid[y][x] == next {
			frontier = append(frontier, point)
		}
	}
	return frontier
}

func getKey(point Point) string {
	return fmt.Sprintf("%d:%d", point.x, point.y)
}

func findPaths(start Point, grid [][]int, findDistinct bool) int {
	numPaths := 0
	current := 1
	frontier := getFrontierPoints(start, grid, current)

	for {
		if len(frontier) == 0 {
			break
		}

		newFrontier := []Point{}
		tried := make(map[string]bool)
		for i := range frontier {
			point := frontier[i]
			if !findDistinct {
				if tried[getKey(point)] {
					continue
				}
				tried[getKey(point)] = true
			}
			if current == 9 {
				numPaths++
			} else {
				newFrontier = append(newFrontier, getFrontierPoints(point, grid, current+1)...)
			}
		}
		current++
		frontier = newFrontier
	}

	return numPaths
}

func createGrid(data string) [][]int {
	lines := strings.Split(data, "\n")
	grid := [][]int{}
	for i := range lines {
		if len(lines[i]) > 0 {
			grid = append(grid, utils.ToIntArray(strings.Split(lines[i], "")))
		}
	}
	return grid
}

func PartTwo(path string) {
	fmt.Println("Day 10 - Part 2: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)

	totalScore := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				score := findPaths(Point{x, y}, grid, true)
				totalScore += score
			}
		}
	}

	fmt.Println("Total path score:", totalScore)
}
