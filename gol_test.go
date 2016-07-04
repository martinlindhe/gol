package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeighbourCountCenter(t *testing.T) {

	area := MakeTwoDimensionalBoolSlice(3, 3)

	area[0][0] = true
	area[0][1] = true
	area[0][2] = true

	area[1][0] = true
	area[1][2] = true

	area[2][0] = true
	area[2][1] = true
	area[2][2] = true

	assert.Equal(t, 8, neighbourCount(area, 1, 1))
}

func TestNeighbourCountTopLeft(t *testing.T) {

	area := MakeTwoDimensionalBoolSlice(3, 3)

	area[0][0] = true // make sure this is not counted
	area[0][1] = true

	area[1][0] = true
	area[1][1] = true

	assert.Equal(t, 3, neighbourCount(area, 0, 0))
}

func TestNeighbourCountBottomRight(t *testing.T) {

	area := MakeTwoDimensionalBoolSlice(3, 3)

	area[1][1] = true
	area[1][2] = true

	area[2][1] = true

	assert.Equal(t, 3, neighbourCount(area, 2, 2))
}

// if correctly implemented, the "blinker" should toggle from | to -
func TestBlinker(t *testing.T) {

	area := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	assert.Equal(t, [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}, Progress(area))
}
