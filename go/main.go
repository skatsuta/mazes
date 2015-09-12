package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skatsuta/mazes/go/alg"
	"github.com/skatsuta/mazes/go/maze"
)

var (
	row, col int
	dijkstra bool
	usage    = `Usage: %s [-row n] [-col m] [-dijkstra] algorithm

ALGORITHMS:
  binarytree
        binary tree algorithm
  sidewinder
        sidewinder algorithm
  recursivebacktracker
        recursive backtracker algorithm

OPTIONS:
`
)

func init() {
	flag.IntVar(&row, "row", 4, "the number of rows of a maze")
	flag.IntVar(&col, "col", 4, "the number of columns of a maze")
	flag.BoolVar(&dijkstra, "dijkstra", false, "show distances using Dijkstra algorthm")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	var (
		grid maze.Grid
		al   alg.Algorithm
	)

	// choose grid type
	if dijkstra {
		grid = maze.NewDistanceGrid(row, col)
	} else {
		grid = maze.NewNormalGrid(row, col)
	}

	// choose algorithm to use
	switch flag.Arg(0) {
	case "binarytree":
		al = alg.NewBinaryTree()
	case "sidewinder":
		al = alg.NewSidewinder()
	case "recursivebacktracker":
		al = alg.NewRecursiveBacktracker(grid.Random())
	default:
		panic("unknown algorithm")
	}

	// generate a maze by al
	al.On(grid)

	if dijkstra {
		start := grid.Get(0, 0)
		grid.(*maze.DistanceGrid).Distances = start.Distances()
	}

	fmt.Println(grid.String())
	fmt.Printf("%d dead-ends\n", len(grid.DeadEnds()))
}
