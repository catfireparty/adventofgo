package main

import (
	"fmt"

	"adventofgo.dev/day12"
)

func main() {
	path := "./day12/data.txt"
	elapsed := day12.PartOne(path)
	fmt.Printf("Part 1 took %s \n", elapsed)

	elapsed = day12.PartTwo(path)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
