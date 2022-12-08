package main

import (
	"reflect"
	"testing"
)

func Test_getLargestThree(t *testing.T) {
	largest := getLargestThree("day01_test.txt")
	expected := []int{10000, 11000, 24000}
	if !reflect.DeepEqual(largest, expected) {
		t.Errorf("Expected %v, got %v", expected, largest)
	}
	/*
		if largest != expected {
			t.Errorf("Expected 10, got %d", largest)
		}
	*/
}
