package maze

import "strconv"

// DistanceGrid is a grid that also holds distances.
type DistanceGrid struct {
	Grid
	*Distances
}

// NewDistanceGrid returns a new DistanceGrid whose size is `rows`-by-`cols`.
func NewDistanceGrid(rows, cols int) Grid {
	return &DistanceGrid{
		Grid: NewNormalGrid(rows, cols),
	}
}

// Get returns a cell in (row, col).
func (g *DistanceGrid) Get(row, col int) *Cell {
	return g.Grid.Get(row, col)
}

// String draws a maze by an ASCII art.
func (g *DistanceGrid) String() string {
	return g.Grid.(*NormalGrid).stringWithContentsFunc(g.contentsOf)
}

// contentsOf returns contents of `cell`.
// This overrides Grid.contentsOf().
func (g *DistanceGrid) contentsOf(cell *Cell) string {
	if g.Distances == nil {
		return " "
	}

	dist := g.Distances.Get(cell)
	if dist < 0 {
		return " "
	}
	base := 36
	return strconv.FormatInt(int64(dist), base)
}
