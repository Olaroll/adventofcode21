package tree

import (
	"fmt"
	"github.com/Olaroll/adventofcode21/utils"
	"strings"
	"unicode"
)

type Tree struct {
	left   *Tree
	right  *Tree
	parent *Tree

	value int
}

func FromString(str string) *Tree {
	reader := strings.NewReader(str)

	return create(reader)
}

func create(reader *strings.Reader) *Tree {
	if read(reader) != '[' {
		return nil
	}

	tree := &Tree{}
	for {
		b := peek(reader)

		switch {
		case b == '[':
			tree.SetEither(create(reader))

		case unicode.IsDigit(rune(b)):
			var numBytes []byte
			for unicode.IsDigit(rune(peek(reader))) {
				numBytes = append(numBytes, read(reader))
			}

			num := utils.Atoi(string(numBytes))
			tree.SetEither(&Tree{value: num})

		case b == ',':
			skip(reader)

		case b == ']':
			skip(reader)
			return tree

		default:
			panic(fmt.Errorf("unknown character (%v)", b))
		}
	}
}

func peek(reader *strings.Reader) byte {
	b := read(reader)
	_, _ = reader.Seek(-1, 1)
	return b
}

func read(reader *strings.Reader) byte {
	b, err := reader.ReadByte()
	if err != nil {
		return 0
	}
	return b
}

func skip(reader *strings.Reader) {
	_, _ = reader.Seek(1, 1)
}

func (t *Tree) SetEither(t2 *Tree) {
	switch {
	case t.left == nil:
		t.SetLeft(t2)
	case t.right == nil:
		t.SetRight(t2)
	default:
		panic("Left and right subtrees already set")
	}
}

func (t *Tree) SetLeft(t2 *Tree) {
	t.left = t2
	t2.parent = t
}

func (t *Tree) SetRight(t2 *Tree) {
	t.right = t2
	t2.parent = t
}

func (t *Tree) Add(t2 *Tree) *Tree {
	outer := &Tree{left: t, right: t2}
	t.parent, t2.parent = outer, outer
	outer.Reduce()
	return outer
}

func (t *Tree) IsValue() bool {
	if t.left == nil && t.right == nil {
		return true
	}
	return false
}

func (t *Tree) SetZero() {
	t.left = nil
	t.right = nil
	t.value = 0
}

func (t *Tree) Value() int {
	if t.IsValue() {
		return t.value
	}

	var total int
	if t.left != nil {
		total += t.left.Value()
	}

	if t.right != nil {
		total += t.right.Value()
	}

	return total
}

func (t *Tree) String() string {
	if t == nil {
		return "[]"
	}
	if t.IsValue() {
		return fmt.Sprint(t.value)
	} else {
		s := ""
		if t.left != nil {
			s += t.left.String() + ","
		}
		if t.right != nil {
			s += t.right.String()
		}
		return "[" + s + "]"
	}
}

func (t *Tree) Magnitude() int {
	if t.IsValue() {
		return t.value
	}

	var total int
	if t.left != nil {
		total += 3 * t.left.Magnitude()
	}

	if t.right != nil {
		total += 2 * t.right.Magnitude()
	}

	return total
}

// Reducing stuff

func (t *Tree) Reduce() {
	for {
		found := t.findDeeperThan(4)
		if found != nil {
			found.explode()
			continue
		}

		found = t.findGreaterThan(9)
		if found != nil {
			found.split()
			continue
		}

		break
	}
}

func (t *Tree) findDeeperThan(target int) *Tree {
	return t.fdt(target, 1)
}

func (t *Tree) fdt(target, depth int) *Tree {
	if depth > target && !t.IsValue() {
		return t
	}

	if t.left != nil {
		found := t.left.fdt(target, depth+1)
		if found != nil {
			return found
		}
	}

	if t.right != nil {
		found := t.right.fdt(target, depth+1)
		if found != nil {
			return found
		}
	}

	return nil
}

func (t *Tree) findGreaterThan(target int) *Tree {
	if t.IsValue() && t.value > target {
		return t
	}

	if t.left != nil {
		found := t.left.findGreaterThan(target)
		if found != nil {
			return found
		}
	}

	if t.right != nil {
		found := t.right.findGreaterThan(target)
		if found != nil {
			return found
		}
	}

	return nil
}

// Exploding

func (t *Tree) explode() {
	var left int
	if t.left != nil {
		left = t.left.Value()
	}
	t.explodeLeft(left)

	var right int
	if t.right != nil {
		right = t.right.Value()
	}
	t.explodeRight(right)

	t.SetZero()
}

func (t *Tree) explodeLeft(num int) {
	if t.parent != nil {
		if t.parent.left == t {
			t.parent.explodeLeft(num)
			return
		} else {
			t.parent.left.explodeLeftDown(num)
		}
	}
}

func (t *Tree) explodeLeftDown(num int) {
	if t.IsValue() {
		t.value += num
	} else if t.right != nil {
		t.right.explodeLeftDown(num)
	}
}

func (t *Tree) explodeRight(num int) {
	if t.parent != nil {
		if t.parent.right == t {
			t.parent.explodeRight(num)
			return
		} else {
			t.parent.right.explodeRightDown(num)
		}
	}
}

func (t *Tree) explodeRightDown(num int) {
	if t.IsValue() {
		t.value += num
	} else if t.left != nil {
		t.left.explodeRightDown(num)
	}
}

// Splitting

func (t *Tree) split() {
	if t.IsValue() {
		left := t.value / 2
		right := t.value/2 + t.value%2
		t.value = 0

		t.SetRight(&Tree{value: right})
		t.SetLeft(&Tree{value: left})
	}
}
