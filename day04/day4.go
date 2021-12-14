package day04

import (
	"github.com/Olaroll/adventofcode21/utils"
	"log"
	"strings"
)

var dir = "./day04/"

// Not too happy with today's solutions and how long they took,
// but at least I got it done in the end

type Board [][]Cell

type Cell struct {
	number int
	found  bool
}

func makeCells(ints []int) []Cell {
	cells := make([]Cell, len(ints))
	for i, v := range ints {
		cells[i].number = v
	}
	return cells
}

func bingo(board Board, num int) bool {
	for y := range board {
		for x := range board[y] {
			if board[y][x].number == num {
				board[y][x].found = true
				if checkBingo(board, x, y) {
					return true
				}
			}
		}
	}
	return false
}

func checkBingo(board Board, x, y int) bool {
	// Check row
	for _, cell := range board[y] {
		if !cell.found {
			goto noRow
		}
	}
	return true
noRow:

	// Check column
	for yi := 0; yi < len(board); yi++ {
		if !board[yi][x].found {
			return false
		}
	}
	return true

}

func sumNotFound(board Board) int {
	var sum int
	for y := range board {
		for x := range board[y] {
			if !board[y][x].found {
				sum += board[y][x].number
			}
		}
	}
	return sum
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	nums := utils.AtoiSlc(strings.Split(lines[0], ","))
	lines = lines[1:]

	var boards []Board
	for len(lines) > 1 {
		lines = lines[1:]
		board := make(Board, 5)
		for i := 0; i < 5; i++ {
			ints := utils.AtoiSlc(strings.Fields(lines[i]))
			board[i] = makeCells(ints)
		}
		boards = append(boards, board)
		lines = lines[5:]
	}

	for _, num := range nums {
		for i := 0; i < len(boards); i++ {
			if bingo(boards[i], num) {
				return sumNotFound(boards[i]) * num
			}
		}
	}

	log.Fatalln("Could not find bingo")
	return 0
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	nums := utils.AtoiSlc(strings.Split(lines[0], ","))
	lines = lines[1:]

	var boards []Board
	for len(lines) > 1 {
		lines = lines[1:]
		board := make(Board, 5)
		for i := 0; i < 5; i++ {
			ints := utils.AtoiSlc(strings.Fields(lines[i]))
			board[i] = makeCells(ints)
		}
		boards = append(boards, board)
		lines = lines[5:]
	}

	for _, num := range nums {
		for i := 0; i < len(boards); i++ {
			if bingo(boards[i], num) {
				if len(boards) == 1 {
					return sumNotFound(boards[i]) * num
				} else {
					boards[i] = boards[len(boards)-1]
					boards = boards[:len(boards)-1]
					i--
				}
			}
		}
	}

	log.Fatalln("Could not find bingo")
	return 0
}
