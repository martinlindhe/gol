package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/martinlindhe/gol"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	world      *gol.World
	noiseImage *image.RGBA
)

func update(screen *ebiten.Image) error {

	world.Progress()
	world.DrawImage(noiseImage)
	screen.ReplacePixels(noiseImage.Pix)

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	return nil
}

func main() {

	population := int((screenWidth * screenHeight) / 10)
	scale := 2.0

	world = gol.NewWorld(screenWidth, screenHeight)
	world.RandomSeed(population)

	noiseImage = image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	if err := ebiten.Run(update, screenWidth, screenHeight, scale, "game of life"); err != nil {
		log.Fatal(err)
	}
}
