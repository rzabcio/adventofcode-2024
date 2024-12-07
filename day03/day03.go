package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day03_1(filename string) (result int) {
	for line := range utils.InputRows(filename) {
		r1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
		for _, mul := range r1.FindAllString(line, -1) {
			mul = strings.ReplaceAll(mul, "mul(", "")
			mul = strings.ReplaceAll(mul, ")", "")
			numbers := strings.Split(mul, ",")
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			result += n1 * n2
		}
	}
	return result
}

func Day03_2(filename string) (result int) {
	return result
}
