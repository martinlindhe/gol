package gol

import (
	"image"
	"math/rand"
	"time"
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	rnd        = rand.New(randSource)
)

// World represents the game state
type World struct {
	area   [][]bool
	width  int
	height int
}

// NewWorld creates a new world
func NewWorld(width, height int) *World {

	world := World{}
	world.area = makeArea(width, height)
	world.width = width
	world.height = height
	return &world
}

// RandomSeed inits world with a random state
func (w *World) RandomSeed(limit int) {

	for i := 0; i < limit; i++ {
		x := rnd.Intn(w.width)
		y := rnd.Intn(w.height)
		w.area[y][x] = true
	}
}

// Progress game state by one tick
func (w *World) Progress() {

	next := makeArea(w.width, w.height)

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {

			pop := w.neighbourCount(x, y)
			switch {
			case pop < 2:
				// rule 1. Any live cell with fewer than two live neighbours
				// dies, as if caused by under-population.
				next[y][x] = false

			case (pop == 2 || pop == 3) && w.area[y][x]:
				// rule 2. Any live cell with two or three live neighbours
				// lives on to the next generation.
				next[y][x] = true

			case pop > 3:
				// rule 3. Any live cell with more than three live neighbours
				// dies, as if by over-population.
				next[y][x] = false

			case pop == 3:
				// rule 4. Any dead cell with exactly three live neighbours
				// becomes a live cell, as if by reproduction.
				next[y][x] = true
			}
		}
	}
	w.area = next
}

// DrawImage paints current game state
func (w *World) DrawImage(img *image.RGBA) {

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			pos := 4*y*w.width + 4*x
			if w.area[y][x] {
				img.Pix[pos] = 0xff
				img.Pix[pos+1] = 0xff
				img.Pix[pos+2] = 0xff
				img.Pix[pos+3] = 0xff
			} else {
				img.Pix[pos] = 0
				img.Pix[pos+1] = 0
				img.Pix[pos+2] = 0
				img.Pix[pos+3] = 0
			}
		}
	}
}

// calculates the Moore neighborhood of x, y
func (w *World) neighbourCount(x, y int) int {

	lowX := 0
	if x > 0 {
		lowX = x - 1
	}

	lowY := 0
	if y > 0 {
		lowY = y - 1
	}

	highX := w.width - 1
	if x < w.width-1 {
		highX = x + 1
	}

	highY := w.height - 1
	if y < w.height-1 {
		highY = y + 1
	}

	near := 0
	for pY := lowY; pY <= highY; pY++ {
		for pX := lowX; pX <= highX; pX++ {
			if !(pX == x && pY == y) && w.area[pY][pX] {
				near++
			}
		}
	}

	return near
}

func makeArea(width, height int) [][]bool {

	area := make([][]bool, height)
	for i := 0; i < height; i++ {
		area[i] = make([]bool, width)
	}
	return area
}
