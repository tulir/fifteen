# 15-puzzle solver
[![Build Status](https://travis-ci.org/tulir/fifteen.svg?branch=master)](https://travis-ci.org/tulir/fifteen)
[![Maintainability](https://api.codeclimate.com/v1/badges/d8575cfd2ecbeaebc4c3/maintainability)](https://codeclimate.com/github/tulir/fifteen/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d8575cfd2ecbeaebc4c3/test_coverage)](https://codeclimate.com/github/tulir/fifteen/test_coverage)
[![GoDoc Reference](https://godoc.org/maunium.net/go/fifteen/fifteen?status.svg)](http://godoc.org/maunium.net/go/fifteen/fifteen)

Data structures and algorithms project, 15-puzzle solver. Written in Go.

* [Project definition](docs/PROJECT_DEFINITION.md)
* Week reports: [1](docs/WEEK_1.md), [2](docs/WEEK_2.md)

## Instructions
### Installation
0. Install [Go](https://golang.org/) 1.11 or higher
1. Clone the repository with `git clone https://github.com/tulir/fifteen.git`
2. Fetch dependencies and compile with `go build`

To update, simply pull changes and `go build` again.

### Usage
Run with `./fifteen <flags>` after building as specified in previous section.
See `./fifteen --help` for a list of flags.

#### Puzzle size
Puzzle size can be set with `-w <size>`. Size must be between 3 and 15 (inclusive).

#### Random puzzles
Shuffled (`-r shuffle -n <moves>`) or randomized (`-r random`). If `-n <moves>`
is specified without specifying the randomization mode with `-r`, `shuffle` is
assumed. If `-r random` is specified, `-n <moves>` is no-op.

Note that `-r random` puzzles are generally much harder to solve than puzzles
shuffled with a low number (<200) of moves.

Randomization means generating a completely random array and swapping values
until the puzzle is solvable. Shuffling starts with a solved puzzle and makes
the given number of random moves.

Seed for randomizer can be set with `-s <seed>`. Seed is set to current unix
time in nanoseconds if not specified.

#### Animated steps
Animate solution in a ncurses-like UI after solving (`-a solution`) or render
each node entered in the IDA* algorithm (`-a steps`).

Duration of animated solution can be set with `-d <seconds>`.

#### I/O
Input file can be set with `-i <path>` and output file can be set with `-o <path>`.
Output defaults to stdout. Either input or randomization mode must be specified.
Output format can be set with `-f <format>` where `<format>` is either `text` or
`json`. Defaults to `text`.

### Tests
Run tests and output coverage results:
```bash
go test ./... -coverprofile=coverage.out
```

Read coverage info in terminal:
```bash
go tool cover -func=coverage.out
```

Generate coverage report and open in browser:
```bash
go tool cover -html=coverage.out
```

Generate coverage report and save to file:
```bash
go tool cover -html=coverage.out -o file.html
```
