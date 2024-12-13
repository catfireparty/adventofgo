package day13

import (
	"fmt"
	"regexp"
	"strings"

	"adventofgo.dev/utils"
)

type Vector struct {
	x int
	y int
}

type Machine struct {
	A     Vector
	B     Vector
	prize Vector
}

func PartOne(path string) {
	fmt.Println("Day 13 - Part 1: ")
	data := utils.ReadFile(path)
	machines := parseToMachines(data)

	tokens := 0
	for i := range machines {
		optimal := solveForMachine(machines[i])
		tokens += optimal
	}

	fmt.Println("Fewest tokens possible:", tokens)
}

func solveForMachine(machine Machine) int {
	mintokens := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			dx := a*machine.A.x + b*machine.B.x
			dy := a*machine.A.y + b*machine.B.y
			cost := a*3 + b
			if dx == machine.prize.x && dy == machine.prize.y {
				// possible
				if mintokens == 0 || cost < mintokens {
					mintokens = cost
				}
			}
		}
	}
	return mintokens
}

func parseToMachines(data string) []Machine {
	lines := strings.Split(data, "\n")
	machines := []Machine{}
	for i := 0; i < len(lines); i += 4 {

		machines = append(machines, Machine{
			A:     getVector(lines[i]),
			B:     getVector(lines[i+1]),
			prize: getVector(lines[i+2]),
		})
	}
	return machines
}

func getVector(line string) Vector {
	matcher := regexp.MustCompile(`\d+`)
	numbers := matcher.FindAllString(line, -1)
	return Vector{
		x: utils.ToInt(numbers[0]),
		y: utils.ToInt(numbers[1]),
	}
}

func solveEquationForMachine(machine Machine) int {
	// use cramers rule
	// https://en.wikipedia.org/wiki/Cramer's_rule

	a1 := machine.A.x
	b1 := machine.B.x
	c1 := machine.prize.x + 10000000000000

	a2 := machine.A.y
	b2 := machine.B.y
	c2 := machine.prize.y + 10000000000000

	// i.e. where we're trying to solve for
	// x - number of A presses
	// y - number of B presses
	// a1x + b1y = c1
	// a2x + b2y = c2

	D := a1*b2 - b1*a2
	if D == 0 {
		// no solution
		return 0
	}

	xD := (c1*b2 - b1*c2)
	yD := (a1*c2 - c1*a2)

	x := xD / D
	y := yD / D

	if xD%D != 0 || yD%D != 0 || x < 0 || y < 0 {
		return 0
	}

	return x*3 + y
}

func PartTwo(path string) {
	fmt.Println("Day 13 - Part 2: ")
	data := utils.ReadFile(path)
	machines := parseToMachines(data)

	tokens := 0
	for i := range machines {
		optimal := solveEquationForMachine(machines[i])
		tokens += optimal
	}

	fmt.Println("Fewest tokens possible:", tokens)
}
