package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"adventofgo.dev/utils"
)

type Operation int

const (
	adv Operation = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type Machine struct {
	A       int
	B       int
	C       int
	program []int
	pointer int
}

func PartOne(data string) {
	fmt.Println("Day 17 - Part 1: ")

	machine, _ := parseData(data)

	output := runMachine(machine)

	fmt.Println(machine)
	fmt.Println("Output is:", strings.Join(output, ","))
}

func runMachine(machine Machine) []string {
	output := []string{}
runloop:
	for machine.pointer < len(machine.program) {
		operation := Operation(machine.program[machine.pointer])
		operand := machine.program[machine.pointer+1]
		switch operation {
		case adv:
			machine.A = divideA(&machine, operand)

		case bxl:
			machine.B = machine.B ^ operand

		case bst:
			comboValue := getComboOperandValue(&machine, operand)
			machine.B = comboValue % 8

		case jnz:
			if machine.A != 0 {
				machine.pointer = operand
				continue runloop
			}

		case bxc:
			machine.B = machine.B ^ machine.C

		case out:
			comboValue := getComboOperandValue(&machine, operand)
			output = append(output, strconv.Itoa(comboValue%8))

		case bdv:
			machine.B = divideA(&machine, operand)

		case cdv:
			machine.C = divideA(&machine, operand)

		}

		machine.pointer += 2
	}

	return output
}

func divideA(machine *Machine, operand int) int {
	numerator := machine.A
	denominator := math.Pow(2, float64(getComboOperandValue(machine, operand)))
	return int(float64(numerator) / denominator)
}

func getComboOperandValue(machine *Machine, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return machine.A
	case 5:
		return machine.B
	case 6:
		return machine.C
	case 7:
		panic("Unexpected operand: 7")
	default:
		panic(fmt.Sprintf("Unexpected operand: %d", operand))
	}
}

func parseData(data string) (Machine, []string) {
	lines := strings.Split(data, "\n")

	A := utils.ToInt(strings.Split(lines[0], ": ")[1])
	B := utils.ToInt(strings.Split(lines[1], ": ")[1])
	C := utils.ToInt(strings.Split(lines[2], ": ")[1])

	strProgram := strings.Split(lines[4], ": ")[1]
	program := utils.ToIntArray(strings.Split(strProgram, ","))

	return Machine{
		A,
		B,
		C,
		program,
		0,
	}, strings.Split(strProgram, ",")
}

func PartTwo(data string) {
	fmt.Println("Day 17 - Part 2: ")
	machine, program := parseData(data)

	// Found by other means, just checking it here
	// TODO : come back to this with a proper solution
	// 			  it's something to do with octal and << 3
	value := 105734774294938
	machine.A = value
	output := runMachine(machine)
	if strings.Join(output, ",") == strings.Join(program, ",") {
		fmt.Println("Success:", value)
	}
}
