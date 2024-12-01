package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day01_1(filename string) (result int) {
	l, r := readFile(filename)
	sum := 0
	for i := range l {
		diff := r[i] - l[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func Day01_2(filename string) (result int) {
	return 0
}

func readFile(filename string) (lList, rList []int) {
	for line := range utils.InputRows(filename) {
		s := strings.Split(line, "   ")
		lInt, _ := strconv.Atoi(s[0])
		rInt, _ := strconv.Atoi(s[1])
		lList = append(lList, lInt)
		rList = append(rList, rInt)
	}
	sort.Ints(lList)
	sort.Ints(rList)
	return lList, rList
}
