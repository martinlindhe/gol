package main

import (
	"fmt"

	"github.com/gizak/termui"
	"github.com/martinlindhe/gol"
)

var (
	offsetsPar *termui.Par
)

func main() {

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})

	width := termui.TermWidth() - 2
	height := termui.TermHeight() - 2
	population := width * height

	offsetsPar = termui.NewPar("")
	offsetsPar.Width = width + 2
	offsetsPar.Height = height + 2

	area := gol.MakeTwoDimensionalBoolSlice(width, height)

	area = gol.RandomSeed(area, population)

	it := 0

	offsetsPar.BorderLabel = fmt.Sprintf("it %d", it)
	offsetsPar.Text = gol.PrettyPrint(area)
	termui.Render(offsetsPar)

	termui.Handle("/timer/1s", func(e termui.Event) {

		it++
		area = gol.Progress(area)

		offsetsPar.BorderLabel = fmt.Sprintf("it %d", it)
		offsetsPar.Text = gol.PrettyPrint(area)
		termui.Render(offsetsPar)
	})

	termui.Loop()
}
