package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	// worry level of items it holds
	items []int
	// Operation to perform after inspecting items
	//operation string
	operation func(int) int
	// number to divide by when testing
	testNum int
	// recipient monkey depending on the outcome of the test
	recipient map[bool]int
	// number of items inspected
	counter int
}

func (m *monkey) throw(monkeys []monkey, receiver int) {
	item := m.items[0]
	m.items = m.items[1:]
	monkeys[receiver].items = append(monkeys[receiver].items, item)
}

func (m *monkey) inspect(monkeys []monkey) {
	for len(m.items) > 0 {
		m.counter++
		m.items[0] = m.operation(m.items[0])
		testResult := m.items[0]%m.testNum == 0
		m.throw(monkeys, m.recipient[testResult])
	}
}

func main() {
	monkeys := readInput("input.txt")
	for i := 0; i < 20; i++ {
		for x := 0; x < len(monkeys); x++ {
			monkeys[x].inspect(monkeys)
		}
	}
	counters := make([]int, len(monkeys))
	for x := 0; x < len(monkeys); x++ {
		counters[x] = monkeys[x].counter
	}

	//sort the counters
	sort.IntSlice(counters).Sort()
	fmt.Println("Part A:")
	fmt.Println(counters[6] * counters[7])
}

func readInput(filename string) []monkey {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var monkeys []monkey
	var trueReceiver int
	var falseReceiver int

	for scanner.Scan() {
		monkey := new(monkey)
		// Ignore Monkey #: line
		// read Starting items: #, #, ...
		scanner.Scan()
		tokens := strings.Split(scanner.Text(), ": ")
		itemStrList := strings.Split(tokens[1], ", ")
		monkey.items = make([]int, len(itemStrList))
		for i, v := range itemStrList {
			monkey.items[i], _ = strconv.Atoi(v)
		}

		// read Operation: ...
		// All operations start with "new = old " (then an operation, then a number)
		scanner.Scan()
		tokens = strings.Split(scanner.Text(), "= old ")
		monkey.operation = buildOperation(tokens[1])

		// read Test: divisible by #
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &monkey.testNum)

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &trueReceiver)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &falseReceiver)
		monkey.recipient = map[bool]int{true: trueReceiver, false: falseReceiver}

		// read the extra blank line
		scanner.Scan()
		monkeys = append(monkeys, *monkey)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return monkeys
}

func buildOperation(operation string) func(int) int {
	// All we should have in the operation at this point is the operation and the number
	line := strings.Split(operation, " ")
	if line[1] == "old" {
		if line[0] == "*" {
			return func(x int) int { return (x * x) / 3 }
		}
		// otherwise it should be +
		return func(x int) int { return (x + x) / 3 }
	}
	// if the second number isn't old, it should be a number
	num, _ := strconv.Atoi(line[1])
	if line[0] == "*" {
		return func(x int) int { return (x * num) / 3 }
	}
	// otherwise it should be +
	return func(x int) int { return (x + num) / 3 }
}
