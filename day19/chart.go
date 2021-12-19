package day19

import "fmt"

type Chart map[Beacon]Empty

func NewChart() Chart {
	return make(map[Beacon]Empty)
}

func (c Chart) Count() int {
	return len(c)
}

func (c Chart) AddScanner(scanner Scanner) {
	for _, beacon := range scanner {
		c[*beacon] = Empty{}
	}
}

func (c Chart) TryAllRotations(scanner Scanner) bool {
	for i := 0; i < 24; i++ {
		if c.TryAllPositions(scanner) {
			return true
		}
		scanner.RotateFromTo(i, i+1)
	}
	return false
}

func (c Chart) TryAllPositions(scanner Scanner) bool {
	for i := 0; i < len(scanner)-12; i++ {
		testBcn := *scanner[i]

		for setBcn := range c {
			x := setBcn.x - testBcn.x
			y := setBcn.y - testBcn.y
			z := setBcn.z - testBcn.z

			if c.CountOverlap(x, y, z, scanner) {
				scanner.ApplyOffset(x, y, z)
				fmt.Printf("Success!! Found at X:%v Y:%v Z:%v\n", x, y, z)
				foundScanners = append(foundScanners, [3]int{x, y, z})
				return true
			}
		}

	}
	return false
}

func (c Chart) CountOverlap(x, y, z int, scanner Scanner) bool {
	var count int
	for i := 0; i < len(scanner)-(12-count)+1; i++ {
		beacon := *scanner[i]
		beacon.x += x
		beacon.y += y
		beacon.z += z

		_, overlap := c[beacon]
		if overlap {
			count++
			if count == 12 {
				return true
			}
		}
	}
	return false
}
