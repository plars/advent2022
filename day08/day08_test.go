package main

import (
	"testing"
)

func Test_partA(t *testing.T) {
	expected := 21
	actual := partA("day08_test.txt")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func Test_partB(t *testing.T) {
	expected := 8
	actual := partB("day08_test.txt")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
