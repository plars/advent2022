package main

import "testing"

func Test_partA(t *testing.T) {
	data := readFile("day10_test.txt")
	cycles, xvals := interpret(data)
	answer := partA(cycles, xvals)
	expected := 13140
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
	}
}
