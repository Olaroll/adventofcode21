package day02

import (
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
)

var dir = "./day02/"

type Command struct {
	cmd   string
	count int
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	commands := make([]Command, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " ")

		command := Command{split[0], utils.Atoi(split[1])}
		commands = append(commands, command)
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

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	commands := make([]Command, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " ")

		command := Command{split[0], utils.Atoi(split[1])}
		commands = append(commands, command)
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
