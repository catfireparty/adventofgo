package day18

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Point struct {
	x int
	y int
}

type Grid [71][71]string

// type Grid [7][7]string

var directions = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type FrontierPoint struct {
	point Point
	from  []FrontierPoint
}

func PartOne(data string) {
	fmt.Println("Day 18 - Part 1: ")
	corrupted := parseData(data)
	grid := Grid{}

	// populate the "free" squares
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = "."
		}
	}

	// populate the corruption
	for i := range 1024 {
		c := corrupted[i]
		grid[c.y][c.x] = "#"
	}

	// printGrid(grid)

	path, _ := traverseGrid(grid)

	current := path
	count := 0
	for {
		grid[current.point.y][current.point.x] = "O"
		if len(current.from) == 0 {
			break
		} else {
			current = current.from[0]
			count++
		}
	}

	// printGrid(grid)

	fmt.Println("Minimum steps:", count)
}

func traverseGrid(grid Grid) (FrontierPoint, bool) {
	yMax := len(grid) - 1
	xMax := len(grid[0]) - 1
	start := Point{0, 0}
	end := Point{len(grid) - 1, len(grid[0]) - 1}
	frontier := []FrontierPoint{{start, []FrontierPoint{}}}
	visited := make(map[string]bool)
	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if current.point.x == end.x && current.point.y == end.y {
			return current, true
		}

		for i := range directions {
			d := directions[i]
			x1 := current.point.x + d.x
			y1 := current.point.y + d.y

			key := getKey(Point{x1, y1})
			if visited[key] {
				continue
			}

			if x1 < 0 || y1 < 0 || x1 > xMax || y1 > yMax {
				continue
			}

			if grid[y1][x1] == "#" {
				continue
			}

			frontier = append(frontier, FrontierPoint{Point{x1, y1}, []FrontierPoint{current}})
			visited[key] = true
		}
	}

	return FrontierPoint{}, false
}

func getKey(point Point) string {
	return fmt.Sprintf("%d:%d", point.x, point.y)
}

func parseData(data string) []Point {
	points := []Point{}
	lines := strings.Split(data, "\n")
	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		coords := strings.Split(lines[i], ",")
		points = append(points, Point{utils.ToInt(coords[0]), utils.ToInt(coords[1])})
	}
	return points
}

func printGrid(grid Grid) {
	for y := range grid {
		fmt.Println(strings.Join(grid[y][:], ""))
	}
}

func PartTwo(data string) {
	fmt.Println("Day 18 - Part 2: ")
	corrupted := parseData(data)
	grid := Grid{}

	// populate the "free" squares
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = "."
		}
	}

	// populate the corruption
	for i := range 1024 {
		c := corrupted[i]
		grid[c.y][c.x] = "#"
	}

	breaking := Point{}
	for i := 1024; i < len(corrupted); i++ {
		c := corrupted[i]
		grid[c.y][c.x] = "#"
		_, possible := traverseGrid(grid)
		if !possible {
			breaking = corrupted[i]
			break
		}
	}

	fmt.Println("Breaking corruption:", fmt.Sprintf("%d,%d", breaking.x, breaking.y))
}
