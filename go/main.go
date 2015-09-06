package main

import (
	"flag"
	"fmt"

	"github.com/skatsuta/mazes/go/alg"
	"github.com/skatsuta/mazes/go/maze"
)

var row, col int

func init() {
	flag.IntVar(&row, "row", 4, "the number of rows of a maze")
	flag.IntVar(&col, "col", 4, "the number of columns of a maze")
	flag.Parse()
}

func main() {
	mode := flag.Arg(0)
	switch mode {
	case "binarytree", "sidewinder":
		grid(mode)
	case "dijkstra":
		distanceGrid(mode)
	}
}

func grid(mode string) {
	grid := maze.NewNormalGrid(row, col)

	var algorithm alg.Algorithm
	switch mode {
	case "binarytree":
		algorithm = alg.NewBinaryTree()
	case "sidewinder":
		algorithm = alg.NewSidewinder()
	}

	algorithm.On(grid)
	fmt.Println(grid.String())
}

func distanceGrid(mode string) {
	grid := maze.NewDistanceGrid(row, col)

	algorithm := alg.NewBinaryTree()

	algorithm.On(grid)
	start := grid.Get(0, 0)
	grid.(*maze.DistanceGrid).Distances = start.Distances()

	fmt.Println(grid.String())
}
