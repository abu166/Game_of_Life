package elements

func (g *GameOfLife) CountLiveCells() int {
	count := 0
	for _, row := range g.Grid {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}
