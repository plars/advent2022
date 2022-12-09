package main

import (
	"testing"
)

func Test_CountOverlaps(t *testing.T) {
	overlaps := CountOverlaps("day04_test.txt")
	expected := 2
	if overlaps != expected {
		t.Errorf("Expected %d, got %d", expected, overlaps)
	}
}
