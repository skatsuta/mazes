require 'distance_grid'
require 'binary_tree'

grid = DistanceGrid.new(ARGV[0].to_i, ARGV[1].to_i)
BinaryTree.on(grid)

start = grid[0, 0]
distances = start.distances

grid.distances = distances
puts grid
