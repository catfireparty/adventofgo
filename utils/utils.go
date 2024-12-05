package utils

import (
	"os"
	"strconv"
)

func ReadFile(path string) string {
	data, error := os.ReadFile(path)

	if error != nil {
		panic(error)
	}

	return string(data)
}

func ToInt(input string) int {
	value, error := strconv.Atoi(input)
	if error != nil {
		panic(error)
	}
	return value
}

func ToIntArray(input []string) []int {
	array := []int{}
	for i := range input {
		array = append(array, ToInt(input[i]))
	}
	return array
}
