package elements

func (g *GameOfLife) NextTick(edgePortal bool) {
	newGrid := make([][]bool, g.Rows)
	for i := range g.Grid {
		newGrid[i] = make([]bool, g.Cols)
		for j := range g.Grid[i] {
			liveNeighbors := g.CountLiveNeighbors(edgePortal, i, j)
			if g.Grid[i][j] && (liveNeighbors == 2 || liveNeighbors == 3) {
				newGrid[i][j] = true
			} else if !g.Grid[i][j] && liveNeighbors == 3 {
				newGrid[i][j] = true
			} else {
				newGrid[i][j] = false
			}

			if g.Footprints && (g.Grid[i][j] || newGrid[i][j]) {
				g.Visited[i][j] = true
			}
		}
	}
	g.Grid = newGrid
	g.Tick++
}
