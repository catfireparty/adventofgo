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

func ToUint64(input string) uint64 {
	var result, error = strconv.ParseUint(input, 10, 64)
	if error != nil {
		panic(input)
	}
	return result
}

func ToUint64Array(input []string) []uint64 {
	array := []uint64{}
	for i := range input {
		array = append(array, ToUint64(input[i]))
	}
	return array
}

func ToIntArray(input []string) []int {
	array := []int{}
	for i := range input {
		array = append(array, ToInt(input[i]))
	}
	return array
}
