package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day9"
)

func main() {
	path := "./day9/data.txt"
	start := time.Now()
	day9.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day9.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
