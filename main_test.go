package main

import (
	"testing"

	"github.com/rzabcio/adventofcode-2024/day01"
)

func TestDay01(t *testing.T) {
	got, want := day01.Day01_1("input-files/day01-test1"), 11
	if got != want {
		t.Errorf("Day01_1(test1) = %d; want %d", got, want)
	}
	got, want = day01.Day01_2("input-files/day01-test1"), 31
	if got != want {
		t.Errorf("Day01_2(test1) = %d; want %d", got, want)
	}
}
