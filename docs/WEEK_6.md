# Week 6 report
Implemented improvement ideas from peer review
* Fixed build instructions
* Included suggested tests
* Made parser check validity and set the blank position
* Fixed blank position in NewSolvedPuzzle
* Moved getting valid moves to `Puzzle`
* Added move validation to `Move()`

Second peer review, of @alafuzof's tira-puristin [here](https://github.com/alafuzof/tira-puristin)

I looked into Kenichiro Takahashi's heuristics more, but it turns out they're
specific to 4x4 puzzles and are not documented properly. Increasing the IDA*
bound by 2 instead of 1 every iteration was the only improvement that was easy
enough to understand, so I added it as a special case for size 4 puzzles.

The puzzle solvability verification in the random puzzle generator had a bug
where it would fail to make the puzzle solvable if the blank position was at
the end of the generated puzzle. This issue has been fixed and a new test was
added for it.

Next steps:
* Add performance tests for data structures
