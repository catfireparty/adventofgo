package day5

import (
	"fmt"
	"strings"

	"adventofgo.dev/utils"
)

type Rule struct {
	first  int
	second int
}

func PartOne(path string) {
	fmt.Println("Day 5 - Part 1: ")
	data := utils.ReadFile(path)
	lines := strings.Split(data, "\n")
	rules, updates := getRulesAndUpdates(lines)

	middleSum := 0

	// check validity of each update
	for i := range updates {
		conforms := true
		for j := range rules {
			conforms = conformsToRule(updates[i], rules[j])
			if !conforms {
				break
			}
		}
		if conforms {
			// get the middle number
			middleSum += updates[i][len(updates[i])/2]
		}
	}

	fmt.Println("Middle sum := ", middleSum)
}

func conformsToRule(update []int, rule Rule) bool {
	found := make(map[int]bool)
	for i := range update {
		if update[i] == rule.first && found[rule.second] {
			return false
		}
		found[update[i]] = true
	}
	return true
}

func getRulesAndUpdates(lines []string) ([]Rule, [][]int) {
	rules := []Rule{}
	updates := [][]int{}

	for i := range lines {
		line := lines[i]

		if len(line) < 1 {
			continue
		}

		if strings.Contains(line, "|") {
			// parse rule
			pages := strings.Split(line, "|")
			rules = append(rules, Rule{
				first:  utils.ToInt(pages[0]),
				second: utils.ToInt(pages[1]),
			})
		}

		if strings.Contains(line, ",") {
			// parse pages
			updates = append(updates, utils.ToIntArray(strings.Split(line, ",")))
		}
	}

	return rules, updates
}

func PartTwo(path string) {
	fmt.Println("Day 5 - Part 2: ")
	data := utils.ReadFile(path)
	lines := strings.Split(data, "\n")
	rules, updates := getRulesAndUpdates(lines)

	badUpdates := [][]int{}

	// find the incorrect updates
	for i := range updates {
		for j := range rules {
			if !conformsToRule(updates[i], rules[j]) {
				badUpdates = append(badUpdates, updates[i])
				break
			}
		}
	}

	middleSum := 0

	// for every non-conforming update, lets go through
	// and when we hit a non-conformance, move the current
	// number to the front of the update list
	for i := range badUpdates {
		for {
			changed := false
			update := badUpdates[i]

			// check each of the rules against this update
			for j := range rules {
				rule := rules[j]
				found := make(map[int]bool)

				for k := range update {

					if update[k] == rule.first && found[rule.second] {
						// the update does not conform, so...
						// move the non-conforming number to the
						// front of the queue and start over
						badUpdates[i] = moveToFront(update, k)
						changed = true
						break
					}

					if changed {
						break
					}

					found[update[k]] = true
				}

				if changed {
					break
				}
			}

			if !changed {
				break
			}
		}
		middleSum += badUpdates[i][len(badUpdates[i])/2]
	}

	fmt.Println("Middle sum := ", middleSum)
}

func moveToFront(array []int, indexToMove int) []int {
	newArray := []int{array[indexToMove]}
	for i := range array {
		if i == indexToMove {
			continue
		}
		newArray = append(newArray, array[i])
	}
	return newArray
}
