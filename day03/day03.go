package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) []string {
	//read filename and return a slice of strings
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func calcPriority(a string) int {
	var priority int
	for _, c := range a {
		val := int(c)
		if 97 <= val && val <= 122 {
			priority += val - 96
		} else if 65 <= val && val <= 90 {
			priority += val - 38
		}
	}
	return priority
}

func findSharedItemsA(a string, b string) string {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			return string(c)
		}
	}
	return ""
}

func findSharedItemsB(line []string) string {
	// slight difference, take 3 sets and only return the shared items if they
	// are in all 3 sets
	for _, c := range line[0] {
		if strings.Contains(line[1], string(c)) && strings.Contains(line[2], string(c)) {
			return string(c)
		}
	}
	return ""
}

func partA(data []string) {
	var sharedItems string
	var priority int

	for _, line := range data {
		set1 := line[:len(line)/2]
		set2 := line[len(line)/2:]
		sharedItems += findSharedItemsA(set1, set2)
	}
	priority += calcPriority(sharedItems)
	fmt.Println("Part A:")
	fmt.Println(priority)
}

func partB(data []string) {
	var sharedItems string
	var priority int

	for i := 0; i < len(data)-2; i += 3 {
		sharedItems += findSharedItemsB(data[i : i+3])
	}
	priority += calcPriority(sharedItems)
	fmt.Println("Part B:")
	fmt.Println(priority)
}

func main() {
	data := readFile("input.txt")
	partA(data)
	partB(data)
}
