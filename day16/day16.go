package day16

import (
	"encoding/hex"
	"github.com/Olaroll/adventofcode21/utils"
)

var dir = "./day16/"

var versionSum int

func Solve1(file string) int {
	line := utils.GetLines(dir + file)[0]

	hexBytes, _ := hex.DecodeString(line)
	reader := &BitReader{data: hexBytes}
	readPacket(reader)

	return versionSum
}

func Solve2(file string) int {
	line := utils.GetLines(dir + file)[0]

	hexBytes, _ := hex.DecodeString(line)
	reader := &BitReader{data: hexBytes}

	return readPacket(reader)
}
