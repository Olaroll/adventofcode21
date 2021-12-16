package main

import (
	"flag"
	"fmt"
	"github.com/Olaroll/adventofcode21/day16"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "input.txt", "Sets the filename that's used as input")
	flag.Parse()

	fmt.Println(day16.Solve2(file))
}
