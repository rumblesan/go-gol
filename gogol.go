package main

import "fmt"
import "strconv"

type coord struct {
	x int
	y int
}

func (c coord) neighbours() []coord {

	n := make([]coord, 0)

	for _, xVal := range []int{-1, 0, 1} {
		for _, yVal := range []int{-1, 0, 1} {
			if xVal == 0 && yVal == 0 {
				continue
			}
			n = append(n, coord{c.x + xVal, c.y + yVal})
		}
	}
	return n
}

type board struct {
	cells map[coord]bool
}

func createBoard(coords ...coord) board {

	cells := make(map[coord]bool)

	for _, c := range coords {
		cells[c] = true
	}
	return board{cells}
}

func (b *board) cellHealth(c coord) bool {
	return b.cells[c]
}

func (b board) String() string {

	output := "Board\n"

	xMax := 0
	xMin := 0
	yMax := 0
	yMin := 0

	for c := range b.cells {
		if c.x < xMin {
			xMin = c.x
		}
		if c.x > xMax {
			xMax = c.x
		}
		if c.y < yMin {
			yMin = c.y
		}
		if c.y > yMax {
			yMax = c.y
		}
	}

	output += "  "
	for x := xMin; x < xMax; x += 1 {
		if x < 0 {
			output += "|"
		} else {
			output += " "
		}
	}
	output += "\n  "
	for x := xMin; x <= xMax; x += 1 {
		if x < 0 {
			output += strconv.Itoa(x * -1)
		} else {
			output += strconv.Itoa(x)
		}
	}
	output += "\n"

	for y := yMin; y <= yMax; y += 1 {
		if y < 0 {
			output += strconv.Itoa(y)
		} else {
			output += " " + strconv.Itoa(y)
		}
		for x := xMin; x <= xMax; x += 1 {
			if b.cellHealth(coord{x, y}) {
				output += "X"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output

}

func nextGen(b *board) board {

	toCheck := make(map[coord]bool)

	nextCells := make(map[coord]bool)

	for k := range b.cells {
		for _, c := range k.neighbours() {
			toCheck[c] = true
			toCheck[k] = true
		}
	}

	for k := range toCheck {
		if cellLife(b, b.cellHealth(k), k.neighbours()) {
			nextCells[k] = true
		}
	}
	return board{nextCells}
}

func cellLife(b *board, alive bool, neighbours []coord) bool {
	a := 0

	for _, c := range neighbours {
		if b.cells[c] {
			a += 1
		}
	}

	if alive {
		switch a {
		case 0, 1:
			return false
		case 2, 3:
			return true
		default:
			return false
		}
	} else {
		switch a {
		case 3:
			return true
		default:
			return false
		}
	}
}

func main() {

	fmt.Println("Conway's Game of Life in Go")

	b := createBoard(
		coord{1, 0},
		coord{-3, -2},
		coord{0, 1},
		coord{2, 1},
		coord{0, 2},
		coord{2, 2},
		coord{4, 2},
		coord{3, 5},
		coord{2, 5},
		coord{5, 5},
		coord{1, 3})

	fmt.Println("first gen", b)

	generations := 15

	for i := 0; i < generations; i += 1 {
		b = nextGen(&b)
		fmt.Println(b)
	}

	fmt.Println(b)

}
