package day16

import (
	"fmt"
	"math"
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

type Cost struct {
	cost int
	from []FrontierPoint
}

func getFaceKey(x int, y int, face Direction) string {
	return fmt.Sprintf("%d:%d:%d", x, y, face)
}

func PartTwo(data string) {
	fmt.Println("Day 16 - Part 2: ")
	grid := parseInput(data)
	start := Point{1, len(grid) - 2}
	if grid[start.y][start.x] != "S" {
		panic("Start not found")
	}

	end := Point{len(grid[0]) - 2, 1}
	if grid[end.y][end.x] != "E" {
		panic("End not found")
	}

	costSoFar := map[string]*Cost{}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != "#" {
				costSoFar[getFaceKey(x, y, North)] = &Cost{math.MaxInt32, []FrontierPoint{}}
				costSoFar[getFaceKey(x, y, South)] = &Cost{math.MaxInt32, []FrontierPoint{}}
				costSoFar[getFaceKey(x, y, East)] = &Cost{math.MaxInt32, []FrontierPoint{}}
				costSoFar[getFaceKey(x, y, West)] = &Cost{math.MaxInt32, []FrontierPoint{}}
			}
		}
	}

	// set the start square to zero cost
	costSoFar[getFaceKey(start.x, start.y, East)] = &Cost{0, []FrontierPoint{}}

	frontier := []FrontierPoint{{start, 0, East}}
	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]
		if current.point.x == end.x && current.point.y == end.y {
			// we've found one way
			continue
		}

		for face, d := range directions {
			x1 := current.point.x + d.x
			y1 := current.point.y + d.y

			if grid[y1][x1] == "#" {
				continue
			}

			nextKey := getFaceKey(x1, y1, face)
			currentKey := getFaceKey(current.point.x, current.point.y, current.facing)
			currentCost := costSoFar[currentKey]

			newCost := currentCost.cost + getCost(face, current.facing)
			existingCost := costSoFar[nextKey]

			if existingCost.cost == newCost {
				// add current point to from list
				existingCost.from = append(existingCost.from, current)
				continue
			}

			if newCost < existingCost.cost {
				costSoFar[nextKey] = &Cost{newCost, []FrontierPoint{{Point{current.point.x, current.point.y}, current.cost, current.facing}}}
				frontier = append(frontier, FrontierPoint{Point{x1, y1}, newCost, face})
				sort.Slice(frontier, func(i, j int) bool {
					return frontier[i].cost < frontier[j].cost
				})
			}
		}
	}

	// work back from the end and fill in all the paths?
	points := []FrontierPoint{{Point{end.x, end.y}, 0, North}}
	for len(points) > 0 {
		current := points[0]
		points = points[1:]
		grid[current.point.y][current.point.x] = "O"
		currentCost := costSoFar[getFaceKey(current.point.x, current.point.y, current.facing)]
		if len(currentCost.from) > 0 {
			points = append(points, currentCost.from...)
		}
	}

	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				count++
			}
		}
	}

	// printGrid(grid)

	fmt.Println("Number of tiles:", count)
}

func printGrid(grid Grid) {
	for y := range grid {
		fmt.Println(strings.Join(grid[y], ""))
	}
}
