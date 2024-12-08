package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day8"
)

func main() {
	path := "./day8/data.txt"
	start := time.Now()
	day8.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day8.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
