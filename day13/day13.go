package day13

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

var dir = "./day13/"

var Empty struct{}

type Dot struct{ x, y int }
type Dots map[Dot]struct{}

func foldY(index int, dots Dots) {
	for dot := range dots {
		if dot.y > index {
			delete(dots, dot)
			dot.y = dot.y - (dot.y-index)*2
			dots[dot] = Empty
		}
	}
}

func foldX(index int, dots Dots) {
	for dot := range dots {
		if dot.x > index {
			delete(dots, dot)
			dot.x = dot.x - (dot.x-index)*2
			dots[dot] = Empty
		}
	}
}

func printDots(dots Dots, sizeX, sizeY int) {
	board := make([][]rune, sizeY)
	for i := range board {
		board[i] = []rune(strings.Repeat(".", sizeX))
	}

	for dot := range dots {
		if dot.y < sizeY && dot.x < sizeX {
			board[dot.y][dot.x] = '#'
		}
	}

	for _, line := range board {
		fmt.Println(string(line))
	}
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	dots := make(Dots)

	for i := range lines {
		if lines[i] == "" {
			lines = lines[i+1:]
			break
		}

		xy := strings.Split(lines[i], ",")
		dot := Dot{x: utils.Atoi(xy[0]), y: utils.Atoi(xy[1])}
		dots[dot] = Empty
	}

	for i := range lines {
		split := strings.Split(lines[i][11:], "=")

		axis := split[0]
		index := utils.Atoi(split[1])

		switch axis {
		case "y":
			foldY(index, dots)
		case "x":
			foldX(index, dots)
		default:
			panic("invalid fold axis")
		}

		break // Solve1 is different only on this line
	}

	printDots(dots, 50, 8)

	return len(dots)
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	dots := make(Dots)

	for i := range lines {
		if lines[i] == "" {
			lines = lines[i+1:]
			break
		}

		xy := strings.Split(lines[i], ",")
		dot := Dot{x: utils.Atoi(xy[0]), y: utils.Atoi(xy[1])}
		dots[dot] = Empty
	}

	for i := range lines {
		split := strings.Split(lines[i][11:], "=")

		axis := split[0]
		index := utils.Atoi(split[1])

		switch axis {
		case "y":
			foldY(index, dots)
		case "x":
			foldX(index, dots)
		default:
			panic("invalid fold axis")
		}
	}

	printDots(dots, 50, 8)

	return len(dots)
}
