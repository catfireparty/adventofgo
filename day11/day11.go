package day11

import (
	"fmt"
	"strconv"
	"strings"

	"adventofgo.dev/utils"
)

type Stone struct {
	number string
	next   *Stone
}

func PartOne(path string) {
	fmt.Println("Day 11 - Part 1: ")
	data := utils.ReadFile(path)
	line := strings.Split(data, "\n")[0]
	start, _ := toStones(strings.Split(line, " "))

	numBlinks := 25
	for range numBlinks {
		naiveBlink(start)
	}

	// fmt.Println(stonesToString(start))

	fmt.Println("Number of stones after", numBlinks, "blinks = ", countStones(start))
}

func countStones(start *Stone) int {
	numStones := 0
	current := start
	for {
		if current == nil {
			break
		}

		numStones++
		current = current.next
	}
	return numStones
}

func stonesToString(start *Stone) string {
	line := []string{}
	current := start
	for {
		if current == nil {
			break
		}
		line = append(line, current.number)
		current = current.next
	}
	return strings.Join(line, " ")
}

func toStones(line []string) (*Stone, *Stone) {
	var end *Stone
	var start *Stone
	var previous *Stone
	for i := 0; i < len(line); i++ {
		end = &Stone{line[i], nil}
		if i == 0 {
			start = end
		}
		if previous != nil {
			previous.next = end
		}
		previous = end
	}
	return start, end
}

func naiveBlink(start *Stone) {
	current := start
	for {
		if current == nil {
			break
		}

		if current.number == "0" {
			current.number = "1"
		} else if len(current.number)%2 == 0 {
			half := len(current.number) / 2

			// insert a new one
			next := current.next
			current.next = &Stone{strconv.Itoa(utils.ToInt(current.number[half:])), next}
			current.number = current.number[:half]
			current = current.next
		} else {
			current.number = strconv.Itoa(utils.ToInt(current.number) * 2024)
		}

		current = current.next
	}
}

func mapBlink(cache map[string]int) map[string]int {
	newCache := make(map[string]int)
	for number, count := range cache {
		if number == "0" {
			newCache["1"] += count
		} else if len(number)%2 == 0 {
			half := len(number) / 2
			newCache[number[:half]] += count
			newCache[strconv.Itoa(utils.ToInt(number[half:]))] += count
		} else {
			newCache[strconv.Itoa(utils.ToInt(number)*2024)] += count
		}
	}
	return newCache
}

func PartTwo(path string) {
	fmt.Println("Day 11 - Part 2: ")
	data := utils.ReadFile(path)
	line := strings.Split(data, "\n")[0]
	numbers := strings.Split(line, " ")
	cache := make(map[string]int)

	for i := range numbers {
		cache[numbers[i]] += 1
	}

	for range 75 {
		cache = mapBlink(cache)
	}

	numStones := 0
	for _, value := range cache {
		numStones += value
	}

	fmt.Println("Number of stones after 75 blinks = ", numStones)
}
