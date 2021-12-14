package day03

import (
	"github.com/Olaroll/adventofcode21/utils"
	"log"
	"strconv"
)

var dir = "./day03/"

// Not too happy with today's solutions and how long they took,
// but at least I got it done in the end

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	var binlines []int
	for _, line := range lines {
		bin, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			log.Fatalln(err)
		}

		binlines = append(binlines, int(bin))
	}

	binLen := 12
	binCounts := make([]int, binLen)

	for i := 0; i < binLen; i++ {
		for _, line := range binlines {
			if 1<<i&line > 0 {
				binCounts[i]++
			}
		}
	}

	var gamma int
	var epsilon int
	for i, v := range binCounts {
		len2 := len(binlines) / 2
		if v > len2 {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}

	return gamma * epsilon
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	oxyList := make([]string, len(lines))
	copy(oxyList, lines)
	co2List := make([]string, len(lines))
	copy(co2List, lines)

	oxyStr := doNarrowing(oxyList, false)
	co2Str := doNarrowing(co2List, true)

	oxy, _ := strconv.ParseInt(oxyStr, 2, 0)
	co2, _ := strconv.ParseInt(co2Str, 2, 0)

	return int(oxy * co2)
}

func doNarrowing(nums []string, invert bool) string {
	for i := 0; i < 12 && len(nums) > 1; i++ {
		nums = narrowList(nums, i, invert)
	}

	if len(nums) == 1 {
		return nums[0]
	}

	log.Fatalln("Couldn't narrow down to 1")
	return ""
}

func countOnes(nums []string, index int) int {
	var count int
	for _, num := range nums {
		if num[index] == '1' {
			count++
		}
	}
	return count
}

func narrowList(nums []string, index int, invert bool) (narrowed []string) {
	ones := countOnes(nums, index)
	passOne := len(nums)-ones <= len(nums)/2
	passOne = passOne != invert

	for _, num := range nums {
		isOne := num[index] == '1'

		if isOne == passOne {
			narrowed = append(narrowed, num)
		}
	}

	return narrowed
}
