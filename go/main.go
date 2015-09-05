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
	grid := maze.NewGrid(row, col)

	var algorithm alg.Algorithm
	switch flag.Arg(0) {
	case "binarytree":
		algorithm = alg.NewBinaryTree()
	case "sidewinder":
		algorithm = alg.NewSidewinder()
	}

	algorithm.On(grid)
	fmt.Println(grid.String())
}
