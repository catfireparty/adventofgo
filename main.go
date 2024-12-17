package main

import (
	"fmt"
	"time"

	"adventofgo.dev/day16"
	"adventofgo.dev/utils"
)

func main() {
	data := utils.ReadFile("./day16/data.txt")

	start := time.Now()
	day16.PartOne(data)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s \n", elapsed)

	start = time.Now()
	day16.PartTwo(data)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s \n", elapsed)
}
