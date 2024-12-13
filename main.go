package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day13"
)

func main() {
	path := "./day13/data.txt"
	start := time.Now()
	day13.PartOne(path)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day13.PartTwo(path)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
