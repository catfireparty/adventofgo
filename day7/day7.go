package day7

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Calibration struct {
	result uint64
	values []uint64
}

type Operation uint64

var Multiply = uint64(1)
var Addition = uint64(2)
var Concat = uint64(3)

func PartOne(path string) {
	fmt.Println("Day 7 - Part 1: ")
	data := utils.ReadFile(path)
	calibrations := readCalibrations(data)

	totalCalibrationResult := uint64(0)
	for i := range calibrations {
		if isPossible(calibrations[i].result, calibrations[i].values, []Operation{Operation(Multiply), Operation(Addition)}) {
			totalCalibrationResult += calibrations[i].result
		}
	}

	fmt.Println("Total calibration result:", totalCalibrationResult)
}

type EquationStep struct {
	value uint64
	index int
}

// Use something like Dijkstra's algorithm to explore the equation space
func isPossible(expected uint64, values []uint64, operations []Operation) bool {
	frontier := []EquationStep{
		{values[0], 0},
	}

	for {
		current := frontier[0]
		frontier = frontier[1:]

		for i := range operations {
			nextValue := current.value
			nextIndex := current.index + 1
			switch operations[i] {
			case Operation(Multiply):
				nextValue = current.value * values[nextIndex]
			case Operation(Addition):
				nextValue = current.value + values[nextIndex]
			case Operation(Concat):
				nextValue = utils.ToUint64(fmt.Sprintf("%d%d", current.value, values[nextIndex]))
			}

			if nextIndex == (len(values) - 1) {
				if nextValue == expected {
					return true
				}
			} else {
				frontier = append(frontier, EquationStep{
					nextValue,
					nextIndex,
				})
			}
		}

		if len(frontier) == 0 {
			break
		}
	}

	return false
}

func evaluateExpression(expression []uint64, expected uint64) bool {
	result := expression[0]
	for i := 2; i < len(expression); i += 2 {
		operator := expression[i-1]
		if operator == Addition {
			result += expression[i]
		}
		if operator == Multiply {
			result *= expression[i]
		}
		if operator == Concat {
			result = utils.ToUint64(fmt.Sprintf("%d%d", result, expression[i]))
		}
		if result > expected {
			return false
		}
	}
	return result == expected
}

func generateExpressions(values []uint64, operators []uint64) [][]uint64 {
	expressions := [][]uint64{{values[0]}}

	for i := 1; i < len(values); i++ {
		updated := [][]uint64{}
		for j := range expressions {
			// for each expression, create another for each operator
			value := values[i]
			length := len(expressions[j]) + 2

			for k := range operators {
				operation := make([]uint64, length)

				copy(operation, expressions[j])
				operation[length-2] = operators[k]
				operation[length-1] = value

				updated = append(updated, operation)
			}
		}
		expressions = updated
	}

	return expressions
}

func readCalibrations(data string) []Calibration {
	calibration := []Calibration{}
	lines := strings.Split(data, "\n")
	for i := range lines {
		line := lines[i]
		if len(line) < 1 {
			continue
		}

		parts := strings.Split(line, ": ")

		calibration = append(calibration, Calibration{
			utils.ToUint64(parts[0]),
			utils.ToUint64Array(strings.Split(parts[1], " ")),
		})
	}
	return calibration
}

func PartTwo(path string) {
	fmt.Println("Day 7 - Part 2: ")
	data := utils.ReadFile(path)
	calibrations := readCalibrations(data)

	totalCalibrationResult := uint64(0)
	for i := range calibrations {
		if isPossible(calibrations[i].result, calibrations[i].values, []Operation{Operation(Multiply), Operation(Addition), Operation(Concat)}) {
			totalCalibrationResult += calibrations[i].result
		}
	}

	fmt.Println("Total calibration result:", totalCalibrationResult)
}
