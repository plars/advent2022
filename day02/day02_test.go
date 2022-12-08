package main

import "testing"

func Test_calcScore(t *testing.T) {
	scoreMapA := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	score := calcScore("day02_test.txt", scoreMapA)
	expected := 15
	if score != expected {
		t.Errorf("Expected %d, got %d", expected, score)
	}
}
