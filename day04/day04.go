package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	overlaps := CountOverlaps("input.txt")
	fmt.Println("Part A:")
	fmt.Println(overlaps)
	overlaps = CountOverlaps2("input.txt")
	fmt.Println("Part B:")
	fmt.Println(overlaps)
}

func CountOverlaps(filename string) int {
	// Count only overlaps that are completely contained within the other
	var numOverlaps int
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		first := strings.Split(numbers[0], "-")
		second := strings.Split(numbers[1], "-")

		elf1Start, _ := strconv.Atoi(first[0])
		elf1End, _ := strconv.Atoi(first[1])
		elf2Start, _ := strconv.Atoi(second[0])
		elf2End, _ := strconv.Atoi(second[1])

		if elf1Start >= elf2Start && elf1End <= elf2End || elf2Start >= elf1Start && elf2End <= elf1End {
			numOverlaps++
		}
	}
	return numOverlaps
}

func CountOverlaps2(filename string) int {
	// Count any kind of overlap
	var numOverlaps int
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// split the line
		numbers := strings.Split(line, ",")
		first := strings.Split(numbers[0], "-")
		second := strings.Split(numbers[1], "-")

		elf1Start, _ := strconv.Atoi(first[0])
		elf1End, _ := strconv.Atoi(first[1])
		elf2Start, _ := strconv.Atoi(second[0])
		elf2End, _ := strconv.Atoi(second[1])

		if (elf1Start <= elf2Start && elf2Start <= elf1End) ||
			(elf1Start <= elf2End && elf2End <= elf1End) ||
			(elf2Start <= elf1Start && elf1Start <= elf2End) ||
			(elf2Start <= elf1End && elf1End <= elf2End) {
			numOverlaps++
		}
	}
	return numOverlaps
}
