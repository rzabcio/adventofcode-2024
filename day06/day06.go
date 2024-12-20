package day06

import (
	"strings"

	"github.com/rzabcio/adventofcode-2024/utils"
)

func Day06_1(filename string) (result int) {
	ml := NewManufacturingLab(filename)
	for ml.Next() {
	}
	return ml.counter
}

func Day06_2(filename string) (result int) {
	return result
}

type ManufacturingLab struct {
	area     [][]rune
	starting Guard
	current  Guard
	counter  int
}

func NewManufacturingLab(filename string) (ml ManufacturingLab) {
	ml = ManufacturingLab{}
	ml.counter = 1
	y := 0
	for line := range utils.InputRows(filename) {
		if strings.Contains(line, "^") {
			x := strings.Index(line, "^")
			ml.starting = Guard{x, y, 0, -1}
			ml.current = Guard{x, y, 0, -1}
			line = strings.ReplaceAll(line, "^", "x")
		}
		ml.area = append(ml.area, []rune(line))
		y++
	}
	return ml
}

func (ml *ManufacturingLab) Next() (result bool) {
	nextX := ml.current.NextX()
	nextY := ml.current.NextY()
	if nextX < 0 || nextX >= len(ml.area[0]) || nextY < 0 || nextY >= len(ml.area) {
		return false // out of boundary - finish
	}
	if ml.area[nextY][nextX] == '#' {
		ml.current.TurnRight()
		return true // just turn without an actual movement
	}
	ml.current.x = nextX
	ml.current.y = nextY
	if ml.area[ml.current.y][ml.current.x] != 'x' {
		ml.area[ml.current.y][ml.current.x] = 'x'
		ml.counter++
	}
	return true
}

func (ml ManufacturingLab) String() (result string) {
	for _, line := range ml.area {
		result += string(line) + "\n"
	}
	return result
}

type Guard struct {
	x, y       int
	dirX, dirY int
}

func (g *Guard) TurnRight() {
	if g.dirY == -1 { // up
		g.dirX = 1
		g.dirY = 0
	} else if g.dirX == 1 { // right
		g.dirX = 0
		g.dirY = 1
	} else if g.dirY == 1 { // down
		g.dirX = -1
		g.dirY = 0
	} else if g.dirX == -1 { // left
		g.dirX = 0
		g.dirY = -1
	}
}

func (g Guard) NextX() int {
	return g.x + g.dirX
}

func (g Guard) NextY() int {
	return g.y + g.dirY
}
