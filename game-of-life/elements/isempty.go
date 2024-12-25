package elements

func (g *GameOfLife) IsEmpty() bool {
	for _, row := range g.Grid {
		for _, cell := range row {
			if cell {
				return false
			}
		}
	}
	return true
}
