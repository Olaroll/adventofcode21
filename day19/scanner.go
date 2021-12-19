package day19

type Scanner []*Beacon

func NewScanner() Scanner {
	return make(Scanner, 0)
}

func (s Scanner) AddBeacon(b *Beacon) Scanner {
	return append(s, b)
}

func (s Scanner) RotateFromTo(from, to int) {
	for _, beacon := range s {
		beacon.RotateFromTo(from, to)
	}
}

func (s Scanner) ApplyOffset(x, y, z int) {
	for _, beacon := range s {
		beacon.x += x
		beacon.y += y
		beacon.z += z
	}
}
