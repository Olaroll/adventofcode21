package day14

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

var dir = "./day14/"

func makeRules(lines []string) map[[2]byte]byte {
	rules := make(map[[2]byte]byte)

	for _, line := range lines {
		split := strings.Split(line, " -> ")

		var key [2]byte
		copy(key[:], split[0])

		rules[key] = []byte(split[1])[0]
	}

	return rules
}

func getMinMax(counts map[byte]int) (min int, max int) {
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	return
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	state := lines[0]
	lines = lines[2:]

	rules := makeRules(lines)

	fmt.Printf("Initial: %v\n", state)
	for i := 0; i < 10; i++ {
		state = polymerizeNaive(state, rules)
		fmt.Printf("Step %v: %v\n", i+1, state)
	}

	counts := make(map[byte]int)
	for _, char := range state {
		counts[byte(char)]++
	}

	min, max := getMinMax(counts)

	return max - min
}

func polymerizeNaive(initial string, rules map[[2]byte]byte) string {
	state := make([]byte, len(initial), len(initial)*2)
	copy(state, initial)

	for i := len(state) - 1; i > 0; i-- {
		key := [2]byte{state[i-1], state[i]}

		insert, ok := rules[key]
		if !ok {
			continue
		}

		state = append(state[:i+1], state[i:]...)
		state[i] = insert
	}

	return string(state)
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	init := lines[0]
	lines = lines[2:]

	rules := makeRules(lines)

	state, counts := processInit(init)
	for i := 0; i < 40; i++ {
		state, counts = polymerize(state, rules, counts)
	}

	min, max := getMinMax(counts)

	return max - min
}

func polymerize(state map[[2]byte]int, rules map[[2]byte]byte, counts map[byte]int) (map[[2]byte]int, map[byte]int) {
	newState := make(map[[2]byte]int)
	for k, v := range state {
		insert, ok := rules[k]
		if !ok {
			continue
		}

		newState[[2]byte{k[0], insert}] += v
		newState[[2]byte{insert, k[1]}] += v

		counts[insert] += v
	}

	return newState, counts
}

func processInit(init string) (map[[2]byte]int, map[byte]int) {
	state := make(map[[2]byte]int)
	counts := make(map[byte]int)

	for i := 0; i < len(init)-1; i++ {
		counts[init[i]]++
		state[[2]byte{init[i], init[i+1]}]++
	}
	counts[init[len(init)-1]]++
	return state, counts
}
