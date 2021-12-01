package day1

import "github.com/Olaroll/adventofcode21/utils"

func Solve1() int {
	nums := utils.GetLinesAsInts("./day1/input.txt")

	var prev int
	var count int
	for _, num := range nums {
		if prev < num && prev != 0 {
			count++
		}
		prev = num
	}

	return count
}

func Solve2() int {
	nums := utils.GetLinesAsInts("./day1/input.txt")

	windowLen := 3
	var prev []int
	var count int
	for i := 0; i <= len(nums)-windowLen; i++ {
		current := make([]int, 0, windowLen)
		for j := 0; j < windowLen; j++ {
			current = append(current, nums[i+j])
		}

		if compareWindow(current, prev) {
			count++
		}

		prev = current
	}

	return count
}

func compareWindow(current, prev []int) bool {
	if prev == nil {
		return false
	}
	return addNums(current...) > addNums(prev...)
}

func addNums(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
