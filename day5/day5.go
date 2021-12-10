package day5

import (
	"github.com/Olaroll/adventofcode21/utils"
	"math"
	"strings"
)

var dir = "./day5/"

type Vent struct {
	x1, y1 int
	x2, y2 int
}

type Board [][]int

func linesToVents(lines []string, simple bool) []Vent {
	vents := make([]Vent, 0, 500)
	for _, line := range lines {
		oneTwo := strings.Split(line, " -> ")
		one := strings.Split(oneTwo[0], ",")
		two := strings.Split(oneTwo[1], ",")

		if simple && one[0] != two[0] && one[1] != two[1] {
			continue
		}

		var vent Vent
		vent.x1 = utils.Atoi(one[0])
		vent.y1 = utils.Atoi(one[1])
		vent.x2 = utils.Atoi(two[0])
		vent.y2 = utils.Atoi(two[1])
		vents = append(vents, vent)
	}

	return vents
}

func makeBoard(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	return board
}

func (b Board) drawVent(vent Vent) {
	plotLine(vent.x1, vent.y1, vent.x2, vent.y2, func(x, y int) {
		b[y][x]++
	})
}

func (b Board) countGreaterThan(num int) int {
	var count int
	for _, row := range b {
		for _, cell := range row {
			if cell > num {
				count++
			}
		}
	}
	return count
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	vents := linesToVents(lines, true)
	board := makeBoard(1000)

	for _, vent := range vents {
		board.drawVent(vent)
	}

	return board.countGreaterThan(1)
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	vents := linesToVents(lines, false)
	board := makeBoard(1000)

	for _, vent := range vents {
		board.drawVent(vent)
	}

	return board.countGreaterThan(1)
}

// LINE PLOTTING STUFF (from https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm)

func plotLine(x0, y0, x1, y1 int, plot func(x, y int)) {
	if math.Abs(float64(y1-y0)) < math.Abs(float64(x1-x0)) {
		if x0 > x1 {
			plotLineLow(x1, y1, x0, y0, plot)
		} else {
			plotLineLow(x0, y0, x1, y1, plot)
		}
	} else {
		if y0 > y1 {
			plotLineHigh(x1, y1, x0, y0, plot)
		} else {
			plotLineHigh(x0, y0, x1, y1, plot)
		}
	}
}

func plotLineLow(x0, y0, x1, y1 int, plot func(x, y int)) {
	dx := x1 - x0
	dy := y1 - y0

	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}

	D := (2 * dy) - dx
	y := y0

	for x := x0; x <= x1; x++ {
		plot(x, y)
		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + 2*dy
		}
	}
}

func plotLineHigh(x0, y0, x1, y1 int, plot func(x, y int)) {
	dx := x1 - x0
	dy := y1 - y0

	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}

	D := (2 * dx) - dy
	x := x0

	for y := y0; y <= y1; y++ {
		plot(x, y)
		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + 2*dx
		}
	}
}
