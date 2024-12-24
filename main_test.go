package main

import (
	"testing"

	"github.com/rzabcio/adventofcode-2024/day01"
	"github.com/rzabcio/adventofcode-2024/day02"
	"github.com/rzabcio/adventofcode-2024/day03"
	"github.com/rzabcio/adventofcode-2024/day05"
	"github.com/rzabcio/adventofcode-2024/day06"
	"github.com/rzabcio/adventofcode-2024/day07"
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

	got, want = day02.Day02_1("input-files/day02-test"), 2
	if got != want {
		t.Errorf("Day02_1(test1) = %d; want %d", got, want)
	}
	got, want = day02.Day02_2("input-files/day02-test"), 6
	if got != want {
		t.Errorf("Day02_2(test1) = %d; want %d", got, want)
	}

	got, want = day03.Day03_1("input-files/day03-test"), 161
	if got != want {
		t.Errorf("Day03_1(test) = %d; want %d", got, want)
	}
	got, want = day03.Day03_2("input-files/day03-test2"), 48
	if got != want {
		t.Errorf("Day03_2(test1) = %d; want %d", got, want)
	}

	// got, want = day04.Day04_1("input-files/day04-test"), 18
	// if got != want {
	// 	t.Errorf("Day04_1(test) = %d; want %d", got, want)
	// }
	// got, want = day04.Day04_2("input-files/day04-test"), 0
	// if got != want {
	// 	t.Errorf("Day04_2(test) = %d; want %d", got, want)
	// }

	got, want = day05.Day05_1("input-files/day05-test"), 143
	if got != want {
		t.Errorf("Day05_1(test) = %d; want %d", got, want)
	}
	got, want = day05.Day05_2("input-files/day05-test"), 0
	if got != want {
		t.Errorf("Day05_2(test) = %d; want %d", got, want)
	}

	got, want = day06.Day06_1("input-files/day06-test"), 41
	if got != want {
		t.Errorf("Day06_1(test) = %d; want %d", got, want)
	}
	got, want = day06.Day06_2("input-files/day06-test"), 0
	if got != want {
		t.Errorf("Day06_2(test) = %d; want %d", got, want)
	}

	got, want = day07.Day07_1("input-files/day07-test"), 3749
	if got != want {
		t.Errorf("Day07_1(test) = %d; want %d", got, want)
	}
	got, want = day07.Day07_2("input-files/day07-test"), 0
	if got != want {
		t.Errorf("Day07_2(test) = %d; want %d", got, want)
	}
}
