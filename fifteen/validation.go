// fifteen - 15-puzzle solver.
// Copyright (C) 2019  Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package fifteen

// Solvable checks if the puzzle is solvable.
//
// Time complexity: O(n²)
// Space complexity: O(1)
func (puzzle *Puzzle) Solvable() bool {
	var inversions, blankSpotRow int
	found := make([]bool, puzzle.n*puzzle.n)
	// Count inversions in puzzle while making sure it's valid.
	for i, val := range puzzle.data {
		if val < 0 || val >= puzzle.n*puzzle.n || found[val] {
			// Duplicate or out of range value, puzzle is not valid and therefore not solvable.
			return false
		}
		found[val] = true

		if val == 0 {
			_, blankSpotRow = puzzle.Coordinates(i)
		}
		for j := i + 1; j < len(puzzle.data); j++ {
			jVal := puzzle.data[j]
			if jVal != 0 && jVal < val {
				inversions++
			}
		}
	}
	// Check if puzzle is solvable based on number of inversions and row number of blank spot.
	if puzzle.n%2 == 0 {
		// For puzzles with even sizes, the puzzle is solvable if the number of
		// inversions and the row number of the blank spot are both even or both odd.
		return inversions%2 == blankSpotRow%2
	}
	// For puzzles with odd sizes, the puzzle is solvable if the number of inversions is even.
	return inversions%2 == 0
}

// IsSolved checks if the puzzle is in its final form.
func (puzzle *Puzzle) IsSolved() bool {
	for i, val := range puzzle.data {
		if val != i+1 {
			return i == len(puzzle.data)-1 && val == 0
		}
	}
	return false
}
