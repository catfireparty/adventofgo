package day16

import (
	"fmt"
	"sort"
	"strings"
)

type Grid = [][]string

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var directions = map[Direction]Point{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

type Point struct {
	x int
	y int
}

type FrontierPoint struct {
	point  Point
	cost   int
	facing Direction
}

func getCost(direction Direction, facing Direction) int {
	if direction == facing {
		return 1
	}

	switch direction {
	case North:
		switch facing {
		case East:
			return 1001
		case South:
			return 2001
		case West:
			return 1001
		}
	case East:
		switch facing {
		case North:
			return 1001
		case South:
			return 1001
		case West:
			return 2001
		}
	case West:
		switch facing {
		case North:
			return 1001
		case South:
			return 1001
		case East:
			return 2001
		}
	case South:
		switch facing {
		case North:
			return 2001
		case West:
			return 1001
		case East:
			return 1001
		}
	}

	panic("Could not calculate cost")
}

func PartOne(data string) {
	fmt.Println("Day 16 - Part 1: ")
	grid := parseInput(data)
	start := Point{1, len(grid) - 2}
	if grid[start.y][start.x] != "S" {
		panic("Start not found")
	}

	end := Point{len(grid[0]) - 2, 1}
	if grid[end.y][end.x] != "E" {
		panic("End not found")
	}

	costSoFar := map[string]int{
		getKey(start.x, start.y): 0,
	}

	frontier := []FrontierPoint{{start, 0, East}}
	for {
		current := frontier[0]
		frontier = frontier[1:]

		if current.point.x == end.x && current.point.y == end.y {
			break
		}

		for face, d := range directions {
			x1 := current.point.x + d.x
			y1 := current.point.y + d.y

			if grid[y1][x1] == "#" {
				continue
			}

			nextKey := getKey(x1, y1)
			newCost := costSoFar[getKey(current.point.x, current.point.y)] + getCost(face, current.facing)
			existingCost := costSoFar[nextKey]
			if existingCost == 0 || newCost < existingCost {
				costSoFar[nextKey] = newCost
				frontier = append(frontier, FrontierPoint{Point{x1, y1}, newCost, face})
				sort.Slice(frontier, func(i, j int) bool {
					return frontier[i].cost < frontier[j].cost
				})
			}
		}
	}

	fmt.Println("Minimum cost to reach the end:", costSoFar[getKey(end.x, end.y)])
}

func getKey(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func parseInput(data string) Grid {
	grid := Grid{}
	lines := strings.Split(data, "\n")
	for i := range lines {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func PartTwo(data string) {
	fmt.Println("Day 16 - Part 2: ")
	grid := parseInput(data)
	start := Point{1, len(grid) - 2}
	if grid[start.y][start.x] != "S" {
		panic("Start not found")
	}

	fmt.Println(len(grid))
}
