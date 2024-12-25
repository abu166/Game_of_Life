package elements

func (g *GameOfLife) CountLiveNeighbors(edgesPortal bool,row, col int) int {
	count := 0
	dirs := []struct{ x, y int }{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		/* cell */ {0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, d := range dirs {
		r, c := row+d.x, col+d.y
		if edgesPortal {
			if r < 0 {
				r = g.Rows - 1
			} else if r >= g.Rows {
				r = 0
			}
			if c < 0 {
				c = g.Cols - 1
			} else if c >= g.Cols {
				c = 0
			}
		} else {
			if r < 0 || r >= g.Rows || c < 0 || c >= g.Cols {
				continue
			}
		}
		if g.Grid[r][c] {
			count++
		}

	}
	return count
}
