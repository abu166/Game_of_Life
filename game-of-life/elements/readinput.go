package elements

import (
	"bufio"
	"fmt"
	"strings"
)

type GameOfLife struct {
	Grid          [][]bool
	Rows, Cols    int
	Tick          int
	Colored       bool
	Footprints    bool
	LiveCellChar  string
	DeadCellChar  string
	VisitedChar   string
	LiveCellColor string
	VisitedColor  string
	Visited       [][]bool
}

func (g *GameOfLife) ReadInput(scanner *bufio.Scanner) error {
	// Read grid dimensions
	if !scanner.Scan() {
		return fmt.Errorf("invalid input")
	}
	dimensions := strings.Fields(scanner.Text())
	if len(dimensions) != 2 {
		return fmt.Errorf("invalid input")
	}
	fmt.Sscanf(dimensions[0], "%d", &g.Rows)
	fmt.Sscanf(dimensions[1], "%d", &g.Cols)
	if g.Rows < 3 || g.Cols < 3 || g.Rows > 100 || g.Cols > 100 {
		return fmt.Errorf("Both width and height must be integers in range from 3 to 100")
	}

	// Read grid lines
	g.Grid = make([][]bool, g.Rows)

	for i := 0; i < g.Rows; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("invalid input")
		}
		line := scanner.Text()
		if len(line) != g.Cols {
			return fmt.Errorf("invalid input")
		}
		g.Grid[i] = make([]bool, g.Cols)
		for j, char := range line {
			if char == '#' {
				g.Grid[i][j] = true
			} else if char == '.' {
				g.Grid[i][j] = false
			} else {
				return fmt.Errorf("invalid input")
			}
		}
	}

	return nil
}
