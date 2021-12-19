package day19

type Beacon struct {
	x, y, z int
}

func NewBeacon(x, y, z int) *Beacon {
	return &Beacon{x: x, y: y, z: z}
}

func (b *Beacon) RotateFromTo(from, to int) *Beacon {
	for to < from {
		to += 24
	}

	for i := from + 1; i <= to; i++ {
		b.RotateAt(i)
	}
	return b
}

func (b *Beacon) RotateAt(i int) *Beacon {
	if i%4 == 0 {
		b.Roll()
	} else {
		if i/4%2 == 0 {
			b.TurnCW()
		} else {
			b.TurnCCW()
		}
	}
	return b
}

func (b *Beacon) Roll() *Beacon {
	b.y, b.z = -b.z, b.y
	return b
}

func (b *Beacon) TurnCW() *Beacon {
	b.x, b.y = b.y, -b.x
	return b
}

func (b *Beacon) TurnCCW() *Beacon {
	b.x, b.y = -b.y, b.x
	return b
}
