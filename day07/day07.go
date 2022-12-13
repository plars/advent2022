package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent2022/day07/filesystem"
)

func main() {
	root := processInput("input.txt")
	// print the tree for sanity checking
	printTree(root)
	partA := sumSizeUnder(root, 100000)
	fmt.Println(partA)
	/*
		we need 30000000 bytes free to install the update, and the total
		capacity is 70000000 so to calculate how much we need to delete, we
		need to subtract (70000000-30000000) from the total space in use
	*/

	minDeleteSize := root.Size - 40000000
	partB := findSmallestDirOver(root, minDeleteSize)
	fmt.Println(partB)
}

func readFile(filename string) []string {
	//read filename and return a slice of strings
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func processInput(filename string) *filesystem.Inode {
	root := filesystem.NewInode("/", nil, 0, true)
	var cwd *filesystem.Inode
	data := readFile(filename)
	for _, line := range data {
		if line == "" {
			// skip empty lines
			continue
		}
		if line[0] == '$' {
			//command
			if line[2:4] == "cd" {
				if line[5] == '/' {
					cwd = root
					continue
				}
				if line[5:] == ".." {
					cwd = cwd.Parent
					continue
				}
				for _, inode := range cwd.Children {
					if inode.Name == line[5:] {
						cwd = inode
						break
					}
				}
				continue
			}
			if line[2:4] == "ls" {
				// keep reading lines until we get a $
				// dirs and files are handled separately, no need to process them here
				continue
			}
		}
		if line[0:3] == "dir" {
			// create new directory
			newDir := filesystem.NewInode(line[4:], cwd, 0, true)
			cwd.AddChild(newDir)
			continue
		}
		// otherwise this should be a file in the format "size filename"
		words := strings.Split(line, " ")
		size, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatal(err)
		}
		newFile := filesystem.NewInode(words[1], cwd, size, false)
		cwd.AddChild(newFile)
	}
	return root
}

func printTree(dir *filesystem.Inode) {
	// print the tree
	fmt.Println(dir.Name, dir.Size)
	for _, link := range dir.Children {
		if link.IsDir {
			printTree(link)
		}
	}
}

func sumSizeUnder(dir *filesystem.Inode, maxSize int) int {
	// sum the size of all files under maxSize bytes
	var sum int
	for _, link := range dir.Children {
		if link.IsDir {
			sum += sumSizeUnder(link, maxSize)
			if link.Size <= maxSize {
				sum += link.Size
			}
		}
	}
	return sum
}

func findSmallestDirOver(dir *filesystem.Inode, minSize int) int {
	// find the smallest directory that is over minSize bytes
	// start with current directory size since we know it should be ok
	smallest := dir.Size
	for _, link := range dir.Children {
		if link.IsDir {
			if link.Size >= minSize {
				tmp := findSmallestDirOver(link, minSize)
				if tmp < smallest {
					smallest = tmp
				}
			}
		}
	}
	return smallest
}
