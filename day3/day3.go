package day3

import (
	"fmt"
	"regexp"
	"strings"

	"adventofgo.dev/utils"
)

func PartOne(path string) {
	fmt.Println("Day 3 - Part 1: ")
	data := utils.ReadFile(path)

	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	ops := matcher.FindAllString(data, -1)

	total := int(0)
	for i := range ops {
		total += mul(ops[i])
	}

	fmt.Println("Total: ", total)
}

func mul(op string) int {
	operands := strings.Split(op[4:len(op)-1], ",")
	first := utils.ToInt(operands[0])
	second := utils.ToInt(operands[1])
	return first * second
}

func PartTwo(path string) {
	fmt.Println("Day 3 - Part 2: ")
	data := utils.ReadFile(path)

	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	ops := matcher.FindAllString(data, -1)

	total := int(0)
	shouldMultiply := true
	for i := range ops {
		op := strings.Split(ops[i], "(")[0]

		switch op {
		case "mul":
			if shouldMultiply {
				total += mul(ops[i])
			}
		case "do":
			shouldMultiply = true
		case "don't":
			shouldMultiply = false
		}
	}

	fmt.Println("Total: ", total)
}
