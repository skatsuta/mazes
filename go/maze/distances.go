package maze

// Distances is a collection of distances from a root cell.
type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

// NewDistances returns a new Distances that holds distances from the root cell.
func NewDistances(root *Cell) *Distances {
	return &Distances{
		root:  root,
		cells: make(map[*Cell]int),
	}
}

// Get returns a distance from the root to `cell`.
// If there is no distance to the cell, it returns -1.
func (d *Distances) Get(cell *Cell) int {
	if n, found := d.cells[cell]; found {
		return n
	}
	return -1
}

// Set sets `distance` as a distance to `cell`.
func (d *Distances) Set(cell *Cell, distance int) {
	d.cells[cell] = distance
}

// Cells returns all the cells.
func (d *Distances) Cells() []*Cell {
	keys := make([]*Cell, len(d.cells))
	i := 0
	for k := range d.cells {
		keys[i] = k
		i++
	}
	return keys
}
