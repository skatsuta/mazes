package alg

import (
	"math/rand"
	"time"

	"github.com/skatsuta/mazes/go/maze"
)

// binaryTree is a binary tree algorithm.
// It implements Algorithm interface.
type binaryTree struct {
	r *rand.Rand
}

// NewBinaryTree returns a new binary tree Algorithm.
func NewBinaryTree() Algorithm {
	return binaryTree{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// On generates a maze on a given grid by using a binary tree algorithm.
func (bt binaryTree) On(grid maze.Grid) {
	for _, cell := range grid.EachCell() {
		var neighbors []*maze.Cell
		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}
		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}

		if len(neighbors) < 1 {
			continue
		}

		idx := bt.r.Intn(len(neighbors))
		neighbor := neighbors[idx]
		cell.Link(neighbor)
	}
}
