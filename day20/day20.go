package day20

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
)

var dir = "./day20/"

type Image [][]byte

func (img Image) Count(char byte) int {
	var count int
	for y := range img {
		for x := range img[y] {
			if img[y][x] == char {
				count++
			}
		}
	}
	return count
}

func (img Image) expand(plusBorder int) Image {
	newImg := make(Image, len(img)+plusBorder*2)
	for y := range newImg {
		newImg[y] = make([]byte, len(img[0])+plusBorder*2)
	}
	return newImg
}

func (img Image) ApplyAlgo(algo string, flipNA bool) Image {
	newImg := img.expand(1)
	for y := range newImg {
		for x := range newImg[y] {
			i := img.XYToIndex(x-1, y-1, flipNA)
			newImg[y][x] = algo[i]
		}
	}
	return newImg
}

func (img Image) XYToIndex(x, y int, flipNA bool) int {
	var res int
	res |= img.getBit(x-1, y-1, 8, flipNA)
	res |= img.getBit(x, y-1, 7, flipNA)
	res |= img.getBit(x+1, y-1, 6, flipNA)
	res |= img.getBit(x-1, y, 5, flipNA)
	res |= img.getBit(x, y, 4, flipNA)
	res |= img.getBit(x+1, y, 3, flipNA)
	res |= img.getBit(x-1, y+1, 2, flipNA)
	res |= img.getBit(x, y+1, 1, flipNA)
	res |= img.getBit(x+1, y+1, 0, flipNA)
	return res
}

func (img Image) getBit(x, y, bitIndex int, flipNA bool) int {
	if !img.isValid(x, y) {
		if flipNA {
			return 1 << bitIndex
		}
		return 0
	}

	if img[y][x] == '.' {
		return 0
	}
	return 1 << bitIndex
}

func (img Image) isValid(x, y int) bool {
	if y < 0 || x < 0 || y >= len(img) || x >= len(img[y]) {
		return false
	}
	return true
}

func (img Image) print() {
	for _, line := range img {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)
	algo := lines[0]
	lines = lines[2:]

	img := make(Image, len(lines))
	for i := range lines {
		img[i] = []byte(lines[i])
	}

	doFlip := algo[0] == '#'

	for i := 0; i < 2; i++ {
		img = img.ApplyAlgo(algo, doFlip && i%2 == 1)
	}

	return img.Count('#')
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)
	algo := lines[0]
	lines = lines[2:]

	img := make(Image, len(lines))
	for i := range lines {
		img[i] = []byte(lines[i])
	}

	doFlip := algo[0] == '#'

	for i := 0; i < 50; i++ {
		img = img.ApplyAlgo(algo, doFlip && i%2 == 1)
	}

	return img.Count('#')
}
