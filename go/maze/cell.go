package maze

const numNeighbors = 4

// Cell is a cell in a maze.
type Cell struct {
	Row, Col                 int
	North, South, East, West *Cell
	links                    map[*Cell]bool
}

// NewCell returns a new Cell put in (row, col).
func NewCell(row, col int) *Cell {
	return &Cell{
		Row: row,
		Col: col,
	}
}

// Link links c and cell bidirectionally.
func (c *Cell) Link(cell *Cell) {
	c.LinkDi(cell, true)
}

// LinkDi links c and cell bidirectionally if bidi is true.
// Otherwise, it only links c to cell.
func (c *Cell) LinkDi(cell *Cell, bidi bool) {
	c.links[cell] = true
	if bidi {
		cell.LinkDi(c, false)
	}
}

// Unlink unlinks c and cell bidirectionally.
func (c *Cell) Unlink(cell *Cell) {
	c.UnlinkDi(cell, true)
}

// UnlinkDi unlinks c and cell bidirectionally if bidi is true.
// Otherwise, it only unlinks c to cell.
func (c *Cell) UnlinkDi(cell *Cell, bidi bool) {
	delete(c.links, cell)
	if bidi {
		cell.UnlinkDi(c, false)
	}
}

// Links returns all the cells linked with c.
func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, len(c.links))
	i := 0
	for k := range c.links {
		keys[i] = k
		i++
	}
	return keys
}

// IsLinked returns true if cell is linked with c.
// Otherwise, it returns false.
func (c *Cell) IsLinked(cell *Cell) bool {
	_, exists := c.links[cell]
	return exists
}

// Neighbors returns all the neighbors of c.
func (c *Cell) Neighbors() []*Cell {
	var nb []*Cell

	if c.North != nil {
		nb = append(nb, c.North)
	}
	if c.South != nil {
		nb = append(nb, c.South)
	}
	if c.East != nil {
		nb = append(nb, c.East)
	}
	if c.West != nil {
		nb = append(nb, c.West)
	}

	return nb
}
