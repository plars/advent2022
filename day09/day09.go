package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	answerA := partA("input.txt")
	fmt.Println("Part A:", answerA)
}

func partA(filename string) int {
	visited := map[string]int{"0,0": 1}
	data := readFile(filename)
	head := []int{0, 0}
	tail := []int{0, 0}
	for _, line := range data {
		if line == "" {
			continue
		}
		move := strings.Split(line, " ")
		newTailLocations := moveRope(move, head, tail)
		for _, coords := range newTailLocations {
			if _, exists := visited[coords]; exists == false {
				visited[coords] = 1
			} else {
				visited[coords]++
			}
		}
	}
	return len(visited)
}

func readFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func moveRope(move []string, head, tail []int) []string {
	// return the new locations the tail has moved to
	var newTailLocations []string
	numMoves, err := strconv.Atoi(move[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < numMoves; i++ {
		oldHead := make([]int, 2)
		copy(oldHead, head)
		switch move[0] {
		case "U":
			head[1]++
		case "D":
			head[1]--
		case "L":
			head[0]--
		case "R":
			head[0]++
		}
		if math.Abs(float64(head[0]-tail[0])) <= 1.0 && math.Abs(float64(head[1]-tail[1])) <= 1.0 {
			// if we are adjacent to the tail, don't move the tail
			continue
		}
		if head[0] != tail[0] && head[1] != tail[1] {
			// head has moved diagonally, so move the tail to the old head position
			copy(tail, oldHead)
		} else if head[0] != tail[0] {
			// otherwise we moved along just one axis, so move the tail along that axis
			tail[0] = (tail[0] + head[0]) / 2
		} else {
			tail[1] = (tail[1] + head[1]) / 2
		}
		newTailLocations = append(newTailLocations, fmt.Sprintf("%d,%d", tail[0], tail[1]))
	}
	return newTailLocations
}
