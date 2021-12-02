package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetLines(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read file %v\n", path)
	}

	lines := strings.Split(strings.ReplaceAll(string(bytes), "\r\n", "\n"), "\n")

	if len(lines) == 0 {
		log.Fatalf("no lines found in %v\n", path)
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func GetLinesAsInts(path string) []int {
	lines := GetLines(path)

	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		num := Atoi(line)
		nums = append(nums, num)
	}

	return nums
}

func Atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(fmt.Errorf("could not convert %v to number", str))
	}
	return num
}
