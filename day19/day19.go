package day19

import (
	"fmt"
	"regexp"
	"strings"
)

func PartOne(data string) {
	fmt.Println("Day 19 - Part 1: ")
	towels, designs := parseData(data)

	possible := getPossibleDesigns(towels, designs)

	fmt.Println("Possible designs:", len(possible))
}

func parseData(data string) ([]string, []string) {
	lines := strings.Split(data, "\n")

	towels := strings.Split(lines[0], ", ")
	designs := []string{}

	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		designs = append(designs, lines[i])
	}

	return towels, designs
}

func getPossibleDesigns(towels []string, designs []string) []string {
	matcher := regexp.MustCompile(`^(` + strings.Join(towels, "|") + `)+$`)
	possible := []string{}
	for i := range designs {
		if matcher.MatchString(designs[i]) {
			possible = append(possible, designs[i])
		}
	}
	return possible
}

func findWaysToMake(design string, towels []string, cache map[string]int) int {
	if cache[design] > 0 {
		return cache[design]
	}

	count := 0
	for i := range towels {
		towel := towels[i]
		if design == towel {
			count++
		} else if strings.HasPrefix(design, towel) {
			count += findWaysToMake(strings.TrimPrefix(design, towel), towels, cache)
		}
	}

	cache[design] = count
	return count
}

func PartTwo(data string) {
	fmt.Println("Day 19 - Part 2: ")
	towels, designs := parseData(data)
	possible := getPossibleDesigns(towels, designs)

	// for each possible design, figure out how many ways there are to build it
	count := 0
	cache := make(map[string]int)
	for i := range possible {
		count += findWaysToMake(possible[i], towels, cache)
	}

	fmt.Println("Possible combinations:", count)
}
