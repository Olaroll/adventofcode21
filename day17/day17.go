package day17

import (
	"github.com/Olaroll/adventofcode21/utils"
	"regexp"
)

var dir = "./day17/"

func Solve1(file string) int {
	line := utils.GetLines(dir + file)[0]
	r := regexp.MustCompile(`[xy]=([\d-]*)\.\.([\d-]*)`)
	found := r.FindAllStringSubmatch(line, 2)

	// xlow := found[0][0]
	// xhigh := found[0][1]
	ylow := -utils.Atoi(found[1][1]) - 1
	// yhigh := utils.Atoi(found[1][2])

	return ylow * (ylow + 1) / 2
}

type Target struct {
	xlow  int
	xhigh int
	ylow  int
	yhigh int
}

func (t Target) simulate(xvel, yvel int) bool {
	var x, y int

	for y >= t.ylow && x <= t.xhigh {
		if y <= t.yhigh && x >= t.xlow {
			return true
		}
		x += xvel
		y += yvel

		if xvel > 0 {
			xvel--
		}
		yvel--
	}

	return false
}

func Solve2(file string) int {
	line := utils.GetLines(dir + file)[0]
	r := regexp.MustCompile(`[xy]=([\d-]*)\.\.([\d-]*)`)
	found := r.FindAllStringSubmatch(line, 2)

	target := Target{}
	target.xlow = utils.Atoi(found[0][1])
	target.xhigh = utils.Atoi(found[0][2])
	target.ylow = utils.Atoi(found[1][1])
	target.yhigh = utils.Atoi(found[1][2])

	var count int
	for x := 0; x <= target.xhigh; x++ {
		for y := -target.ylow - 1; y >= target.ylow; y-- {
			if target.simulate(x, y) {
				count++
			}
		}
	}

	return count
}
