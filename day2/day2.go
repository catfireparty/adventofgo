package day2

import (
	"fmt"
	"strconv"
	"strings"

	"adventofgo.dev/utils"
)

func PartOne(path string) {
	fmt.Println("Day 2 - Part 1: ")
	data := utils.ReadFile(path)

	lines := strings.Split(data, "\n")
	safeCount := 0
	for _, line := range lines {
		report := strings.Split(line, " ")
		if len(report) <= 1 {
			continue
		}

		if isSafe(report, -1) {
			safeCount++
		}
	}

	fmt.Println("Safe reports: ", safeCount)
}

func toInt(input string) int {
	value, error := strconv.Atoi(input)
	if error != nil {
		panic(error)
	}
	return value
}

func isSafe(report []string, toExclude int) bool {
	isSafe := true

	firstIndex := 0
	if toExclude == 0 {
		firstIndex = 1
	}

	startIndex := firstIndex + 1
	if toExclude == startIndex {
		startIndex++
	}

	lastIndex := len(report) - 1
	if toExclude == lastIndex {
		lastIndex--
	}

	previous := toInt(report[firstIndex])
	decreasing := previous > toInt(report[lastIndex])
	for i := startIndex; i < len(report); i++ {
		if i == toExclude {
			continue
		}

		current := utils.ToInt(report[i])

		if decreasing {
			if current >= previous {
				isSafe = false
				break
			}

			if previous-current > 3 {
				isSafe = false
				break
			}
		} else {
			if current <= previous {
				isSafe = false
				break
			}

			if current-previous > 3 {
				isSafe = false
				break
			}
		}

		previous = current
	}

	return isSafe
}

func PartTwo(path string) {
	fmt.Println("Day 2 - Part 2: ")
	data := utils.ReadFile(path)
	lines := strings.Split(data, "\n")
	safeCount := 0
	for _, line := range lines {
		report := strings.Split(line, " ")
		if len(report) <= 1 {
			continue
		}

		if isSafe(report, -1) {
			safeCount++
		} else {
			// use problem dampener
			for i := range report {
				if isSafe(report, i) {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println("Safe reports: ", safeCount)
}
