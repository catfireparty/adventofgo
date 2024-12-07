package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day7"
)

func main() {
	path := "./day7/data.txt"
	start := time.Now()
	day7.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day7.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
