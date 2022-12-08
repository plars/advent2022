package main

import (
	"bufio"
	"fmt"
	"os"
)

func calcScore(filename string, scoreMap map[string]int) int {
	// Read the file and calculate the score
	input, _ := os.Open(filename)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var score int

	for scanner.Scan() {
		line := scanner.Text()
		score += scoreMap[line]
	}
	return score
}

func main() {
	scoreMapA := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	fmt.Println("Part A solution:")
	fmt.Println(calcScore("day02.txt", scoreMapA))
	scoreMapB := map[string]int{"B X": 1, "C Y": 6, "A Z": 8, "A X": 3, "B Y": 5, "C Z": 7, "C X": 2, "A Y": 4, "B Z": 9}
	fmt.Println("Part B solution:")
	fmt.Println(calcScore("day02.txt", scoreMapB))
}
