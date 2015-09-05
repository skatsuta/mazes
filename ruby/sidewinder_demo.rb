require 'grid'
require 'sidewinder'

grid = Grid.new(ARGV[0].to_i, ARGV[1].to_i)
Sidewinder.on(grid)

puts grid
