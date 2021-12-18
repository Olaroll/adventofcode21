package day18

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/day18/tree"
	"github.com/Olaroll/adventofcode21/utils"
)

var dir = "./day18/"

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	t := tree.FromString(lines[0])
	t.Reduce()
	lines = lines[1:]
	for _, line := range lines {
		t2 := tree.FromString(line)
		t2.Reduce()
		t = t.Add(t2)
	}

	fmt.Printf("Final snailfish number: %v\n", t.String())

	return t.Magnitude()
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	var max int
	for _, line := range lines {
		for _, line2 := range lines {
			if line == line2 {
				continue
			}

			t1 := tree.FromString(line)
			t1.Reduce()

			t2 := tree.FromString(line2)
			t2.Reduce()

			added := t1.Add(t2)
			mag := added.Magnitude()
			fmt.Printf("Magnitude %v for %v\n", mag, added.String())
			if mag > max {
				max = mag
			}
		}
	}

	return max
}
