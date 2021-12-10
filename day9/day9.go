package day9

import (
	"github.com/Olaroll/adventofcode21/utils"
	"sort"
)

var dir = "./day9/"

type Board [][]byte

func checkLow(board Board, x, y int) int {
	if y > 0 && board[y-1][x] <= board[y][x] {
		return 0
	}
	if y < len(board)-1 && board[y+1][x] <= board[y][x] {
		return 0
	}
	if x > 0 && board[y][x-1] <= board[y][x] {
		return 0
	}
	if x < len(board[y])-1 && board[y][x+1] <= board[y][x] {
		return 0
	}
	return int(board[y][x] + 1)
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for y := range board {
		board[y] = []byte(lines[y])
		for x := range board[y] {
			board[y][x] -= '0'
		}
	}

	var sum int
	for y := range board {
		for x := range board[y] {
			sum += checkLow(board, x, y)
		}
	}

	return sum
}

func expand(board Board, x, y int) int {
	counter := 1
	board[y][x] = 9

	// Check UP
	if y > 0 && board[y-1][x] != 9 {
		counter += expand(board, x, y-1)
	}

	// Check DOWN
	if y < len(board)-1 && board[y+1][x] != 9 {
		counter += expand(board, x, y+1)
	}

	// Check LEFT
	if x > 0 && board[y][x-1] != 9 {
		counter += expand(board, x-1, y)
	}

	// Check RIGHT
	if x < len(board[y])-1 && board[y][x+1] != 9 {
		counter += expand(board, x+1, y)
	}

	return counter
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for y := range board {
		board[y] = []byte(lines[y])
		for x := range board[y] {
			board[y][x] -= '0'
		}
	}

	var sizes []int
	for y := range board {
		for x := range board[y] {
			if checkLow(board, x, y) > 0 {
				count := expand(board, x, y)
				sizes = append(sizes, count)
			}
		}
	}

	sort.Ints(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}
