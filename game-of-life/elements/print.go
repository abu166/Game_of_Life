package elements

import (
	"fmt"
)

func (g *GameOfLife) PrintGrid() {
	for i, row := range g.Grid {
		for j, cell := range row {
			if cell {
				if g.Colored {
					fmt.Print(g.LiveCellColor + g.LiveCellChar + " " + "\033[0m")
				} else {
					fmt.Print(g.LiveCellChar + " ")
				}
			} else if g.Footprints && g.Visited[i][j] {
				if g.Colored {
					fmt.Print(g.VisitedColor + g.VisitedChar + " " + "\033[0m")
				} else {
					fmt.Print(g.VisitedChar + " ")
				}
			} else {
				fmt.Print(g.DeadCellChar + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
