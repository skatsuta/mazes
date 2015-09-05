package alg

import (
	"math/rand"
	"time"

	"github.com/skatsuta/mazes/go/maze"
)

// sidewinder is a sidewinder algorithm.
// It implements Algorithm interface.
type sidewinder struct {
	r *rand.Rand
}

// NewSidewinder returns a new sidewinder Algorithm.
func NewSidewinder() Algorithm {
	return sidewinder{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// On generates a maze on a given grid by using a sidewinder algorithm.
func (sw sidewinder) On(grid *maze.Grid) {
	for _, row := range grid.EachRow() {
		run := make([]*maze.Cell, 0, len(row))

		for _, cell := range row {
			run = append(run, cell)

			atEasternBoundary := cell.East == nil
			atNorthanBoundary := cell.North == nil

			shouldCloseOut := atEasternBoundary ||
				(!atNorthanBoundary && sw.r.Intn(2) == 0)

			if shouldCloseOut {
				idx := sw.r.Intn(len(run))
				member := run[idx]
				if member.North != nil {
					member.Link(member.North)
				}
				run = run[:0]
			} else {
				cell.Link(cell.East)
			}
		}
	}
}
