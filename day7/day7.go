package day7

import (
	"github.com/Olaroll/adventofcode21/utils"
	"math"
	"strings"
)

func Solve1() int {
	nums := utils.AtoiSlc(strings.Split(utils.GetLines("./day7/input.txt")[0], ","))

	var max int
	var min int
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	//var best int
	bestSum := math.MaxInt
	for i := min; i <= max; i++ {
		var sum int
		for _, n := range nums {
			sum += int(math.Abs(float64(n - i)))
		}

		if sum < bestSum {
			bestSum = sum
			//best = i
		}
	}

	return bestSum
}

func Solve2() int {
	nums := utils.AtoiSlc(strings.Split(utils.GetLines("./day7/input.txt")[0], ","))

	var max int
	var min int
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	//var best int
	bestSum := math.MaxInt
	for i := min; i <= max; i++ {
		var sum int
		for _, n := range nums {
			d := int(math.Abs(float64(n - i)))
			sum += d * (d + 1) / 2
		}

		if sum < bestSum {
			bestSum = sum
			//best = i
		}
	}

	return bestSum
}
