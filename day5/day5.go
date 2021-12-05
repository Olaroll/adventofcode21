package day5

import (
	"github.com/Olaroll/adventofcode21/utils"
	"log"
	"strings"
)

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
	switch {
	case vent.x1 == vent.x2:
		x := vent.x1
		if vent.y1 > vent.y2 {
			vent.y1, vent.y2 = vent.y2, vent.y1
		}

		for y := vent.y1; y <= vent.y2; y++ {
			b[y][x]++
		}

	case vent.y1 == vent.y2:
		y := vent.y1
		if vent.x1 > vent.x2 {
			vent.x1, vent.x2 = vent.x2, vent.x1
		}

		for x := vent.x1; x <= vent.x2; x++ {
			b[y][x]++
		}

	case (vent.x1 < vent.x2 && vent.y1 < vent.y2) || (vent.x1 > vent.x2 && vent.y1 > vent.y2):
		if vent.x1 > vent.x2 {
			vent.x1, vent.x2 = vent.x2, vent.x1
			vent.y1, vent.y2 = vent.y2, vent.y1
		}

		for i := 0; i <= vent.x2-vent.x1; i++ {
			b[vent.y1+i][vent.x1+i]++
		}

	case (vent.x1 < vent.x2 && vent.y1 > vent.y2) || (vent.x1 > vent.x2 && vent.y1 < vent.y2):
		if vent.x1 > vent.x2 {
			vent.x1, vent.x2 = vent.x2, vent.x1
			vent.y1, vent.y2 = vent.y2, vent.y1
		}

		for i := 0; i <= vent.x2-vent.x1; i++ {
			b[vent.y1-i][vent.x1+i]++
		}

	default:
		log.Fatalln("unhandled vent")
	}
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

func Solve1() int {
	lines := utils.GetLines("./day5/input.txt")

	vents := linesToVents(lines, true)
	board := makeBoard(1000)

	for _, vent := range vents {
		board.drawVent(vent)
	}

	return board.countGreaterThan(1)
}

func Solve2() int {
	lines := utils.GetLines("./day5/input.txt")

	vents := linesToVents(lines, false)
	board := makeBoard(1000)

	for _, vent := range vents {
		board.drawVent(vent)
	}

	return board.countGreaterThan(1)
}
