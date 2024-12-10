package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day10"
)

func main() {
	path := "./day10/data.txt"
	start := time.Now()
	day10.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day10.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
