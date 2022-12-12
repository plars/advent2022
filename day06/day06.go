package main

import (
	"fmt"
	"os"
)

func checkUnique(packet []byte) bool {
	seen := make(map[byte]bool)
	for _, v := range packet {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func findStartOfPacketIndex() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := 4; i < len(data); i++ {
		if checkUnique(data[i-3 : i+1]) {
			// increment by 1 because we started the count at 0
			fmt.Println(i + 1)
			break
		}
	}
}

func findStartOfMessageIndex() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := 14; i < len(data); i++ {
		if checkUnique(data[i-13 : i+1]) {
			// increment by 1 because we started the count at 0
			fmt.Println(i + 1)
			break
		}
	}
}

func main() {
	fmt.Println("Part A:")
	findStartOfPacketIndex()
	fmt.Println("Part B:")
	findStartOfMessageIndex()
}
