package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Stack []string

func (s *Stack) push(v string) {
	*s = append(*s, v)
}

// push a given number of elements onto the stack
func (s *Stack) pushN(v []string) {
	*s = append(*s, v...)
}

func (s *Stack) reverse() {
	var newStack Stack
	max := len(*s)
	for i := 0; i < max; i++ {
		newStack.push(s.pop())
	}
	*s = newStack
}

func (s *Stack) pop() string {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

// pop a given number of elements from the stack
func (s *Stack) popN(n int) []string {
	l := len(*s)
	popped := (*s)[l-n:]
	*s = (*s)[:l-n]
	return popped
}

func (s *Stack) peek() string {
	l := len(*s)
	return (*s)[l-1]
}

func readFile(filename string) []string {
	//read filename and return a slice of strings
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return Stack(lines)
}

/*
Take a scanner and read an array of 9 stacks from a file in the format:
[X] [H]     [Q]
[R] [P] [F] [J] [B] [C] [M] [D] [G]

	1   2   3   4   5   6   7   8   9

Everything after the line with numbers can be ignored for now
*/
func readStacks(scanner *bufio.Scanner) ([]Stack, []string) {
	// The top part of the file is the stacks, followed by the moves
	var stacks []Stack = make([]Stack, 9)
	var moves []string

	for scanner.Scan() {
		line := scanner.Text()
		// if the line has a stack number in it, we're done reading the stacks
		if strings.Contains(line, "1") {
			break
		}
		// read the items and push them on the right stack if it exists
		for stackNum := 0; stackNum < 9; stackNum++ {
			item := string(line[1+stackNum*4])
			// handle empty item for a stack
			if item == " " {
				continue
			}
			//push the item on the stack
			stacks[stackNum].push(item)
		}
	}
	// now reverse the stacks because we pushed from top to bottom
	for i := 0; i < 9; i++ {
		stacks[i].reverse()
	}

	// read the blank line
	scanner.Scan()
	// Now read the moves
	for scanner.Scan() {
		moves = append(moves, scanner.Text())
	}
	return stacks, moves
}

func processMoves(stacks []Stack, moves []string) {
	var move, from, to int
	for _, line := range moves {
		n, err := fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		if err != nil {
			fmt.Println(n)
			fmt.Println(err)
		}
		for i := 0; i < move; i++ {
			item := stacks[from-1].pop()
			stacks[to-1].push(item)
		}
	}
}

// PWL func processMovesN(stacks []Stack, scanner *bufio.Scanner) {
func processMovesN(stacks []Stack, moves []string) {
	// just like processMoves, but pop and push N items for day05b
	var move, from, to int
	for _, line := range moves {
		n, err := fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		if err != nil {
			fmt.Println(n)
			fmt.Println(err)
		}
		items := stacks[from-1].popN(move)
		stacks[to-1].pushN(items)
	}
}

func processInput(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	stacksA, moves := readStacks(scanner)

	// make a copy of the stacks for part b
	var stacksB []Stack = make([]Stack, len(stacksA))
	copy(stacksB, stacksA)

	// remove the empty line
	_ = scanner.Scan()

	fmt.Println("Part A:")
	processMoves(stacksA, moves)
	for i := 0; i < 9; i++ {
		fmt.Print(stacksA[i].peek())
	}
	fmt.Println()

	fmt.Println("Part B:")
	processMovesN(stacksB, moves)
	for i := 0; i < 9; i++ {
		fmt.Print(stacksB[i].peek())
	}
}

func main() {
	processInput("input.txt")
}
