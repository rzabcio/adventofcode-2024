package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day03_1(filename string) (result int) {
	return multiply(filename)
}

func Day03_2(filename string) (result int) {
	return multiply(filename)
}

func multiply(filename string) (result int) {
	enabled := true
	for line := range utils.InputRows(filename) {
		r1 := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
		for _, part := range r1.FindAllString(line, -1) {
			if part == "don't()" {
				enabled = false
				continue
			} else if part == "do()" {
				enabled = true
				continue
			} else if !enabled {
				continue
			}
			part = strings.ReplaceAll(part, "mul(", "")
			part = strings.ReplaceAll(part, ")", "")
			numbers := strings.Split(part, ",")
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			result += n1 * n2
		}
	}
	return result
}
