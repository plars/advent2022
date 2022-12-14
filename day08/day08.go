package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeMap struct {
	MaxX int
	MaxY int
	Map  [][]int
}

func main() {
	fmt.Println("Part A:")
	solutionA := partA("input.txt")
	fmt.Println(solutionA)
	fmt.Println("Part B:")
	solutionB := partB("input.txt")
	fmt.Println(solutionB)
}

func partA(filename string) int {
	trees := readInput(filename)
	var numVisible int
	for y := 0; y <= trees.MaxY; y++ {
		for x := 0; x <= trees.MaxX; x++ {
			if checkVisible(trees, x, y) {
				numVisible++
			}
		}
	}
	return numVisible
}

func partB(filename string) int {
	// find the highest scenic score
	var highScore int
	trees := readInput(filename)
	for y := 0; y <= trees.MaxY; y++ {
		for x := 0; x <= trees.MaxX; x++ {
			score := findScenicScore(trees, x, y)
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func readInput(filename string) TreeMap {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var trees TreeMap
	for scanner.Scan() {
		line := scanner.Text()
		treeRow := make([]int, len(line))
		for i, v := range line {
			treeRow[i], err = strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
		}
		trees.Map = append(trees.Map, treeRow)
	}
	trees.MaxX = len(trees.Map[0]) - 1
	trees.MaxY = len(trees.Map) - 1
	return trees
}

func findScenicScore(trees TreeMap, x, y int) int {
	if x == 0 || y == 0 || x == trees.MaxX || y == trees.MaxY {
		// this one is on the edge so the score will be 0
		return 0
	}
	height := trees.Map[y][x]
	_, west := checkPathVisible(height, reverse(trees.Map[y][:x]))
	_, east := checkPathVisible(height, trees.Map[y][x+1:])
	var northTrees []int
	for i := y - 1; i >= 0; i-- {
		northTrees = append(northTrees, trees.Map[i][x])
	}
	_, north := checkPathVisible(height, northTrees)
	var southTrees []int
	//for i := y - 1; i > y; i-- {
	for i := y + 1; i <= trees.MaxY; i++ {
		southTrees = append(southTrees, trees.Map[i][x])
	}
	_, south := checkPathVisible(height, southTrees)
	return north * south * east * west
}

func checkVisible(trees TreeMap, x, y int) bool {
	// Trees at the edges are always visible
	if x == 0 || y == 0 || x == trees.MaxX || y == trees.MaxY {
		return true
	}
	height := trees.Map[y][x]
	// First check the row
	west, _ := checkPathVisible(height, trees.Map[y][:x])
	east, _ := checkPathVisible(height, trees.Map[y][x+1:])
	// now check the column
	var northTrees []int
	for i := 0; i < y; i++ {
		northTrees = append(northTrees, trees.Map[i][x])
	}
	north, _ := checkPathVisible(height, northTrees)

	var southTrees []int
	for i := y + 1; i <= trees.MaxY; i++ {
		southTrees = append(southTrees, trees.Map[i][x])
	}
	south, _ := checkPathVisible(height, southTrees)

	return north || south || west || east
}

func checkPathVisible(height int, path []int) (bool, int) {
	// For a slice of heights, return true only if all heights are less than the given height
	// and also return the distance that can be seen
	if path[0] >= height {
		return false, 1
	}
	if len(path) == 1 {
		return true, 1
	}
	visible, distance := checkPathVisible(height, path[1:])
	return visible, distance + 1
}

func reverse(s []int) []int {
	// Return a reversed copy of s
	s2 := make([]int, len(s))
	for i := 0; i < len(s)-1; i++ {
		s2[i] = s[len(s)-i-1]
	}
	return s2
}
