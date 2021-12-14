package day08

import (
	"github.com/Olaroll/adventofcode21/utils"
	"log"
	"strings"
)

var dir = "./day08/"

// Very iffy solution today, dunno why I didn't do the proper way >_>

const (
	n0 = int64('a') * 'b' * 'c' * 'e' * 'f' * 'g'
	n1 = int64('c') * 'f'
	n2 = int64('a') * 'c' * 'd' * 'e' * 'g'
	n3 = int64('a') * 'c' * 'd' * 'f' * 'g'
	n4 = int64('b') * 'c' * 'd' * 'f'
	n5 = int64('a') * 'b' * 'd' * 'f' * 'g'
	n6 = int64('a') * 'b' * 'd' * 'e' * 'f' * 'g'
	n7 = int64('a') * 'c' * 'f'
	n8 = int64('a') * 'b' * 'c' * 'd' * 'e' * 'f' * 'g'
	n9 = int64('a') * 'b' * 'c' * 'd' * 'f' * 'g'
)

func getProduct(str string, cipher map[rune]rune) int64 {
	var product int64 = 1
	for _, r := range str {
		product *= int64(cipher[r])
	}
	return product
}

func decode(str string, cipher map[rune]rune) int {
	switch getProduct(str, cipher) {
	case n0:
		return 0
	case n1:
		return 1
	case n2:
		return 2
	case n3:
		return 3
	case n4:
		return 4
	case n5:
		return 5
	case n6:
		return 6
	case n7:
		return 7
	case n8:
		return 8
	case n9:
		return 9
	}
	return -1
}

func decypher(hints []string) map[rune]rune {
	segments := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	gen := utils.NewPermutationGenerator(segments)

	cipher := make(map[rune]rune)
next:
	for {
		perm := gen()
		if perm == nil {
			return nil
		}

		for i, k := range segments {
			cipher[k] = perm[i].(rune)
		}

		for _, str := range hints {
			if decode(str, cipher) == -1 {
				continue next
			}
		}
		return cipher
	}
}

func Solve1(file string) int {
	lines := utils.GetLines(dir + file)

	var count int
	for _, line := range lines {
		split := strings.Split(line, "|")

		hints := strings.Fields(split[0])
		final := strings.Fields(split[1])

		cipher := decypher(hints)
		if cipher == nil {
			log.Fatalln("Could not find valid cipher")
		}

		for _, v := range final {
			n := decode(v, cipher)

			if n == 1 || n == 4 || n == 7 || n == 8 {
				count++
			}
			// result *= 10
			// result += n
		}
	}

	return count
}

func Solve2(file string) int {
	lines := utils.GetLines(dir + file)

	var sum int
	for _, line := range lines {
		split := strings.Split(line, "|")

		hints := strings.Fields(split[0])
		final := strings.Fields(split[1])

		cipher := decypher(hints)
		if cipher == nil {
			log.Fatalln("Could not find valid cipher")
		}

		var result int
		for _, v := range final {
			n := decode(v, cipher)

			result *= 10
			result += n
		}

		sum += result
	}

	return sum
}
