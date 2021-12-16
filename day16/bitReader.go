package day16

import (
	"fmt"
	"strconv"
	"strings"
)

type BitReader struct {
	data []byte
	i    int
}

func (reader *BitReader) Len() int {
	return len(reader.data) * 8
}

func (reader *BitReader) BitsRead() int {
	return reader.i
}

func (reader *BitReader) ReadInt(n int) int {
	ret, _ := strconv.ParseInt(reader.Read(n), 2, 0)
	return int(ret)
}

func (reader *BitReader) Read(n int) string {
	if n == 0 {
		return ""
	}

	lastByte := (reader.i+n-1)/8 + 1
	if lastByte > len(reader.data) {
		panic(fmt.Errorf("trying to read more bits than available in reader"))
	}

	relevantBytes := reader.data[reader.i/8 : lastByte]
	str := bytesToString(relevantBytes)

	start := reader.i % 8
	end := (reader.i+n-1)%8 + 1

	reader.i += n

	return str[start : len(str)-8+end]
}

func bytesToString(bytes []byte) string {
	var builder strings.Builder
	builder.Grow(len(bytes) * 8)
	for i := 0; i < len(bytes); i++ {
		for j := 0; j < 8; j++ {
			zeroOrOne := bytes[i] >> (7 - j) & 1
			builder.WriteByte('0' + zeroOrOne)
		}
	}
	return builder.String()
}
