package day08

import (
	"fmt"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day08_1(filename string) (result int) {
	am := NewAntennaMap(filename)
	fmt.Println(am)
	return result
}

func Day08_2(filename string) (result int) {
	return result
}

type AntennaMap struct {
	area         map[int]map[int]Antenna
	sizeX, sizeY int
}

func NewAntennaMap(filename string) (am AntennaMap) {
	am = AntennaMap{}
	am.area = make(map[int]map[int]Antenna)
	y := 0
	for line := range utils.InputRows(filename) {
		am.sizeX = len(line)
		for x, r := range line {
			if _, ok := am.area[x]; !ok {
				am.area[x] = make(map[int]Antenna)
			}
			a := NewAntenna(r, x, y)
			am.area[x][y] = a
		}
		y++
	}
	am.sizeY = y
	return am
}

func (am AntennaMap) String() (s string) {
	for y := 0; y < am.sizeY; y++ {
		for x := 0; x < am.sizeX; x++ {
			if a, ok := am.area[x][y]; ok {
				s += string(a.antenna)
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

type Antenna struct {
	x, y      int
	antenna   rune
	antinodes []string
}

func NewAntenna(r rune, x, y int) (a Antenna) {
	a = Antenna{}
	a.x = x
	a.y = y
	a.antenna = r
	return a
}

func (a Antenna) String() (s string) {
	return fmt.Sprintf("[%d,%d] %c (%v)", a.x, a.y, a.antenna, a.antinodes)
}
