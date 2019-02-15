# Implementation document

## File structure
Program is split into 4 packages:
* Main package (repo root) contains most I/O stuff (CLI flags, animating puzzle
  steps with a TUI, file IO)
* `fifteen` package contains `Puzzle` struct which has methods for the
  algorithm and puzzle validation, randomization, parsing and stringifying
* `fifteen/datastructures` contains the `IntStack`, `LinkedMoveStack` and
  `Position` data structures which are utilized by many parts of the `fifteen`
  package.
* `fifteen/util` contains simple utility functions like `Abs()` (absolute value
  of integer) and `Digits()` (number of digits in integer).

Tests for each file are in another file in the same directory with the
`_test.go` suffix as is customary in Go.

## Data structures
The project didn't require many data structures. The main struct is `Puzzle`,
which contains the puzzle size, data and the position of the blank spot. The
puzzle data is stored in a 1-dimensional array with a length of `size^2`. The
array acts as a flattened 2-dimensional array and the cell at x, y is at
index `y*size+x` in the array.

Two stack types are used during the algorithm: `IntStack` is an array-based
data structure with Push(), Remove() and Contains(), while `LinkedMoveStack`
is a linked list with Push(), Pop() and a method to convert it into an array.
Both of these are only used inside the algorithm for keeping track of past
moves in the current move chain.

Finally, `Position` is a simple struct containing `X` and `Y` ints. It is used
to represent the positions that are clicked to make a move.

## Algorithm
The solver uses the [iterative deepening A*](https://en.wikipedia.org/wiki/Iterative_deepening_A*)
algorithm with a [manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry) heuristic.
