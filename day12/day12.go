package day12

import (
	"fmt"
	"strings"

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

func PartOne(path string) {
	fmt.Println("Day 12 - Part 1: ")
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
}

func findPerimeter(points []Point, grid [][]string) int {
	p := 0
	for i := range points {
		start := points[i]
		for i := range directions {
			x := start.x + directions[i].x
			y := start.y + directions[i].y
			if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[0])-1 {
				p++
			} else if grid[y][x] != grid[start.y][start.x] {
				p++
			}
		}
	}
	return p
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

func getKey(point Point) string {
	return fmt.Sprintf("%d:%d", point.x, point.y)
}

func getFrontierPoints(start Point, grid [][]string, currentPlant string) []Point {
	frontier := []Point{}
	for i := range directions {
		x := start.x + directions[i].x
		y := start.y + directions[i].y
		point := Point{x, y}
		if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[0])-1 {
			continue
		}
		if grid[y][x] == currentPlant {
			frontier = append(frontier, point)
		}
	}
	return frontier
}

func findAreas(grid [][]string) [][]Point {
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

func getPlant(x int, y int, grid [][]string) string {
	if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[0])-1 {
		return "*"
	}
	return grid[y][x]
}

func findCorners(area []Point, grid [][]string) int {
	// concave and convex corners

	corners := 0
	for i := range area {
		current := area[i]
		plant := grid[current.y][current.x]
		x := current.x
		y := current.y

		n := getPlant(x, y+1, grid)
		ne := getPlant(x+1, y+1, grid)
		e := getPlant(x+1, y, grid)
		se := getPlant(x+1, y-1, grid)
		s := getPlant(x, y-1, grid)
		sw := getPlant(x-1, y-1, grid)
		w := getPlant(x-1, y, grid)
		nw := getPlant(x-1, y+1, grid)

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

func PartTwo(path string) {
	fmt.Println("Day 12 - Part 2: ")
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
}
