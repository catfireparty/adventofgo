package day6

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Direction int

var directions = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

const (
	North Direction = iota
	East
	South
	West
)

type Point struct {
	x int
	y int
}

type Vector struct {
	point     Point
	direction Direction
}

func PartOne(path string) {
	fmt.Println("Day 6 - Part 1: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)
	current := findGuard(grid)

	outOfBounds := false
	for {
		grid[current.point.y][current.point.x] = "X"
		current, outOfBounds = getNextVector(grid, current)
		if outOfBounds {
			break
		}
	}

	// print out the path used
	// for i := range grid {
	// 	fmt.Println(strings.Join(grid[i], ""))
	// }

	fmt.Println("Unique positions visited: ", getUniquePositions(grid))
}

func getUniquePositions(grid [][]string) int {
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "X" {
				count++
			}
		}
	}
	return count
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

func getNextVector(grid [][]string, vector Vector) (Vector, bool) {
	direction := vector.direction
	for {
		d := directions[direction]
		v0 := Vector{point: Point{x: vector.point.x + d[0], y: vector.point.y + d[1]}, direction: direction}
		if v0.point.x < 0 || v0.point.y < 0 || v0.point.x >= len(grid[0]) || v0.point.y >= len(grid) {
			return v0, true
		}

		// check the grid square
		if grid[v0.point.y][v0.point.x] == "#" {
			direction = (direction + 1) % 4
			continue
		}

		return v0, false
	}
}

func findGuard(grid [][]string) Vector {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "^" {
				return Vector{Point{x, y}, North}
			}
		}
	}

	panic("Guard not found!")
}

func PartTwo(path string) {
	fmt.Println("Day 6 - Part 2: ")
	data := utils.ReadFile(path)
	grid := createGrid(data)
	start := findGuard(grid)

	// the only possible positions are ones on the route the guard currently takes
	possible := []Point{}
	visited := make(map[string]bool)
	current := Vector{Point{x: start.point.x, y: start.point.y}, start.direction}
	outOfBounds := false
	for {
		key := fmt.Sprintf(`%d:%d`, current.point.x, current.point.y)
		if !visited[key] {
			possible = append(possible, current.point)
			visited[key] = true
		}
		current, outOfBounds = getNextVector(grid, current)
		if outOfBounds {
			break
		}
	}

	// we know we're in a loop when we're at a previously visited location facing the same way as before
	// iterate through all possible positions and block them off, see if they loop

	obstaclePositions := []Point{}
	for i := range possible {
		test := possible[i]

		if test.x == start.point.x && test.y == start.point.y {
			// ignore the starting position
			continue
		}

		grid[test.y][test.x] = "#"

		current := Vector{Point{x: start.point.x, y: start.point.y}, North}
		visited := make(map[string]bool)
		for {
			key := fmt.Sprintf(`%d:%d:%d`, current.point.x, current.point.y, current.direction)
			if visited[key] {
				// we're in a loop
				obstaclePositions = append(obstaclePositions, Point{x: test.x, y: test.y})
				break
			}
			visited[key] = true
			current, outOfBounds = getNextVector(grid, current)
			if outOfBounds {
				break
			}
		}

		grid[test.y][test.x] = "."
	}

	fmt.Println("Possible obstacle positions: ", len(obstaclePositions))
}
