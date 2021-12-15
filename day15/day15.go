package day15

import (
	"container/heap"
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

var dir = "./day15/"

type Board [][]int

func (board Board) isValid(x, y int) bool {
	return y >= 0 &&
		x >= 0 &&
		y < len(board) &&
		x < len(board[y])
}

func (board Board) repeat(times int) Board {
	state := board.copy()

	for i := 1; i < times; i++ {
		state.append(board.plusN(i))
	}

	stateRow := state.copy()

	for i := 1; i < times; i++ {
		state = append(state, stateRow.plusN(i)...)
	}

	return state
}

func (board Board) plusN(n int) Board {
	newBoard := board.copy()
	for y := range newBoard {
		for x := range newBoard[y] {
			newBoard[y][x] += n
			if newBoard[y][x] > 9 {
				newBoard[y][x] %= 9
			}
		}
	}
	return newBoard
}

func (board Board) copy() Board {
	newBoard := make(Board, len(board))
	for y := range board {
		newBoard[y] = make([]int, len(board[y]))
		copy(newBoard[y], board[y])
	}
	return newBoard
}

func (board Board) append(board2 Board) {
	for y := range board {
		board[y] = append(board[y], board2[y]...)
	}
}

type Path struct {
	length int
	x, y   int
}

type PathQueue []Path

func (pq PathQueue) Len() int {
	return len(pq)
}

func (pq PathQueue) Less(i, j int) bool {
	return pq[i].length < pq[j].length
}

func (pq PathQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PathQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Path))
}

func (pq *PathQueue) Pop() interface{} {
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return item
}

var neighbours = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func (pq *PathQueue) advance(board Board) int {
	path := heap.Pop(pq).(Path)

	for _, off := range neighbours {
		x, y := off[0]+path.x, off[1]+path.y
		if board.isValid(x, y) && board[y][x] >= 0 {
			newPath := path
			newPath.x = x
			newPath.y = y

			newPath.length += board[newPath.y][newPath.x]
			heap.Push(pq, newPath)
			heap.Fix(pq, len(*pq)-1)
			board[newPath.y][newPath.x] = -1

			if newPath.y == len(board)-1 && newPath.x == len(board[newPath.y])-1 {
				return newPath.length
			}
		}
	}

	return -1
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for i := range lines {
		board[i] = utils.AtoiSlc(strings.Split(lines[i], ""))
	}

	pq := &PathQueue{Path{}}
	board[0][0] = -1

	var result int
	for {
		result = pq.advance(board)
		if result != -1 {
			break
		}
	}

	return result
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	board := make(Board, len(lines))
	for i := range lines {
		board[i] = utils.AtoiSlc(strings.Split(lines[i], ""))
	}

	board = board.repeat(5)

	pq := &PathQueue{Path{}}
	board[0][0] = -1

	var result int
	for {
		result = pq.advance(board)
		if result != -1 {
			break
		}
	}

	return result
}
