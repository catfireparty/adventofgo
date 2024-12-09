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
	id     int
	length int
}

func expandFileSystemToFiles(line []string) []File {
	expanded := []File{}
	for i := range line {
		// even numbers are file blocks
		// odd numbers are free space
		length := utils.ToInt(line[i])
		if i%2 == 0 {
			expanded = append(expanded, File{(i + 1) / 2, length})
		} else {
			expanded = append(expanded, File{-1, length})
		}
	}
	return expanded
}

func compactFileSystemByFile(files []File) []File {
	moved := make(map[int]bool)
	fileIndex := len(files) - 1
	for {
		fileId := files[fileIndex].id
		if moved[fileId] {
			break
		}
		files = moveFileIfPossible(files, fileIndex)
		// fmt.Println(expandFileSystemTwo(files))
		moved[fileId] = true

		// get next file to move
		for i := len(files) - 1; i > 0; i-- {
			file := files[i]
			if file.id == -1 || moved[file.id] {
				continue
			}
			fileIndex = i
			break
		}
	}
	return files
}

func moveFileIfPossible(files []File, index int) []File {
	fileToMove := files[index]
	newFileSystem := []File{}
	moved := false
	for i := range files {
		current := files[i]
		if !moved && current.id == -1 && current.length >= fileToMove.length && i < index {
			newFileSystem = append(newFileSystem, File{fileToMove.id, fileToMove.length})
			if fileToMove.length < current.length {
				newFileSystem = append(newFileSystem, File{-1, current.length - fileToMove.length})
			}
			moved = true
		} else if current.id == fileToMove.id && moved {
			newFileSystem = append(newFileSystem, File{-1, fileToMove.length})
		} else {
			newFileSystem = append(newFileSystem, File{current.id, current.length})
		}
	}
	return newFileSystem
}

func expandFileSystemTwo(files []File) []string {
	expanded := []string{}
	for i := range files {
		file := files[i]
		content := strconv.Itoa(file.id)
		if file.id == -1 {
			content = "."
		}
		for range file.length {
			expanded = append(expanded, content)
		}
	}
	return expanded
}

func PartTwo(path string) {
	fmt.Println("Day 9 - Part 2: ")
	data := utils.ReadFile(path)
	line := strings.Split(strings.Split(data, "\n")[0], "")
	expanded := expandFileSystemToFiles(line)
	// fmt.Println(expandFileSystemTwo(expanded))
	compacted := compactFileSystemByFile(expanded)

	checksum := calculateCheckSum(expandFileSystemTwo(compacted))
	// fmt.Println(expandFileSystemTwo(compacted))

	fmt.Println("Checksum:", checksum)
}
