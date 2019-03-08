# Usage
Run with `./fifteen-bin <flags>` after building as specified README.
See `./fifteen-bin --help` for a list of flags or see below for in-depth
explanations of each flag.

If using windows, it is recommended to use a unix-like terminal, such as
Cygwin, Windows Subsystem for Linux or Git Bash.

## Puzzle size
Puzzle size can be set with `-w <size>`. Size must be between 3 and 15 (inclusive).

## Random puzzles
Shuffled (`-r shuffle -n <moves>`) or randomized (`-r random`). If `-n <moves>`
is specified without specifying the randomization mode with `-r`, `shuffle` is
assumed. If `-r random` is specified, `-n <moves>` is no-op.

Note that `-r random` puzzles are generally much harder to solve than puzzles
shuffled with a low number (<200) of moves. Shuffles with less than 100 moves
are generally solved nearly instantly

Randomization means generating a completely random array and swapping values
until the puzzle is solvable. Shuffling starts with a solved puzzle and makes
the given number of random moves.

Seed for randomizer can be set with `-s <seed>`. Seed is set to current unix
time in nanoseconds if not specified.

## Animated steps
Animate solution in a ncurses-like UI after solving (`-a solution`) or render
each node entered in the IDA* algorithm (`-a steps`).

Duration of animated solution can be set with `-d <seconds>`.

## I/O
Input file can be set with `-i <path>` and output file can be set with `-o <path>`.
Output defaults to stdout. Either input or randomization mode must be specified.
Output format can be set with `-f <format>` where `<format>` is either `text` or
`json`. Defaults to `text`.

### Input format
The input is either a plaintext number grid or a two-dimensional JSON array.
Zero represents the blank spot. Plaintext input must have columns separated by
spaces and rows separated by newlines.

Since the input is so simple, the input format is automatically detected by
trying to parse the input in both formats.

#### Example
Solved size 4 puzzle.

Plaintext:
```
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 0
```

JSON:
```json
[
	[1, 2, 3, 4],
	[5, 6, 7, 8],
	[9, 10, 11, 12],
	[13, 14, 15, 0]
]
```

### Output format
The output contains the puzzle that was inputted into the solver and the
ordered coordinates of the tiles that a player would click to solve the puzzle.
The time it took to find the solution will also be included except if `-a steps`
is used.

#### Example
Output with flags `-s 1234 -r shuffle -n 10`

Plaintext:
```
Puzzle:
 1  2  3  4
 5     7  8
 9  6 10 12
13 14 11 15

Solution (click coordinates):
2, 3
3, 3
3, 4
4, 4

Duration to solve: 37.931 µs
```

JSON:
Puzzle data as 2-dimensional array under `puzzle`, list of clicks as objects
with `x` and `y` under `clicks` and duration in nanoseconds under `duration`.

```json
{
  "puzzle": [
    [ 1,  2,  3,  4],
    [ 5,  0,  7,  8],
    [ 9,  6, 10, 12],
    [13, 14, 11, 15]
  ],
  "clicks": [
    {"x": 2, "y": 3},
    {"x": 3, "y": 3},
    {"x": 3, "y": 4},
    {"x": 4, "y": 4}
  ],
  "duration": 37931
}
```
