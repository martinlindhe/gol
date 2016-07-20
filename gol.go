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
			w.drawPixel(img, x, y)
		}
	}
}

func (w *World) drawPixel(img *image.RGBA, x, y int) {
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

// calculates the Moore neighborhood of x, y
func (w *World) neighbourCount(x, y int) int {

	norm := w.normalizeCoords(x, y)
	near := 0
	for pY := norm.lowY; pY <= norm.highY; pY++ {
		for pX := norm.lowX; pX <= norm.highX; pX++ {
			if !(pX == x && pY == y) && w.area[pY][pX] {
				near++
			}
		}
	}
	return near
}

type normalizedCoords struct {
	lowX  int
	lowY  int
	highX int
	highY int
}

func (w *World) normalizeCoords(x, y int) normalizedCoords {

	r := normalizedCoords{}
	r.highX = w.width - 1
	r.highY = w.height - 1
	if x > 0 {
		r.lowX = x - 1
	}
	if y > 0 {
		r.lowY = y - 1
	}
	if x < w.width-1 {
		r.highX = x + 1
	}
	if y < w.height-1 {
		r.highY = y + 1
	}
	return r
}

func makeArea(width, height int) [][]bool {

	area := make([][]bool, height)
	for i := 0; i < height; i++ {
		area[i] = make([]bool, width)
	}
	return area
}
