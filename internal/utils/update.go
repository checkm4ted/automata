package utils

func (g *Game) Update() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			c := g.Grid[i][j]
			if g.GetNeighbor(*c, 1, 0).Alive {
				c.Alive = !c.Alive
			}
		}
	}
}