package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type pointInfo struct {
	point
	depth int
}

func main() {
	heightMap, start, end, possibleStarts := readMap("input.txt")
	paths := findPaths(heightMap)
	answerA, _ := shortestPath(heightMap, paths, start, end)
	fmt.Println("Answer A:", answerA)
	answerB := 9999
	for _, p := range possibleStarts {
		testB, err := shortestPath(heightMap, paths, p, end)
		if err != nil {
			continue
		}
		if testB < answerB {
			answerB = testB
		}
	}
	fmt.Println("Answer B:", answerB)
}

func shortestPath(heightMap [][]rune, paths map[point]([]point), start, end point) (int, error) {
	var toVisit []pointInfo
	var current pointInfo
	var visited map[point]bool = make(map[point]bool)
	toVisit = append(toVisit, pointInfo{start, 0})
	for current.point != end {
		if len(toVisit) == 0 {
			return -1, errors.New("No path found")
		}
		current = toVisit[0]
		toVisit = toVisit[1:]
		if visited[current.point] {
			continue
		}
		visited[current.point] = true
		for _, neighbor := range paths[current.point] {
			todo := pointInfo{neighbor, current.depth + 1}
			toVisit = append(toVisit, todo)
		}
	}
	return current.depth, nil
}

func readMap(filename string) ([][]rune, point, point, []point) {
	var line []rune
	var start, end point
	heightMap := make([][]rune, 0)
	// for partB, we need all the points with height 'a'
	possibleStarts := make([]point, 0)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for x, height := range scanner.Text() {
			if height == 'S' {
				start = point{x, len(heightMap)}
				height = 'a'
			}
			if height == 'E' {
				end = point{x, len(heightMap)}
				height = 'z'
			}
			if height == 'a' {
				possibleStarts = append(possibleStarts, point{x, len(heightMap)})
			}
			line = append(line, height)
		}
		heightMap = append(heightMap, line)
		line = make([]rune, 0)
	}
	return heightMap, start, end, possibleStarts
}

func findPaths(heightMap [][]rune) map[point]([]point) {
	var neighborHeight rune
	neighborOffsets := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	paths := make(map[point]([]point))
	for y, line := range heightMap {
		for x, height := range line {
			p := point{x, y}
			paths[p] = make([]point, 0)

			for _, offset := range neighborOffsets {
				neighborX, neighborY := x+offset.x, y+offset.y
				if neighborX < 0 || neighborX >= len(line) || neighborY < 0 || neighborY >= len(heightMap) {
					continue
				}
				neighborHeight = heightMap[y+offset.y][x+offset.x]
				if height-neighborHeight >= -1 {
					paths[p] = append(paths[p], point{x + offset.x, y + offset.y})
				}
			}
		}
	}
	return paths
}
