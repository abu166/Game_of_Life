package elements

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (g *GameOfLife) GenerateRandomGrid(dimensions string) error {
	dims := strings.Split(dimensions, "x")
	if len(dims) != 2 {
		return fmt.Errorf("invalid format for --random. Expected format WxH")
	}
	var err1, err2 error
	g.Rows, err1 = strconv.Atoi(dims[0])
	g.Cols, err2 = strconv.Atoi(dims[1])
	rows := g.Rows
	cols := g.Cols
	if err1 != nil || err2 != nil || rows < 3 || cols < 3 || rows > 100 || cols > 100 {
		return fmt.Errorf("invalid dimensions for --random. Both width and height must be integers in range from 3 to 100")
	}

	rand.Seed(time.Now().UnixNano())
	g.Grid = make([][]bool, rows)
	grid := g.Grid
	for i := range grid {
		grid[i] = make([]bool, cols)
		for j := range grid[i] {
			// Randomly assign live or dead cells
			grid[i][j] = rand.Intn(2) == 1
		}
	}

	return nil
}
