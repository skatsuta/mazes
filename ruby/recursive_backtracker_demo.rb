require 'recursive_backtracker'
require 'grid'

grid = Grid.new(ARGV[0].to_i, ARGV[1].to_i)
RecursiveBacktracker.on(grid)

puts grid
