require 'grid'
require 'binary_tree'

grid = Grid.new(ARGV[0].to_i, ARGV[1].to_i)
BinaryTree.on(grid)

puts grid
puts "#{grid.deadends.count} dead-ends"
