package main

import (
	"flag"
	"fmt"

	"github.com/skatsuta/mazes/go/alg/binarytree"
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

	switch flag.Arg(0) {
	case "binarytree":
		binarytree.On(grid)
	}

	fmt.Println(grid.String())
}
