package day11

import (
	"github.com/Olaroll/adventofcode21/utils"
)

var dir = "./day11/"

type Board [][]int

func (board Board) step() int {
	var flashed int
	for y := range board {
		for x := range board {
			board[y][x]++

			if board[y][x] > 9 {
				flashed += board.flash(x, y)
			}
		}
	}

	for y := range board {
		for x := range board {
			if board[y][x] < 0 {
				board[y][x] = 0
			}
		}
	}

	return flashed
}

var neighbors = [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

func (board Board) flash(x, y int) int {
	flashed := 1
	board[y][x] = -99

	for _, coords := range neighbors {
		x2, y2 := x+coords[0], y+coords[1]

		if board.isValid(x2, y2) {
			board[y2][x2]++

			if board[y2][x2] > 9 {
				flashed += board.flash(x2, y2)
			}
		}
	}

	return flashed
}

func (board Board) isValid(x, y int) bool {
	return x >= 0 &&
		y >= 0 &&
		x < len(board[0]) &&
		y < len(board)
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for y, line := range lines {
		board[y] = make([]int, len(lines[y]))
		for x, c := range line {
			n := int(c - '0')
			board[y][x] = n
		}
	}

	var sum int
	for i := 0; i < 100; i++ {
		sum += board.step()
	}

	return sum
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for y, line := range lines {
		board[y] = make([]int, len(lines[y]))
		for x, char := range line {
			n := int(char - '0')
			board[y][x] = n
		}
	}

	var turn int
	for {
		turn++
		flashed := board.step()
		if flashed == len(board)*len(board[0]) {
			break
		}
	}

	return turn
}
