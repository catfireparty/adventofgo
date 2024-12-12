package day12

import (
	"fmt"
	"strings"
	"time"

	"adventofgo.dev/utils"
)

type Point struct {
	x int
	y int
}

// N, S, E, W - no diagonals
var directions = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Grid [][]string

func (grid Grid) getPlant(point Point) string {
	if point.x < 0 || point.y < 0 || point.y > len(grid)-1 || point.x > len(grid[0])-1 {
		return "*"
	}
	return grid[point.y][point.x]
}

func PartOne(path string) time.Duration {
	fmt.Println("Day 12 - Part 1: ")
	start := time.Now()
	data := utils.ReadFile(path)
	grid := createGrid(data)
	areas := findAreas(grid)

	totalPrice := 0
	for i := range areas {
		// point := areas[i][0]
		perimeter := findPerimeter(areas[i], grid)
		// fmt.Println("Area of range", grid[point.y][point.x], len(areas[i]), "perimeter", perimeter)
		totalPrice += len(areas[i]) * perimeter
	}

	fmt.Println("Total price for fencing:", totalPrice)
	return time.Since(start)
}

func findPerimeter(points []Point, grid Grid) int {
	p := 0
	for i := range points {
		current := points[i]
		currentPlant := grid.getPlant(current)
		for i := range directions {
			x := current.x + directions[i].x
			y := current.y + directions[i].y
			plant := grid.getPlant(Point{x, y})
			if plant != currentPlant {
				p++
			}
		}
	}
	return p
}

func createGrid(data string) Grid {
	lines := strings.Split(data, "\n")
	grid := Grid{}
	for i := range lines {
		if len(lines[i]) > 0 {
			grid = append(grid, strings.Split(lines[i], ""))
		}
	}
	return grid
}

func getKey(point Point) string {
	return fmt.Sprintf("%d:%d", point.x, point.y)
}

func getFrontierPoints(start Point, grid Grid, currentPlant string) []Point {
	frontier := []Point{}
	for i := range directions {
		x := start.x + directions[i].x
		y := start.y + directions[i].y
		point := Point{x, y}
		if grid.getPlant(Point{x, y}) == currentPlant {
			frontier = append(frontier, point)
		}
	}
	return frontier
}

func findAreas(grid Grid) [][]Point {
	// let's iterate through the grid, doing a kind
	// of flood fill with a breadth-first search
	// to get all the areas

	cache := make(map[string]bool)
	areas := [][]Point{}
	for y := range grid {
		for x := range grid[y] {
			start := Point{x, y}

			if cache[getKey(start)] {
				continue
			}

			// use flood fill to find area
			currentPlant := grid[y][x]
			currentArea := []Point{{x, y}}
			cache[getKey(start)] = true

			frontier := getFrontierPoints(start, grid, currentPlant)
			for {
				if len(frontier) == 0 {
					break
				}

				newFrontier := []Point{}
				for i := range frontier {
					point := frontier[i]
					if cache[getKey(point)] {
						continue
					}

					currentArea = append(currentArea, point)
					cache[getKey(point)] = true

					newPoints := getFrontierPoints(point, grid, currentPlant)
					if len(newPoints) > 0 {
						newFrontier = append(newFrontier, newPoints...)
					}
				}

				frontier = newFrontier
			}

			areas = append(areas, currentArea)
		}
	}
	return areas
}

func findCorners(area []Point, grid Grid) int {
	// concave and convex corners

	corners := 0
	for i := range area {
		current := area[i]
		plant := grid[current.y][current.x]
		x := current.x
		y := current.y

		n := grid.getPlant(Point{x, y + 1})
		ne := grid.getPlant(Point{x + 1, y + 1})
		e := grid.getPlant(Point{x + 1, y})
		se := grid.getPlant(Point{x + 1, y - 1})
		s := grid.getPlant(Point{x, y - 1})
		sw := grid.getPlant(Point{x - 1, y - 1})
		w := grid.getPlant(Point{x - 1, y})
		nw := grid.getPlant(Point{x - 1, y + 1})

		// convex corners
		if n != plant && e != plant {
			corners++
		}
		if e != plant && s != plant {
			corners++
		}
		if s != plant && w != plant {
			corners++
		}
		if w != plant && n != plant {
			corners++
		}

		// concave corners
		if n == plant && w == plant && nw != plant {
			corners++
		}
		if n == plant && e == plant && ne != plant {
			corners++
		}
		if s == plant && e == plant && se != plant {
			corners++
		}
		if s == plant && w == plant && sw != plant {
			corners++
		}
	}

	return corners
}

func PartTwo(path string) time.Duration {
	fmt.Println("Day 12 - Part 2: ")
	start := time.Now()
	data := utils.ReadFile(path)
	grid := createGrid(data)
	areas := findAreas(grid)

	totalPrice := 0
	for i := range areas {
		// point := areas[i][0]
		corners := findCorners(areas[i], grid)
		// fmt.Println("Area of range", grid[point.y][point.x], len(areas[i]), "corners", corners)
		totalPrice += len(areas[i]) * corners
	}

	fmt.Println("Total price for fencing:", totalPrice)
	return time.Since(start)
}
