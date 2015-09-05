package alg

import "github.com/skatsuta/mazes/go/maze"

// Algorithm is an interface that generates a maze on a given Grid by a specific algorithm.
type Algorithm interface {
	On(g *maze.Grid)
}
