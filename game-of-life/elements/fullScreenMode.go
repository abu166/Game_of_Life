package elements

import (
	"fmt"

	"github.com/inancgumus/screen"
)

func (g *GameOfLife) Fullscreen() {
	width, height := screen.Size()
	height = height / 2
	fmt.Printf("Width: %d, Height: %d\n", width, height)
	g.Rows = height
	g.Cols = width

	fullscreenGreed := make([][]bool, height)
	for i := range fullscreenGreed {
		fullscreenGreed[i] = make([]bool, width) // Create each row with the specified width
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if j < len(g.Grid[0]) && i < len(g.Grid) {
				fullscreenGreed[i][j] = g.Grid[i][j]
			} else {
				fullscreenGreed[i][j] = false // Use a boolean value instead of a string
			}
		}
	}
	g.Grid = fullscreenGreed
}
