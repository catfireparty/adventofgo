package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day11"
)

func main() {
	path := "./day11/data.txt"
	start := time.Now()
	day11.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day11.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
