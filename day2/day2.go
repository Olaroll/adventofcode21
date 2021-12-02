package day2

import (
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

type Command struct {
	cmd   string
	count int
}

func Solve1() int {
	lines := utils.GetLines("./day2/input.txt")

	commands := make([]Command, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " ")

		if len(split) == 2 {
			commands = append(commands, Command{split[0], utils.Atoi(split[1])})
		}
	}

	var X int
	var Y int
	for _, command := range commands {
		switch command.cmd {
		case "forward":
			X += command.count
		case "up":
			Y -= command.count
		case "down":
			Y += command.count
		}
	}

	return X * Y
}

func Solve2() int {
	lines := utils.GetLines("./day2/input.txt")

	commands := make([]Command, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " ")

		if len(split) == 2 {
			commands = append(commands, Command{split[0], utils.Atoi(split[1])})
		}
	}

	var X int
	var Y int
	var aim int
	for _, command := range commands {
		switch command.cmd {
		case "forward":
			X += command.count
			Y += command.count * aim
		case "up":
			aim -= command.count
		case "down":
			aim += command.count
		}
	}

	return X * Y
}
