# Project definition
The aim of the project is to implement an efficient algorithm that solves
[15-puzzles](https://en.wikipedia.org/wiki/15_puzzle). While the primary goal
is solving 15-puzzles specifically, other sizes will also be supported by the
program. This document refers to the width and height of the game board as
"puzzle size". The puzzle size of 15-puzzles is 4.

Based on preliminary research, it seems that [A*](https://en.wikipedia.org/wiki/A*_search_algorithm)
or some similar algorithm can be used along with some heuristics to solve
the problem.

The time complexity of A* depends on the heuristic used, so a relatively fast
heuristic is required. The worst case is `O(b^d)`. `b` is the branching factor,
i.e. the average number of successors per state, and d is the length of the
shortest path. In this case, there are always 2, 3 or 4 possible successors
with 3 being the average when puzzle size = 4. For larger puzzle sizes, the
average moves closer to 4. `b = (4 * (size-2)^2 + 3 * (size-2)*4 + 2 * 4) / size^2`
15-puzzles have been proven to be solvable in 80 moves. Therefore, when puzzle
size = 4, b = 3 and 0 ≤ d ≤ 80.

The heuristic used will be the sum of the manhattan distance from each tile to
the target position of that tile. As it involves the distance of each tile,
the time complexity is O(n) where n is the number of tiles (i.e. puzzle size^2).

The space complexity of A* is the same as the time complexity. Due to this, it
might be better to use [iterative deepening A*](https://en.wikipedia.org/wiki/Iterative_deepening_A*)
instead of A*, as IDA* has a worst-case space complexity of O(d).

The puzzles are stored as one-dimensional arrays. IDA* requires a stack data
structure for remembering the current path.

The program will be written in [Go](https://golang.org/). It will read text
files with the game board columns separated by spaces and rows separated by
newlines or JSON with the board as a two-dimensional array. It will output
the solution as a series of coordinates either separated by commas and newlines
or as JSON objects with an `x` and `y` field.

The program will take options as command-line flags. It will also optionally
show step-by-step moves with a ncurses-like UI.

Sources (also linked inline in text):
* [en.wikipedia.org/wiki/15_puzzle](https://en.wikipedia.org/wiki/15_puzzle)
* [en.wikipedia.org/wiki/A*_search_algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm)
* [en.wikipedia.org/wiki/Iterative_deepening_A*](https://en.wikipedia.org/wiki/Iterative_deepening_A*)
