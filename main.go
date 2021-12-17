package main

import (
	"flag"
	"fmt"
	"github.com/Olaroll/adventofcode21/day17"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "input.txt", "Sets the filename that's used as input")
	flag.Parse()

	fmt.Println(day17.Solve2(file))
}
