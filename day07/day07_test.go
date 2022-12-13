package main

import (
	"testing"
)

func Test_day07a(t *testing.T) {
	root := processInput("day07_test.txt")
	sumA := sumSizeUnder(root, 100000)
	expected := 95437
	if sumA != expected {
		t.Errorf("Expected %d, got %d", expected, sumA)
	}
}

func Test_day07b(t *testing.T) {
	root := processInput("day07_test.txt")
	minDeleteSize := root.Size - 40000000
	sumB := findSmallestDirOver(root, minDeleteSize)
	expected := 24933642
	if sumB != expected {
		t.Errorf("Expected %d, got %d", expected, sumB)
	}
}
