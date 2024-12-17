package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day15"
	"adventofgo.dev/utils"
)

func main() {
	data := utils.ReadFile("./day15/data.txt")

	start := time.Now()
	// day15.PartOne(data)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day15.PartTwo(data)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
