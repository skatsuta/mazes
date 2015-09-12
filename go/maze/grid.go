package maze

import (
	"bytes"
	"math/rand"
	"time"
)

// Grid is an interface of grid.
type Grid interface {
	Rows() int
	Cols() int
	Get(int, int) *Cell
	Random() *Cell
	Size() int
	EachRow() [][]*Cell
	EachCell() []*Cell
	String() string
	DeadEnds() []*Cell
	Braid(float64)
}

// NormalGrid is a grid containing all the cells.
type NormalGrid struct {
	rows, cols int
	grid       [][]*Cell
	r          *rand.Rand
}

// NewNormalGrid returns a new NormalGrid whose size is rows by cols.
func NewNormalGrid(rows, cols int) Grid {
	g := &NormalGrid{
		rows: rows,
		cols: cols,
		grid: prepareGrid(rows, cols),
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	g.configureCells()
	return g
}

// Rows returns the number of rows of `g`.
func (g *NormalGrid) Rows() int {
	return g.rows
}

// Cols returns the number of columns of `g`.
func (g *NormalGrid) Cols() int {
	return g.cols
}

// Get returns a cell on (row, col).
func (g *NormalGrid) Get(row, col int) *Cell {
	if row < 0 || row > g.rows-1 {
		return nil
	}
	if col < 0 || col > g.cols-1 {
		return nil
	}
	return g.grid[row][col]
}

// Random returns a cell chosen randomly from a grid.
func (g *NormalGrid) Random() *Cell {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	row := r.Intn(g.rows)
	col := r.Intn(g.cols)
	return g.Get(row, col)
}

// Size returns a size of g.
func (g *NormalGrid) Size() int {
	return g.rows * g.cols
}

// EachRow returns each row.
func (g *NormalGrid) EachRow() [][]*Cell {
	return g.grid
}

// EachCell returns each cell.
func (g *NormalGrid) EachCell() []*Cell {
	cells := make([]*Cell, 0, g.Size())
	for _, row := range g.grid {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}
	return cells
}

// String draws a maze by an ASCII art.
func (g *NormalGrid) String() string {
	return g.stringWithContentsFunc(g.contentsOf)
}

// stringContentsFunc draws a maze by an ASCII art using `f`.
func (g *NormalGrid) stringWithContentsFunc(f func(*Cell) string) string {
	var (
		space  = "   "
		wall   = "|"
		corner = "+"
		line   = "---"
	)

	var output bytes.Buffer

	_, _ = output.WriteString(corner)
	for i := 0; i < g.cols; i++ {
		_, _ = output.WriteString(line + corner)
	}
	_, _ = output.WriteString("\n")

	mid := bytes.NewBuffer([]byte(wall))
	btm := bytes.NewBuffer([]byte(corner))
	for _, row := range g.EachRow() {
		// initialize all but the first character
		mid.Truncate(1)
		btm.Truncate(1)

		for _, cell := range row {
			if cell == nil {
				// dummy cell
				cell = NewCell(-1, -1)
			}

			body := " " + f(cell) + " "
			_, _ = mid.WriteString(body)
			if cell.IsLinked(cell.East) {
				_, _ = mid.WriteString(" ")
			} else {
				_, _ = mid.WriteString(wall)
			}

			if cell.IsLinked(cell.South) {
				_, _ = btm.WriteString(space)
			} else {
				_, _ = btm.WriteString(line)
			}
			_, _ = btm.WriteString(corner)
		}

		_, _ = output.Write(mid.Bytes())
		_, _ = output.WriteString("\n")
		_, _ = output.Write(btm.Bytes())
		_, _ = output.WriteString("\n")
	}

	return output.String()
}

func (g *NormalGrid) contentsOf(cell *Cell) string {
	return " "
}

func (g *NormalGrid) configureCells() {
	for _, cell := range g.EachCell() {
		row, col := cell.Row, cell.Col
		cell.North = g.Get(row-1, col)
		cell.South = g.Get(row+1, col)
		cell.East = g.Get(row, col+1)
		cell.West = g.Get(row, col-1)
	}
}

// DeadEnds returns all the dead-end cells.
func (g *NormalGrid) DeadEnds() []*Cell {
	var list []*Cell

	for _, cell := range g.EachCell() {
		if len(cell.Links()) == 1 {
			list = append(list, cell)
		}
	}

	return list
}

// Braid rearranges g to "braid" one, that is, a maze without any dead ends.
func (g *NormalGrid) Braid(p float64) {
	for _, cell := range Shuffle(g.DeadEnds()) {
		if len(cell.Links()) != 1 || g.r.Float64() > p {
			continue
		}

		var nbs, best []*Cell
		for _, nb := range cell.Neighbors() {
			if !nb.IsLinked(cell) {
				nbs = append(nbs, nb)
			}
		}

		for _, n := range nbs {
			if len(n.Links()) == 1 {
				best = append(best, n)
			}
		}

		if len(best) == 0 {
			best = nbs
		}

		idx := g.r.Intn(len(best))
		cell.Link(best[idx])
	}
}

// prepareGrid returns a rows-by-cols 2D Cell array.
func prepareGrid(rows, cols int) [][]*Cell {
	grid := make([][]*Cell, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]*Cell, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = NewCell(i, j)
		}
	}
	return grid
}
