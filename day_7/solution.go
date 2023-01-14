package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const DISK_SPACE = 70_000_000
const UNUSED_SPACE = 30_000_000

type DirTree struct {
	root *Dir
}

type Dir struct {
	name     string
	files    []File
	parent   *Dir
	children []*Dir
}

type File struct {
	name string
	size int
}

func main() {
	inputLines := utils.ReadInputAsStrings("input.txt")

	solution1 := part1(inputLines)
	solution2 := part2(inputLines)

	utils.PrintSolution(&[]string{
		fmt.Sprintf("part1: %v", solution1),
		fmt.Sprintf("part2: %v", solution2),
	})
}

func part1(inputLines []string) int {
	dirTree := parseInput(inputLines)

	_, sumDirSizes := calculateDirSizes(*dirTree.root)

	return sumDirSizes
}

func part2(inputLines []string) int {
	dirTree := parseInput(inputLines)

	// Reusing part 1 function to get the total space taken by all dirs, aka
	// the size of the root folder.
	totalOccupiedSpace, _ := calculateDirSizes(*dirTree.root)
	idealSize := UNUSED_SPACE - (DISK_SPACE - totalOccupiedSpace)

	_, targetDirSize := getDirToDelete(*dirTree.root, idealSize)

	return targetDirSize
}

func getDirToDelete(currentDir Dir, idealSize int) (int, int) {
	totalDirSize := 0
	// The max size a single dir to get deleted can be is the total disk space
	// math.MaxInt would also work
	closestToIdeal := DISK_SPACE

	for _, file := range currentDir.files {
		totalDirSize += file.size
	}

	if currentDir.children != nil {
		for _, child := range currentDir.children {
			childDirSize, childClosestToIdeal := getDirToDelete(*child, idealSize)
			totalDirSize += childDirSize

			if childClosestToIdeal < closestToIdeal {
				closestToIdeal = childClosestToIdeal
			}
		}
	}

	// Check if current dir is closest to the ideal size:
	// The dir must be larger than the ideal size to free up enough space
	// The dir size must be smaller than the previous dir closest to ideal size
	if totalDirSize > idealSize && totalDirSize < closestToIdeal {
		closestToIdeal = totalDirSize
	}

	return totalDirSize, closestToIdeal
}

func calculateDirSizes(currentDir Dir) (int, int) {
	var totalDirSize int
	var totalSizeSmallDirs int

	for _, file := range currentDir.files {
		totalDirSize += file.size
	}

	if currentDir.children != nil {
		for _, child := range currentDir.children {
			childDirSize, smallDirsSize := calculateDirSizes(*child)
			totalDirSize += childDirSize
			totalSizeSmallDirs += smallDirsSize
		}
	}

	if totalDirSize <= 100_000 {
		totalSizeSmallDirs += totalDirSize
	}

	return totalDirSize, totalSizeSmallDirs
}

func parseInput(lines []string) DirTree {
	var dirTree DirTree
	var currentDir *Dir

	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		if splitLine[0] == "$" {
			command := splitLine[1]

			if command == "cd" {
				path := splitLine[2]
				currentDir = currentDir.traverseTree(path, &dirTree)
			}

			if command == "ls" {
				continue
			}

		} else {
			currentDir.addItems(splitLine)
		}

	}

	return dirTree
}

func (currentDir *Dir) traverseTree(path string, dirTree *DirTree) *Dir {
	if path == "/" {
		currentDir = &Dir{name: "root", parent: nil, files: []File{}}
		dirTree.root = currentDir
	} else if path == ".." {
		currentDir = currentDir.parent
	} else {
		for _, child := range currentDir.children {
			if child.name == path {
				currentDir = child
			}
		}
	}

	return currentDir
}

func (dir *Dir) addItems(item []string) {
	if item[0] == "dir" {
		isNewDir := true

		// In case a combination of commands revisits the current dir,
		// If there has been a ls command performed for this dir before,
		// it's children should contain the listed dir on the current line
		// and should not be added again
		if dir.children != nil {
			for _, child := range dir.children {
				if child.name == item[1] {
					isNewDir = false
				}
			}
		}

		if isNewDir {
			newDir := Dir{name: item[1], parent: dir}
			dir.children = append(dir.children, &newDir)
		}

	} else {
		size, _ := strconv.Atoi(item[0])

		newFile := File{name: item[1], size: size}
		dir.files = append(dir.files, newFile)
	}
}
