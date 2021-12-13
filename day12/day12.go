package day12

import (
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
	"unicode"
)

var dir = "./day12/"

var (
	paths [][]string
	rooms map[string][]string
)

func pathfind(visited []string, anyLower bool) {
	lowerVisited := removeUpper(visited)
	for _, nextRoom := range rooms[visited[len(visited)-1]] {
		if nextRoom == "end" {
			// If current room is next to end, don't go any further
			visited = append(visited, "end")
			paths = append(paths, visited)
			continue
		}

		nextAnyLower := anyLower
		if isInSet(nextRoom, lowerVisited) {
			if anyLower && nextRoom != "start" {
				nextAnyLower = false
			} else {
				continue
			}
		}

		// Next recursion level with visited list + current room
		newVisited := make([]string, len(visited))
		copy(newVisited, visited)
		newVisited = append(newVisited, nextRoom)
		pathfind(newVisited, nextAnyLower)
	}
}

func isInSet(target string, set []string) bool {
	for _, check := range set {
		if check == target {
			return true
		}
	}
	return false
}

func removeUpper(set []string) []string {
	// Remove big rooms
	lower := make([]string, 0, len(set))
	for _, room := range set {
		if unicode.IsLower(rune(room[0])) {
			lower = append(lower, room)
		}
	}
	return lower
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	rooms = make(map[string][]string)

	for _, line := range lines {
		conn := strings.Split(line, "-")

		rooms[conn[0]] = append(rooms[conn[0]], conn[1])
		rooms[conn[1]] = append(rooms[conn[1]], conn[0])
	}

	pathfind([]string{"start"}, false)

	return len(paths)
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	rooms = make(map[string][]string)

	for _, line := range lines {
		conn := strings.Split(line, "-")

		rooms[conn[0]] = append(rooms[conn[0]], conn[1])
		rooms[conn[1]] = append(rooms[conn[1]], conn[0])
	}

	pathfind([]string{"start"}, true)

	return len(paths)
}
