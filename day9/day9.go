package day9

import (
	"fmt"
	"strconv"
	"strings"

	"adventofgo.dev/utils"
)

func PartOne(path string) {
	fmt.Println("Day 9 - Part 1: ")
	data := utils.ReadFile(path)
	line := strings.Split(strings.Split(data, "\n")[0], "")
	expanded := expandFileSystem(line)
	compactFileSystemByBlock(expanded)
	checksum := calculateCheckSum(expanded)
	// fmt.Println(expanded)

	fmt.Println("Checksum:", checksum)
}

func expandFileSystem(line []string) []string {
	expanded := []string{}
	currentId := 0
	for i := range line {
		// even numbers are file blocks
		// odd numbers are free space
		number := utils.ToInt(line[i])
		content := "."
		if i%2 == 0 {
			content = strconv.Itoa(currentId)
			currentId++
		}
		for range number {
			expanded = append(expanded, content)
		}
	}
	return expanded
}

func compactFileSystemByBlock(line []string) {
	x := 0
	for i := len(line) - 1; i > x; i-- {
		if line[i] == "." {
			continue
		}

		for ; x < i; x++ {
			if line[x] == "." {
				line[x] = line[i]
				line[i] = "."
				break
			}
		}
	}
}

func calculateCheckSum(line []string) int {
	checksum := 0
	for i := range line {
		if line[i] == "." {
			continue
		}
		checksum += utils.ToInt(line[i]) * i
	}
	return checksum
}

type File struct {
	id       int
	length   int
	previous *File
	next     *File
}

func expandFileSystemToFiles(line []string) *File {
	var current *File = nil
	var previous *File = nil
	for i := 0; i < len(line); i++ {
		// even numbers are file blocks
		// odd numbers are free space
		length := utils.ToInt(line[i])
		if i%2 == 0 {
			current = &File{(i + 1) / 2, length, previous, nil}
		} else {
			current = &File{-1, length, previous, nil}
		}
		if previous != nil {
			previous.next = current
		}
		previous = current
	}
	return current
}

func compactFileSystemByFile(file *File) {
	moved := make(map[int]bool)

	// find the start
	start := beginningOfList(file)
	toMove := endOfList(file)

	for {
		current := start
		for {
			if current.id == toMove.id {
				moved[toMove.id] = true
				break
			}

			if current.id == -1 && current.length >= toMove.length {
				// current File is free space with enough to fit toMove

				// unpick toMove from it's current position
				// and replace with empty space
				next := toMove.next
				prev := toMove.previous
				space := &File{-1, toMove.length, nil, nil}
				if prev != nil {
					prev.next = space
					space.previous = prev
				}
				if next != nil {
					next.previous = space
					space.next = next
				}
				toMove.next = nil
				toMove.previous = nil

				current.previous.next = toMove
				toMove.previous = current.previous

				if current.length > toMove.length {
					toMove.next = &File{-1, current.length - toMove.length, toMove, current.next}
					current.next.previous = toMove.next
				} else {
					toMove.next = current.next
					current.next.previous = toMove
				}

				// orphan current
				current.next = nil
				current.previous = nil
				moved[toMove.id] = true
				break
			}

			current = current.next
		}

		toMove = endOfList(toMove)
		// fmt.Println(expandFileSystemTwo(toMove))

		for {
			if toMove.id != -1 && !moved[toMove.id] {
				break
			}
			toMove = toMove.previous
			if toMove == nil {
				break
			}
		}

		if toMove == nil {
			break
		}
	}
}

func beginningOfList(file *File) *File {
	current := file
	for {
		if current.previous == nil {
			break
		}
		current = current.previous
	}
	return current
}

func endOfList(file *File) *File {
	current := file
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	return current
}

func expandFileSystemTwo(file *File) []string {
	expanded := []string{}

	// find the beginning of the list
	current := beginningOfList(file)

	for {
		if current == nil {
			break
		}
		content := strconv.Itoa(current.id)
		if current.id == -1 {
			content = "."
		}
		for range current.length {
			expanded = append(expanded, content)
		}
		current = current.next
	}

	return expanded
}

func PartTwo(path string) {
	fmt.Println("Day 9 - Part 2: ")
	data := utils.ReadFile(path)
	line := strings.Split(strings.Split(data, "\n")[0], "")
	file := expandFileSystemToFiles(line)
	// fmt.Println(expandFileSystemTwo(file))
	compactFileSystemByFile(file)
	// fmt.Println(expandFileSystemTwo(file))

	checksum := calculateCheckSum(expandFileSystemTwo(file))

	fmt.Println("Checksum:", checksum)
}
