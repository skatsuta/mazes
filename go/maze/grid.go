package maze

import (
	"bytes"
	"math/rand"
	"time"
)

// Grid is a grid containing all the cells.
type Grid struct {
	Rows, Cols int
	grid       [][]*Cell
}

// NewGrid returns a new Grid whose size is rows by cols.
func NewGrid(rows, cols int) *Grid {
	g := &Grid{
		Rows: rows,
		Cols: cols,
		grid: prepareGrid(rows, cols),
	}
	g.configureCells()
	return g
}

// Get returns a cell on (row, col).
func (g *Grid) Get(row, col int) *Cell {
	if row < 0 || row > g.Rows-1 {
		return nil
	}
	if col < 0 || col > g.Cols-1 {
		return nil
	}
	return g.grid[row][col]
}

// Random returns a cell chosen randomly from a grid.
func (g *Grid) Random() *Cell {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	row := r.Intn(g.Rows)
	col := r.Intn(g.Cols)
	return g.Get(row, col)
}

// Size returns a size of g.
func (g *Grid) Size() int {
	return g.Rows * g.Cols
}

// EachRow returns each row.
func (g *Grid) EachRow() [][]*Cell {
	return g.grid
}

// EachCell returns each cell.
func (g *Grid) EachCell() []*Cell {
	cells := make([]*Cell, 0, g.Size())
	for _, row := range g.grid {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}
	return cells
}

// String draws a maze by an ASCII art.
func (g *Grid) String() string {
	var (
		body   = "   "
		space  = " "
		wall   = "|"
		corner = "+"
		line   = "---"
	)

	var output bytes.Buffer

	_, _ = output.WriteString(corner)
	for i := 0; i < g.Cols; i++ {
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

			_, _ = mid.WriteString(body)
			if cell.IsLinked(cell.East) {
				_, _ = mid.WriteString(space)
			} else {
				_, _ = mid.WriteString(wall)
			}

			if cell.IsLinked(cell.South) {
				_, _ = btm.WriteString(body)
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

func (g *Grid) configureCells() {
	for _, cell := range g.EachCell() {
		row, col := cell.Row, cell.Col
		cell.North = g.Get(row-1, col)
		cell.South = g.Get(row+1, col)
		cell.East = g.Get(row, col+1)
		cell.West = g.Get(row, col-1)
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
