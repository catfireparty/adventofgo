package day15

import (
	"fmt"
	"strings"
)

type Grid = [][]string

type Point struct {
	x int
	y int
}

var directions = map[string]Point{
	"^": {0, -1},
	"v": {0, 1},
	">": {1, 0},
	"<": {-1, 0},
}

func PartOne(data string) {
	fmt.Println("Day 15 - Part 1: ")
	grid, moves := parseData(data)
	current := findRobot(grid)

	for i := range moves {
		move := moves[i]
		direction := directions[move]

		queue := []Point{current}
		next := Point{current.x + direction.x, current.y + direction.y}
		for {
			content := grid[next.y][next.x]

			if content == "." {
				break
			} else if content == "#" {
				queue = []Point{}
				break
			} else if content == "O" {
				queue = append(queue, next)
			} else {
				panic("Found unexpected content")
			}

			next = Point{next.x + direction.x, next.y + direction.y}
		}

		// move the queue items over one
		// in reverse order
		for i := len(queue) - 1; i >= 0; i-- {
			current = queue[i]
			x := current.x + direction.x
			y := current.y + direction.y

			grid[y][x] = grid[current.y][current.x]
			grid[current.y][current.x] = "."
		}

		if grid[current.y][current.x] == "." {
			current = Point{current.x + direction.x, current.y + direction.y}
		}

		// printGrid(grid)
	}

	fmt.Println("Sum of all GPS coords:", sumGPSCoords(grid))
}

func printGrid(grid Grid) {
	for y := range grid {
		fmt.Println(strings.Join(grid[y], ""))
	}
}

func findRobot(grid Grid) Point {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "@" {
				return Point{x, y}
			}
		}
	}

	panic("Robot not found!")
}

func sumGPSCoords(grid Grid) int {
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" || grid[y][x] == "[" {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func parseData(data string) (Grid, []string) {
	moves := []string{}
	grid := Grid{}

	lines := strings.Split(data, "\n")
	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, "")

		if len(parts) == 0 {
			continue
		}

		if parts[0] == "#" {
			grid = append(grid, parts)
		} else {
			moves = append(moves, parts...)
		}
	}

	return grid, moves
}

func makeGridWider(grid Grid) Grid {
	newGrid := Grid{}
	for y := range grid {
		newGrid = append(newGrid, []string{})

		for x := range grid[y] {
			content := grid[y][x]

			switch content {
			case "#":
				newGrid[y] = append(newGrid[y], "#", "#")
			case "O":
				newGrid[y] = append(newGrid[y], "[", "]")
			case "@":
				newGrid[y] = append(newGrid[y], "@", ".")
			case ".":
				newGrid[y] = append(newGrid[y], ".", ".")
			}
		}
	}
	return newGrid
}

func getKey(point Point) string {
	return fmt.Sprintf("%d:%d", point.x, point.y)
}

func PartTwo(data string) {
	fmt.Println("Day 15 - Part 2: ")
	grid, moves := parseData(data)
	grid = makeGridWider(grid)
	current := findRobot(grid)

	for i := range moves {
		// printGrid(grid)
		move := moves[i]
		// fmt.Println("next move:", move)
		direction := directions[move]

		queue := [][]Point{{current}}
		for {
			next := []Point{}
			frontier := queue[len(queue)-1]
			for i := range frontier {
				x := direction.x + frontier[i].x
				y := direction.y + frontier[i].y
				next = append(next, Point{x, y})
			}

			// if all next is empty, we can break and move
			// if any of next is a wall we break and move nothing
			// if any of next is a box, we add those to the queue

			allEmpty := true
			blocked := false
			added := make(map[string]bool)
			toAdd := []Point{}
			for i := range next {
				content := grid[next[i].y][next[i].x]
				// fmt.Println(content, next[i])

				if content == "." {
					continue
				} else if content == "#" {
					allEmpty = false
					blocked = true
					break
				} else if content == "[" {
					allEmpty = false
					if !added[getKey(next[i])] {
						toAdd = append(toAdd, next[i])
						added[getKey(next[i])] = true
					}

					// if direction is up or down, need to add the one next to it
					if move == "^" || move == "v" {
						box2 := Point{next[i].x + 1, next[i].y}
						if !added[getKey(box2)] {
							toAdd = append(toAdd, box2)
							added[getKey(box2)] = true
						}
					}
				} else if content == "]" {
					allEmpty = false
					if !added[getKey(next[i])] {
						toAdd = append(toAdd, next[i])
						added[getKey(next[i])] = true
					}

					// if direction is up or down, need to add the one next to it
					if move == "^" || move == "v" {
						box2 := Point{next[i].x - 1, next[i].y}
						if !added[getKey(box2)] {
							toAdd = append(toAdd, box2)
							added[getKey(box2)] = true
						}
					}
				} else {
					panic("Found unexpected content")
				}
			}

			if allEmpty {
				break
			}

			if blocked {
				queue = [][]Point{{}}
				break
			}

			queue = append(queue, toAdd)
		}

		// move the queue items over one
		// in reverse order
		for i := len(queue) - 1; i >= 0; i-- {
			movable := queue[i]
			for j := range movable {
				toMove := movable[j]
				x := toMove.x + direction.x
				y := toMove.y + direction.y

				grid[y][x] = grid[toMove.y][toMove.x]
				grid[toMove.y][toMove.x] = "."
			}

		}

		if grid[current.y][current.x] == "." {
			current = Point{current.x + direction.x, current.y + direction.y}
		}
	}

	fmt.Println("Sum of all GPS coords:", sumGPSCoords(grid))
}
