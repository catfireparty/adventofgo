package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day19"
	"adventofgo.dev/utils"
)

func main() {
	data := utils.ReadFile("./day19/data.txt")

	start := time.Now()
	day19.PartOne(data)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day19.PartTwo(data)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
