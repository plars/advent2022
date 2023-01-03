package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readFile("input.txt")
	cycles, xvals := interpret(data)
	fmt.Println("Part A:")
	solutionA := partA(cycles, xvals)
	fmt.Println(solutionA)

	// This one is visual - you need to look at the output
	partB(xvals)
}

func readFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func interpret(data []string) ([]int, []int) {
	cycles := []int{1}
	xvals := []int{1}
	for _, line := range data {
		if line == "" {
			continue
		}
		// we start by repeating the last value no matter what
		cycles = append(cycles, 1+cycles[len(cycles)-1])
		xvals = append(xvals, xvals[len(xvals)-1])

		if line[:4] == "noop" {
			// noop, we're done and can move on
			continue
		}

		// otherwise it has to be an addx
		addx := strings.Split(line, " ")
		x, err := strconv.Atoi(addx[1])
		if err != nil {
			panic(err)
		}
		cycles = append(cycles, 1+cycles[len(cycles)-1])
		xvals = append(xvals, x+xvals[len(xvals)-1])
	}
	return cycles, xvals
}

func partA(cycles, xvals []int) int {
	signalStrength := 20 * xvals[20-1]
	signalStrength += 60 * xvals[60-1]
	signalStrength += 100 * xvals[100-1]
	signalStrength += 140 * xvals[140-1]
	signalStrength += 180 * xvals[180-1]
	signalStrength += 220 * xvals[220-1]
	return signalStrength
}

func partB(xvals []int) {
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			fmt.Println()
		}
		if math.Abs(float64(xvals[i]-(i%40))) <= 1.0 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}
