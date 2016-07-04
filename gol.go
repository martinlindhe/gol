package gol

import (
	"math/rand"
	"strings"
	"time"
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	rnd        = rand.New(randSource)
)

// Progress game for each cell
func Progress(a [][]bool) [][]bool {

	height := len(a)
	width := len(a[0])

	next := MakeTwoDimensionalBoolSlice(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			pop := neighbourCount(a, x, y)
			switch {
			case pop < 2:
				// rule 1. Any live cell with fewer than two live neighbours dies,
				// as if caused by under-population.
				next[y][x] = false

			case (pop == 2 || pop == 3) && a[y][x]:
				// rule 2. Any live cell with two or three live neighbours
				// lives on to the next generation.
				next[y][x] = true

			case pop > 3:
				// rule 3. Any live cell with more than three live neighbours dies,
				// as if by over-population.
				next[y][x] = false

			case pop == 3:
				// rule 4. Any dead cell with exactly three live neighbours becomes
				// a live cell, as if by reproduction.
				next[y][x] = true
			}
		}
	}

	return next
}

// PrettyPrint prints a in a pretty fashion
func PrettyPrint(a [][]bool) string {

	height := len(a)
	width := len(a[0])

	res := []string{}

	for y := 0; y < height; y++ {

		line := ""
		for x := 0; x < width; x++ {
			if a[y][x] {
				line += "*"
			} else {
				line += " "
			}
		}
		res = append(res, line)
	}

	return strings.Join(res, "\n")
}

// RandomSeed inits a with limit values at random places
func RandomSeed(a [][]bool, limit int) [][]bool {

	height := len(a)
	width := len(a[0])

	for i := 0; i < limit; i++ {
		w := rnd.Intn(width)
		h := rnd.Intn(height)
		a[h][w] = true
	}

	return a
}

// MakeTwoDimensionalBoolSlice creates b[height][width]
func MakeTwoDimensionalBoolSlice(width, height int) [][]bool {

	a := make([][]bool, height)
	for i := 0; i < height; i++ {
		a[i] = make([]bool, width)
	}
	return a
}

// counts neughbours next to x,y (min 0, max 8)
func neighbourCount(a [][]bool, x, y int) int {

	height := len(a)
	width := len(a[0])

	lowX := 0
	if x > 0 {
		lowX = x - 1
	}
	lowY := 0
	if y > 0 {
		lowY = y - 1
	}

	highX := width - 1
	if x < width-1 {
		highX = x + 1
	}
	highY := height - 1
	if y < height-1 {
		highY = y + 1
	}

	// fmt.Println("-- X", width, lowX, highX)
	// fmt.Println("-- Y", height, lowY, highY)

	near := 0
	for pY := lowY; pY <= highY; pY++ {
		for pX := lowX; pX <= highX; pX++ {
			if pX == x && pY == y {

			} else if a[pY][pX] {
				near++
			}
		}
	}

	return near
}
