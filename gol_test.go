package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeighbourCountCenter(t *testing.T) {

	world := NewWorld(3, 3)

	world.area[0][0] = true
	world.area[0][1] = true
	world.area[0][2] = true

	world.area[1][0] = true
	world.area[1][2] = true

	world.area[2][0] = true
	world.area[2][1] = true
	world.area[2][2] = true

	assert.Equal(t, 8, world.neighbourCount(1, 1))
}

func TestNeighbourCountTopLeft(t *testing.T) {

	world := NewWorld(3, 3)

	world.area[0][0] = true // make sure this is not counted
	world.area[0][1] = true

	world.area[1][0] = true
	world.area[1][1] = true

	assert.Equal(t, 3, world.neighbourCount(0, 0))
}

func TestNeighbourCountBottomRight(t *testing.T) {

	world := NewWorld(3, 3)

	world.area[1][1] = true
	world.area[1][2] = true

	world.area[2][1] = true

	assert.Equal(t, 3, world.neighbourCount(2, 2))
}

// if correctly implemented, the "blinker" should toggle from | to -
func TestBlinker(t *testing.T) {

	world := NewWorld(5, 5)

	world.area = [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	world.Progress()
	assert.Equal(t, [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}, world.area)
}
