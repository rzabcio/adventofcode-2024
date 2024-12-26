package day08

import (
	"fmt"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day08_1(filename string) (result int) {
	am := NewAntennaMap(filename)
	result = am.FindAllAntinodes()
	return result
}

func Day08_2(filename string) (result int) {
	return result
}

type AntennaMap struct {
	area         map[int]map[int]*Antenna
	antennas     map[rune][]*Antenna
	sizeX, sizeY int
}

func NewAntennaMap(filename string) (am AntennaMap) {
	am = AntennaMap{}
	am.area = make(map[int]map[int]*Antenna)
	am.antennas = make(map[rune][]*Antenna)
	y := 0
	for line := range utils.InputRows(filename) {
		am.sizeX = len(line)
		for x, r := range line {
			if _, ok := am.area[x]; !ok {
				am.area[x] = make(map[int]*Antenna)
			}
			a := NewAntenna(r, x, y)
			am.area[x][y] = a
			am.antennas[r] = append(am.antennas[r], a)
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
				s += a.String()
			} else {
				s += ". "
			}
		}
		s += "\n"
	}
	return s
}

func (am *AntennaMap) FindAllAntinodes() int {
	for r := range am.antennas {
		if r == '.' {
			continue
		}
		am.FindAntinodes(r)
	}
	return am.CountAntinodes()
}

func (am *AntennaMap) FindAntinodes(r rune) {
	for i := 0; i < len(am.antennas[r])-1; i++ {
		a1 := am.antennas[r][i]
		for j := i + 1; j < len(am.antennas[r]); j++ {
			a2 := am.antennas[r][j]
			am.FindAntinodesForAntennas(a1, a2)
		}
	}
}

func (am *AntennaMap) FindAntinodesForAntennas(a1, a2 *Antenna) {
	distX := a1.x - a2.x
	distY := a1.y - a2.y
	// antinode 1
	antiX, antiY := a1.x+distX, a1.y+distY
	am.MarkAntinode(antiX, antiY, a1.antenna)
	// antinode 2
	antiX, antiY = a2.x-distX, a2.y-distY
	am.MarkAntinode(antiX, antiY, a2.antenna)
}

func (am *AntennaMap) MarkAntinode(x, y int, r rune) bool {
	var a *Antenna
	var ok bool
	if am.InBounds(x, y) {
		if a, ok = am.area[x][y]; !ok {
			a = NewAntenna('.', x, y)
		}
		a.antinodes = append(a.antinodes, r)
		am.area[x][y] = a
		return true
	}
	return false
}

func (am *AntennaMap) InBounds(x, y int) bool {
	return x >= 0 && x < am.sizeX && y >= 0 && y < am.sizeY
}

func (am *AntennaMap) CountAntinodes() (count int) {
	for y := 0; y < am.sizeY; y++ {
		for x := 0; x < am.sizeX; x++ {
			if a, ok := am.area[x][y]; ok {
				if a.IsAntinode() {
					count++
				}
			}
		}
	}
	return count
}

type Antenna struct {
	x, y      int
	antenna   rune
	antinodes []rune
}

func NewAntenna(r rune, x, y int) (a *Antenna) {
	a = &Antenna{}
	a.x = x
	a.y = y
	a.antenna = r
	return a
}

func (a *Antenna) IsAntinode() bool {
	return len(a.antinodes) > 0
}

func (a Antenna) String() (s string) {
	if a.IsAntinode() {
		return fmt.Sprintf("%c#", a.antenna)
	}
	return fmt.Sprintf("%c ", a.antenna)
}
