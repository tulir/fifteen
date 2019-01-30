# Week 3 report
Algorithm is now ~2x faster than last week thanks to reduced memory waste.

Progress this week (approximate chronological order):
* Added usage instructions.
* Added more comments.
* Stopped using strings.Builder within algorithm.
* Simplified parts of algorithm.
* Switched to FNV-1 hashes for checking if the path contains a specific puzzle already.
* Reduced memory waste in IDA* by making and reversing moves instead of copying the puzzle struct.

Next steps:
* Adding performance tests
* Writing implementation and testing documents

Time spent: 6h
