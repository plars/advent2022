package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getLargestThree(filename string) []int {
	input, _ := os.Open(filename)
	defer input.Close()

	var data []int
	var current int
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			data = append(data, current)
			current = 0
			continue
		}
		current += number
	}
	sort.Ints(data)
	return data[len(data)-3:]
}

func main() {
	largest := getLargestThree("input.txt")
	dataLen := len(largest)
	fmt.Println("Largest three (use largest for part a):")
	fmt.Println(largest[dataLen-3:])
	fmt.Println("Part B answer:")
	fmt.Println(largest[dataLen-3] + largest[dataLen-2] + largest[dataLen-1])
}
