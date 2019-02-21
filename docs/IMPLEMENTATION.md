# Implementation document
The project is a solver for [15-puzzles](https://en.wikipedia.org/wiki/15_puzzle).
It can theoretically solve any such n*n puzzle between 3x3 and 15x15, but it is
mostly intended for 3x3, 4x4 and 5x5 puzzles. The application can either generate
a random puzzle, make a given number of moves to shuffle a solved puzzle or read
a puzzle from file or user input. After solving, the application can output the
steps required to solve as a series of positions to click or as a text animation.

## Algorithm
The solver uses the [iterative deepening A*](https://en.wikipedia.org/wiki/Iterative_deepening_A*)
algorithm with a [manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry) heuristic.

For 4x4 puzzles, it uses a special case optimization of increasing the IDA*
bound by two instead of one every iteration. This optimization is from
[Kenichiro Takahashi's document about 15-puzzle solving heuristics](https://web.archive.org/web/20141224035932/http://juropollo.xe0.ru/stp_wd_translation_en.htm).

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

## Problems and potential improvements
Kenichiro Takahashi had developed other heuristics for speeding up 15-puzzle
solving, but they were not implemented due to poor documentation and partly due
to being specific to 4x4 puzzles. An improved version could use his walking
distance heuristic when solving 4x4 puzzles, or maybe even generalize it to
work on other sizes.

Memory usage could have probably been reduced by using smaller types (e.g. uint8
instead of int), but these gains might not have been that great.
