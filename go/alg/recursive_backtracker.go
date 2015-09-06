package alg

import (
	"math/rand"
	"time"

	"github.com/skatsuta/mazes/go/maze"
)

// recursiveBacktracker is the recursive backtracker algorithm.
// It implements Algorithm interface.
type recursiveBacktracker struct {
	r       *rand.Rand
	startAt *maze.Cell
}

// NewRecursiveBacktracker returns a new recursive backtracker Algorithm.
// It uses startAt as a starting point of the algorithm.
func NewRecursiveBacktracker(startAt *maze.Cell) Algorithm {
	return recursiveBacktracker{
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		startAt: startAt,
	}
}

// On generates a maze on a given grid by using the recursive backtracker algorithm.
func (rb recursiveBacktracker) On(grid maze.Grid) {
	stack := []*maze.Cell{rb.startAt}

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		var nbs []*maze.Cell
		for _, nb := range current.Neighbors() {
			if len(nb.Links()) == 0 {
				nbs = append(nbs, nb)
			}
		}

		if len(nbs) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		nb := nbs[rb.r.Intn(len(nbs))]
		current.Link(nb)
		stack = append(stack, nb)
	}
}
