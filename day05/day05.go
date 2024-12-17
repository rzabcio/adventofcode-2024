package day05

import (
	"math"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

type SafetyManual struct {
	rules   map[int]map[int]bool
	updates [][]int
}

func NewSafetyManual(filename string) (sm SafetyManual) {
	sm = SafetyManual{}
	sm.rules = make(map[int]map[int]bool)

	for line := range utils.InputRows(filename) {
		if strings.Contains(line, "|") {
			// read rules: p1|p2 to map (p1 -> p2,p3,...)
			ruleStrings := strings.Split(line, "|")
			p1, _ := strconv.Atoi(ruleStrings[0])
			p2, _ := strconv.Atoi(ruleStrings[1])
			if sm.rules[p1] == nil {
				sm.rules[p1] = make(map[int]bool)
			}
			sm.rules[p1][p2] = true
		} else if strings.Contains(line, ",") {
			// read updates: p1,p2,p3,... to list
			update := []int{}
			for _, pageString := range strings.Split(line, ",") {
				page, _ := strconv.Atoi(pageString)
				update = append(update, page)
			}
			sm.updates = append(sm.updates, update)
		}
	}
	return
}

func (sm *SafetyManual) MiddlePagesSum() (sum int) {
	for i := 0; i < len(sm.updates); i++ {
		update := sm.updates[i]
		if sm.IsUpdateValid(i) {
			// if update valid add to sum middle page
			middlePage := update[int(math.Ceil(float64(len(update)/2.0)))]
			sum += middlePage
		}
	}
	return sum
}

func (sm *SafetyManual) IsUpdateValid(i int) bool {
	// iterate over all pages besiade the first one
	for j := 1; j < len(sm.updates[i]); j++ {
		page := sm.updates[i][j]
		// compare with all previous pages (should be entry current|previous in the rules)
		for k := j - 1; k >= 0; k-- {
			pageOther := sm.updates[i][k]
			val, ok := sm.rules[page][pageOther]
			if ok && val { // if exists and is true (probably overkill, but well
				return false
			}
		}
	}
	return true
}

func Day05_1(filename string) (result int) {
	sm := NewSafetyManual(filename)
	result = sm.MiddlePagesSum()
	return result
}

func Day05_2(filename string) (result int) {
	return result
}
