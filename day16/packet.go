package day16

import (
	"math"
	"strconv"
	"strings"
)

type operatorFunc func(n int) int

var operators = map[int]func() operatorFunc{
	0: sum,
	1: product,
	2: minimum,
	3: maximum,

	5: greater,
	6: less,
	7: equal,
}

func readPacket(reader *BitReader) int {
	version := reader.ReadInt(3)
	versionSum += version

	var ret int
	if typeID := reader.ReadInt(3); typeID == 4 {
		ret = literal(reader)
	} else {
		ret = operator(reader, operators[typeID]())
	}
	return ret
}

func literal(reader *BitReader) int {
	var builder strings.Builder
	for {
		stop := reader.Read(1) == "0"

		builder.WriteString(reader.Read(4))

		if stop {
			break
		}
	}

	num, _ := strconv.ParseInt(builder.String(), 2, 0)
	return int(num)
}

func operator(reader *BitReader, opfunc operatorFunc) int {
	lenType := reader.ReadInt(1)

	var ret int
	switch lenType {
	case 0:
		limit := reader.ReadInt(15)

		initial := reader.BitsRead()
		for reader.BitsRead() < initial+limit {
			ret = opfunc(readPacket(reader))
		}
	case 1:
		limit := reader.ReadInt(11)

		for i := 0; i < limit; i++ {
			ret = opfunc(readPacket(reader))
		}
	}

	return ret
}

func sum() operatorFunc {
	var total int
	return func(n int) int {
		total += n
		return total
	}
}

func product() operatorFunc {
	var total int = 1
	return func(n int) int {
		total *= n
		return total
	}
}

func minimum() operatorFunc {
	min := math.MaxInt
	return func(n int) int {
		if n < min {
			min = n
		}
		return min
	}
}

func maximum() operatorFunc {
	max := 0
	return func(n int) int {
		if n > max {
			max = n
		}
		return max
	}
}

func greater() operatorFunc {
	var first int
	var firstSet bool
	return func(n int) int {
		if !firstSet {
			first = n
			firstSet = true
		}

		var ret int
		if first > n {
			ret = 1
		}

		return ret
	}
}

func less() operatorFunc {
	var first int
	var firstSet bool
	return func(n int) int {
		if !firstSet {
			first = n
			firstSet = true
		}

		var ret int
		if first < n {
			ret = 1
		}

		return ret
	}
}

func equal() operatorFunc {
	var first int
	var firstSet bool
	return func(n int) int {
		if !firstSet {
			first = n
			firstSet = true
		}

		var ret int
		if first == n {
			ret = 1
		}

		return ret
	}
}
