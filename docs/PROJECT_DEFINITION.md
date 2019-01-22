# Project definition
The aim of the project is to implement an efficient way to solve
[15-puzzles](https://en.wikipedia.org/wiki/15_puzzle).

Based on preliminary research, it seems that [A*](https://en.wikipedia.org/wiki/A*_search_algorithm)
or some similar algorithm can be used along with some heuristics to solve
the problem.

The time complexity of A* depends on the heuristic used, so a relatively fast
heuristic is required. The worst case is O(b^d), where b is the branching
factor (average number of successors per state, directly determined by the
heuristic used) and d is the length of the shortest path.

The space complexity of A* is the same as the time complexity. Due to this, it
might be better to use [iterative deepening A*](https://en.wikipedia.org/wiki/Iterative_deepening_A*)
instead of A*, as IDA* has a worst-case space complexity of O(d).

The puzzles are stored as one-dimensional arrays. A* or IDA* 

The program will be written in [Go](https://golang.org/). It will read and
output text files with the game board columns separated by spaces and rows
separated by newlines.
