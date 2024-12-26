package day09

import (
	"fmt"
	"strconv"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day09_1(filename string) (result int) {
	dr := NewDiskReader(filename)
	dr.Defragment()
	result = dr.CalculateCheckSum()
	return result
}

func Day09_2(filename string) (result int) {
	return result
}

type DiskReader struct {
	diskMap         string
	disk            []int
	firstEmptyBlock int
}

func NewDiskReader(filename string) (dr DiskReader) {
	dr = DiskReader{}
	for dr.diskMap = range utils.InputRows(filename) {
	}
	dr.firstEmptyBlock = 0
	dr.ParseDiskMap()
	dr.NextFirstEmptyBlock()
	return dr
}

func (dr *DiskReader) ParseDiskMap() {
	fileID := 0
	fileBlock := true
	for _, r := range dr.diskMap {
		blockFill := -1
		blockLen, _ := strconv.Atoi(string(r))
		if fileBlock {
			blockFill = fileID
			fileID++
		}
		fileBlock = !fileBlock
		for i := 0; i < blockLen; i++ {
			dr.disk = append(dr.disk, blockFill)
		}
	}
}

func (dr *DiskReader) NextFirstEmptyBlock() int {
	for i := dr.firstEmptyBlock; i < len(dr.disk); i++ {
		if dr.disk[i] == -1 {
			dr.firstEmptyBlock = i
			return i
		}
	}
	return -1
}

func (dr *DiskReader) Defragment() {
	for i := len(dr.disk) - 1; i >= 0; i-- {
		if dr.disk[i] == -1 {
			continue
		}
		if i < dr.firstEmptyBlock {
			continue
		}
		dr.disk[dr.firstEmptyBlock] = dr.disk[i]
		dr.disk[i] = -1
		dr.NextFirstEmptyBlock()
	}
}

func (dr *DiskReader) CalculateCheckSum() (result int) {
	for i, block := range dr.disk {
		if block == -1 {
			break
		}
		result += i * block
	}
	return result
}

func (dr DiskReader) String() string {
	return fmt.Sprintf("%v\n", dr.disk)
}
