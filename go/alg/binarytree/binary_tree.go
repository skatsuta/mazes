package binarytree

import (
	"math/rand"
	"time"

	"github.com/skatsuta/mazes/go/maze"
)

// On generates a maze on a given grid by using a binary tree algorithm.
func On(grid *maze.Grid) {
	for _, cell := range grid.EachCell() {
		var neighbors []*maze.Cell
		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}
		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}

		if len(neighbors) < 1 {
			return
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(len(neighbors))
		neighbor := neighbors[idx]

		if neighbor != nil {
			cell.Link(neighbor)
		}
	}
}
