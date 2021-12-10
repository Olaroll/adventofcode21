package day10

import (
	"github.com/Olaroll/adventofcode21/utils"
	"sort"
)

var dir = "./day10/"

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	values := make(map[rune]int)
	values[')'] = 3
	values[']'] = 57
	values['}'] = 1197
	values['>'] = 25137

	var sum int

endline:
	for _, line := range lines {
		var state []rune
		for _, char := range line {
			switch char {
			case '(':
				state = append(state, ')')
			case '[':
				state = append(state, ']')
			case '{':
				state = append(state, '}')
			case '<':
				state = append(state, '>')
			default:
				if len(state) == 0 || state[len(state)-1] != char {
					sum += values[char]
					continue endline
				}
				state = state[:len(state)-1]
			}
		}
	}

	return sum
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	values := make(map[rune]int)
	values[')'] = 1
	values[']'] = 2
	values['}'] = 3
	values['>'] = 4

	var scores []int

endline:
	for _, line := range lines {
		var state []rune
		for _, char := range line {
			switch char {
			case '(':
				state = append(state, ')')
			case '[':
				state = append(state, ']')
			case '{':
				state = append(state, '}')
			case '<':
				state = append(state, '>')
			default:
				if len(state) == 0 || state[len(state)-1] != char {
					continue endline
				}
				state = state[:len(state)-1]
			}
		}

		// Scoring
		var sum int
		for i := len(state) - 1; i >= 0; i-- {
			sum = sum*5 + values[state[i]]
		}
		scores = append(scores, sum)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
