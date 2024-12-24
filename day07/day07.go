package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day07_1(filename string) (result int) {
	for line := range utils.InputRows(filename) {
		eq := NewEquation(line)
		// fmt.Println(eq)
		eq.FindValidOperators()
		if eq.isValid() {
			result += eq.result
		}
	}
	return result
}

func Day07_2(filename string) (result int) {
	return result
}

type Equation struct {
	result     int
	components []int
	ops        []int
}

func NewEquation(line string) Equation {
	eq := Equation{}
	parts := strings.Split(line, ": ")
	eq.result, _ = strconv.Atoi(parts[0])
	for _, s := range strings.Fields(parts[1]) {
		component, _ := strconv.Atoi(s)
		eq.components = append(eq.components, component)
	}
	return eq
}

func (eq Equation) String() (r string) {
	return fmt.Sprintf("%d: %v", eq.result, eq.components)
}

func (eq *Equation) FindValidOperators() {
	maxOps := 1 << (len(eq.components) - 1)
	ops := 0
	for ops < maxOps {
		if eq.result == eq.OperatorsToResult(ops) {
			eq.ops = append(eq.ops, ops)
		}
		ops++
	}
}

func (eq *Equation) OperatorsToResult(ops int) (result int) {
	op := 1
	result = eq.components[0]
	opsString := ""
	for i := 1; i < len(eq.components); i++ {
		// fmt.Printf("ops & op: %d & %d = %d\n", ops, op, ops&op)
		if ops&op > 0 {
			result *= eq.components[i]
			opsString += "*"
		} else {
			result += eq.components[i]
			opsString += "+"
		}
		op = op << 1
	}
	// fmt.Printf("--- %v => %s = %d\n", eq, opsString, result)
	return result
}

func (eq Equation) isValid() bool {
	return len(eq.ops) > 0
}
