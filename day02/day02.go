package day02

import (
	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day02_1(filename string) (result int) {
	for row := range utils.InputRowsAsInts(filename) {
		result++
		firstDiff := row[1] - row[0]
		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]
			// if two numbers are the same
			if diff == 0 {
				result--
				break
			}
			// if it is decreasing series and diff greates than zero or less than -3
			if firstDiff < 0 && (diff > 0 || diff < -3) {
				result--
				break
			}
			// if it is increasing series and diff less than zero or greater than 3
			if firstDiff > 0 && (diff < 0 || diff > 3) {
				result--
				break
			}
		}
	}
	return result
}

func Day02_2(filename string) (result int) {
	for row := range utils.InputRowsAsInts(filename) {
		result++
		firstWrongReport := false
		firstDiff := row[1] - row[0]
		// fmt.Printf("%v\n", row)
		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]
			// fmt.Printf("   %d - %d = %d\n", row[i+1], row[i], diff)
			// if two numbers are the same
			if diff == 0 {
				if !firstWrongReport {
					// fmt.Println("   0: first")
					firstWrongReport = true
					continue
				}
				// fmt.Println("   0: second")
				result--
				break
			}
			// if it is decreasing series and diff greates than zero or less than -3
			if firstDiff < 0 && (diff > 0 || diff < -3) {
				if !firstWrongReport {
					// fmt.Println("   1: first")
					firstWrongReport = true
					continue
				}
				// fmt.Println("   1: second")
				result--
				break
			}
			// if it is increasing series and diff less than zero or greater than 3
			if firstDiff > 0 && (diff < 0 || diff > 3) {
				if !firstWrongReport {
					// fmt.Println("   2: first")
					firstWrongReport = true
					continue
				}
				// fmt.Println("   2: second")
				result--
				break
			}
		}
	}
	return result
}
