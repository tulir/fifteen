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

// HasAllNumbers checks if the puzzle contains all integers from 0 to the size of the puzzle.
// 0 is included, as it represents the empty spot and is required for the puzzle to be valid.
func (puzzle *Puzzle) HasAllNumbers() bool {
	// TODO this method might not be needed as Solvable() does this check as well.
	found := make([]bool, puzzle.n*puzzle.n)
	for _, val := range puzzle.data {
		if val < 0 || val >= puzzle.n*puzzle.n || found[val] {
			// Duplicate or out of range value, not valid.
			return false
		}
		found[val] = true
	}
	return true
}

// Solvable checks if the puzzle is solvable.
//
// Time complexity: O(n²)
// Space complexity: O(1)
func (puzzle *Puzzle) Solvable() bool {
	var inversions, blankSpotRow int
	// We embed the code from HasAllNumbers here to reduce time used from 2*n to n
	found := make([]bool, puzzle.n*puzzle.n)
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
	if puzzle.n%2 == 0 {
		return inversions%2 == blankSpotRow%2
	}
	return inversions%2 == 0
}

// IsSolved checks if the puzzle is in its final form.
func (puzzle *Puzzle) IsSolved() bool {
	for i, val := range puzzle.data {
		if i == len(puzzle.data)-1 {
			if val != 0 {
				return false
			}
		} else if val != i+1 {
			return false
		}
	}
	return true
}
