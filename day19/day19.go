package day19

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
	"regexp"
)

var dir = "./day19/"

func getScanners(lines []string) map[int]Scanner {
	regID := regexp.MustCompile(`--- scanner (\d+) ---`)
	regBeacon := regexp.MustCompile(`([-\d]+),([-\d]+),([-\d]+)`)

	scanners := make(map[int]Scanner)

	for len(lines) >= 14 {
		id := utils.Atoi(regID.FindStringSubmatch(lines[0])[1])
		lines = lines[1:]
		scanner := NewScanner()

		for len(lines) > 0 {
			if lines[0] == "" {
				lines = lines[1:]
				break
			}

			coords := regBeacon.FindStringSubmatch(lines[0])

			beacon := NewBeacon(utils.Atoi(coords[1]), utils.Atoi(coords[2]), utils.Atoi(coords[3]))
			scanner = scanner.AddBeacon(beacon)

			lines = lines[1:]
		}

		scanners[id] = scanner
	}

	return scanners
}

type Empty struct{}

func solve(file string) int {
	lines := utils.GetLines(dir + file)

	scanners := getScanners(lines)

	chart := NewChart()

	fmt.Println("Adding initial scanner")
	chart.AddScanner(scanners[0])
	delete(scanners, 0)

	for len(scanners) > 0 {
		for key, scanner := range scanners {
			fmt.Println("Trying scanner ", key)
			if chart.TryAllRotations(scanner) {
				chart.AddScanner(scanner)
				delete(scanners, key)
			}
		}
	}

	return chart.Count()
}

func Solve1(file string) int {
	return solve(file)
}

var foundScanners = [][3]int{{0, 0, 0}}

func Solve2(file string) int {
	solve(file)

	var max int
	for _, a := range foundScanners {
		for _, b := range foundScanners {
			if a == b {
				continue
			}

			d := abs(a[0]-b[0]) + abs(a[1]-b[1]) + abs(a[2]-b[2])
			if d > max {
				max = d
			}
		}
	}

	return max
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
