package day1

import (
	"fmt"
	"sort"
	"strings"

	"adventofgo.dev/utils"
)

func PartOne(path string) {
	fmt.Println("Day 1 - Part 1: ")
	data := utils.ReadFile(path)
	list1, list2 := convertToSlices(data)
	sort.Ints(list1)
	sort.Ints(list2)

	total := int(0)
	for i := 0; i < len(list1); i++ {
		difference := list1[i] - list2[i]
		if difference < 0 {
			total -= difference
		} else {
			total += difference
		}
	}

	fmt.Println(total)
}

func PartTwo(path string) {
	fmt.Println("Day 1 - Part 2: ")
	data := utils.ReadFile(path)
	list1, list2 := convertToSlices(data)
	sort.Ints(list1)
	sort.Ints(list2)

	countMap := make(map[int]int)

	for i := range list2 {
		countMap[list2[i]] += 1
	}

	score := 0
	for i := range list1 {
		score += list1[i] * countMap[list1[i]]
	}

	fmt.Println(score)
}

func convertToSlices(data string) ([]int, []int) {
	lines := strings.Split(string(data), "\n")
	list1 := []int{}
	list2 := []int{}
	for _, line := range lines {
		split := strings.Split(line, "   ")
		if len(split) < 2 {
			continue
		}

		first := utils.ToInt(split[0])
		second := utils.ToInt(split[1])

		list1 = append(list1, int(first))
		list2 = append(list2, int(second))
	}

	return list1, list2
}
