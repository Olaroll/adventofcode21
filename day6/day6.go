package day6

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

var dir = "./day6/"

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)
	fish := utils.AtoiSlc(strings.Split(lines[0], ","))

	DAYS := 80

	for i := 0; i < DAYS; i++ {
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 7
				fish = append(fish, 8)
			}

			fish[j]--
		}
		// fmt.Printf("DAY %v: %v\n", i+1, fish)
	}

	return len(fish)
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)
	fish := utils.AtoiSlc(strings.Split(lines[0], ","))

	DAYS := 256

	// Calculate a profile of how many fish reproduce every day
	profile := make([]int, DAYS)
	reproduce(0, DAYS, profile)
	fmt.Println(profile)

	var sum int
	for _, n := range fish {
		sum++ // Account for starting fish

		// The fish number is like an offset, so we just start a bit later depending on the number
		for day := n; day < DAYS; day++ {
			sum += profile[day]
		}
	}

	return sum
}

func reproduce(fish int, days int, profile []int) {
	days -= fish
	for ; days > 0; days -= 7 {
		// +1 fish on this day
		profile[days-1]++
		reproduce(9, days, profile)
	}
}

// Old function that took too long
func reproduceOld(fish int, days int) int {
	days -= fish
	sum := 1
	for ; days > 0; days -= 7 {
		sum += reproduceOld(9, days)
	}
	return sum
}
