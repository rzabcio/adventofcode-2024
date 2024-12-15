package day04

import (
	"regexp"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day04_1(filename string) (result int) {
	return findXmas(filename)
}

func Day04_2(filename string) (result int) {
	return findXmas(filename)
}

func findXmas(filename string) (result int) {
	ws := newWordSearch(filename)
	r1 := regexp.MustCompile("XMAS")
	r2 := regexp.MustCompile("SAMX")
	for _, row := range ws.rows {
		result += len(r1.FindAllStringIndex(row, -1))
		result += len(r2.FindAllStringIndex(row, -1))
	}
	for _, col := range ws.cols {
		result += len(r1.FindAllStringIndex(col, -1))
		result += len(r2.FindAllStringIndex(col, -1))
	}
	for _, diag := range ws.diags {
		result += len(r1.FindAllStringIndex(diag, -1))
		result += len(r2.FindAllStringIndex(diag, -1))
	}
	for _, aDiag := range ws.aDiags {
		result += len(r1.FindAllStringIndex(aDiag, -1))
		result += len(r2.FindAllStringIndex(aDiag, -1))
	}
	return result
}

type wordSearch struct {
	rows    []string
	rowsRev []string
	cols    []string
	diags   []string
	aDiags  []string
}

func newWordSearch(filetype string) (ws wordSearch) {
	ws = wordSearch{}
	for line := range utils.InputRows(filetype) {
		ws.rows = append(ws.rows, line)
		ws.rowsRev = append(ws.rowsRev, utils.ReverseString(line))
	}
	for _, col := range utils.InputCols(filetype) {
		ws.cols = append(ws.cols, col)
	}
	// diagonals
	for j := 0; j < len(ws.rows[0]); j++ {
		diag := ""
		for i := 0; i < len(ws.rows) && i <= j; i++ {
			diag += string(ws.rows[i][j-i])
		}
		ws.diags = append(ws.diags, diag)
	}
	for i := 1; i < len(ws.rows); i++ {
		diag := ""
		for j := len(ws.rows[0]) - 1; j >= 0 && i+len(ws.rows[0])-j-1 < len(ws.rows); j-- {
			diag += string(ws.rows[i+len(ws.rows[0])-j-1][j])
		}
		ws.diags = append(ws.diags, diag)
	}
	// anti-diagonals
	for j := 0; j < len(ws.rowsRev[0]); j++ {
		aDiag := ""
		for i := 0; i < len(ws.rowsRev) && i <= j; i++ {
			aDiag += string(ws.rowsRev[i][j-i])
		}
		ws.aDiags = append(ws.aDiags, aDiag)
	}
	for i := 1; i < len(ws.rowsRev); i++ {
		aDiag := ""
		for j := len(ws.rowsRev[0]) - 1; j >= 0 && i+len(ws.rowsRev[0])-j-1 < len(ws.rowsRev); j-- {
			aDiag += string(ws.rowsRev[i+len(ws.rowsRev[0])-j-1][j])
		}
		ws.aDiags = append(ws.aDiags, aDiag)
	}

	return ws
}
