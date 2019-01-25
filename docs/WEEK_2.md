# Week 2 report
Started with adding missing details to project definition as requested in
feedback.

The program can now read input puzzles, generate or shuffle puzzles, solve easy
puzzles fairly quickly, output solutions and render solutions with a fancy TUI
view.

More detailed list of things that were done (approximate chronological order):
* Basic puzzle struct and methods
* String puzzle parser
* Puzzle validation methods
* Manhattan distance calculation
* Puzzle shuffling
* Most unit tests except for the algorithm itself
* Set up Travis CI for running tests, Code Climate for static analysis and Godoc
  for viewing the inline docs.
* Initial IDA* solver
* Main program with I/O and a fancy TUI view of solutions
* Random puzzle generator

Next steps:
* Improve performance of algorithm
* Add unit tests for algorithm

Time spent: 11h
