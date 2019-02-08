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
