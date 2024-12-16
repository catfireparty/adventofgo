package day14

import (
	"fmt"
	"regexp"
	"strings"

	"adventofgo.dev/utils"
)

type Grid = [103][101]int

type Point struct {
	x int
	y int
}

type Speed = Point

type Robot struct {
	current Point
	speed   Speed
}

func PartOne(data string) {
	fmt.Println("Day 14 - Part 1: ")
	robots := dataToRobots(data)
	grid := Grid{}
	xMax := len(grid[0])
	yMax := len(grid)

	for range 100 {
		// fmt.Println(robots)
		// tmpGrid := Grid{}
		// for i := range robots {
		// 	robot := robots[i]
		// 	tmpGrid[robot.current.y][robot.current.x] += 1
		// }

		// for i := range tmpGrid {
		// 	fmt.Println(utils.ToString(tmpGrid[i][:]))
		// }

		// fmt.Println("after", seconds, "seconds")
		for i := range robots {
			robot := &robots[i]
			robot.current.x = robot.current.x + robot.speed.x
			robot.current.y = robot.current.y + robot.speed.y
			if robot.current.x >= xMax {
				robot.current.x -= xMax
			}
			if robot.current.x < 0 {
				robot.current.x += xMax
			}
			if robot.current.y >= yMax {
				robot.current.y -= yMax
			}
			if robot.current.y < 0 {
				robot.current.y += yMax
			}
			// fmt.Println(robot)
		}
	}

	// for i := range grid {
	// 	fmt.Println(utils.ToString(grid[i][:]))
	// }

	sf, _ := safetyFactor(robots)

	fmt.Println("Total safety factor:", sf)
}

func printGrid(robots []Robot) []string {
	grid := Grid{}
	lines := []string{}
	for i := range robots {
		robot := robots[i]
		grid[robot.current.y][robot.current.x] += 1
	}
	for i := range grid {
		lines = append(lines, toString(grid[i][:]))
	}
	return lines
}

func safetyFactor(robots []Robot) (int, Grid) {
	grid := Grid{}
	xMax := len(grid[0])
	yMax := len(grid)

	for i := range robots {
		robot := robots[i]
		grid[robot.current.y][robot.current.x] += 1
	}

	// count the robots in each quadrant
	midX := (xMax - 1) / 2
	midY := (yMax - 1) / 2

	nw := 0
	for x := range midX {
		for y := range midY {
			nw += grid[y][x]
		}
	}

	ne := 0
	for x := midX + 1; x < xMax; x++ {
		for y := range midY {
			ne += grid[y][x]
		}
	}

	sw := 0
	for x := range midX {
		for y := midY + 1; y < yMax; y++ {
			sw += grid[y][x]
		}
	}

	se := 0
	for x := midX + 1; x < xMax; x++ {
		for y := midY + 1; y < yMax; y++ {
			se += grid[y][x]
		}
	}

	return nw * ne * sw * se, grid
}

func toString(input []int) string {
	line := []string{}
	for i := range input {
		if input[i] == 0 {
			line = append(line, ".")
		} else {
			line = append(line, "X")
		}
	}
	return strings.Join(line, "")
}

func dataToRobots(data string) []Robot {
	lines := strings.Split(data, "\n")
	robots := []Robot{}
	for i := range lines {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		start := utils.ToIntArray(strings.Split(strings.Split(parts[0], "=")[1], ","))
		speed := utils.ToIntArray(strings.Split(strings.Split(parts[1], "=")[1], ","))

		robots = append(robots, Robot{
			current: Point{start[0], start[1]},
			speed:   Speed{speed[0], speed[1]},
		})
	}
	return robots
}

func PartTwo(data string) {
	fmt.Println("Day 14 - Part 2: ")
	robots := dataToRobots(data)
	grid := Grid{}
	xMax := len(grid[0])
	yMax := len(grid)

	matcher := regexp.MustCompile(`X{16}`)

	found := false
	for i := range 10000 {
		for r := range robots {
			robot := &robots[r]
			robot.current.x = robot.current.x + robot.speed.x
			robot.current.y = robot.current.y + robot.speed.y
			if robot.current.x >= xMax {
				robot.current.x -= xMax
			}
			if robot.current.x < 0 {
				robot.current.x += xMax
			}
			if robot.current.y >= yMax {
				robot.current.y -= yMax
			}
			if robot.current.y < 0 {
				robot.current.y += yMax
			}
			// fmt.Println(robot)
		}

		lines := printGrid(robots)
		for j := range lines {
			if matcher.MatchString(lines[j]) {
				fmt.Println(i + 1)
				fmt.Println(strings.Join(printGrid(robots), "\n"))
				found = true
				break
			}
		}

		if found {
			break
		}
	}

}
