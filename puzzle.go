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

// Puzzle is the base container for 15-puzzles.
type Puzzle struct {
	data []int
	n    int
}

// NewPuzzle creates a new blank puzzle.
func NewPuzzle(n int) *Puzzle {
	return &Puzzle{
		data: make([]int, n*n),
		n:    n,
	}
}

// NewSolvedPuzzle creates a new puzzle in the finished form.
func NewSolvedPuzzle(n int) *Puzzle {
	puzzle := NewPuzzle(n)
	for i := 1; i < len(puzzle.data); i++ {
		puzzle.data[i-1] = i
	}
	return puzzle
}

// Copy creates a copy of this puzzle.
func (puzzle *Puzzle) Copy() *Puzzle {
	newPuzzle := NewPuzzle(puzzle.n)
	copy(puzzle.data, newPuzzle.data)
	return newPuzzle
}

// Get gets the value of a specific slot.
func (puzzle *Puzzle) Get(x, y int) int {
	if x < 0 || y < 0 || x >= puzzle.n || y >= puzzle.n {
		return -1
	}
	return puzzle.data[y*puzzle.n+x]
}

// Set sets the value of a specific slot.
func (puzzle *Puzzle) Set(x, y, val int) {
	if x < 0 || y < 0 || x >= puzzle.n || y >= puzzle.n {
		return
	}
	puzzle.data[y*puzzle.n+x] = val
}

// Move moves the piece at the given coordinates to the empty slot next to it.
// Return false if there is no empty slot next to the given coordinates or if the coordinates are invalid.
// Returns true if the move was successful.
func (puzzle *Puzzle) Move(x, y int) bool {
	val := puzzle.Get(x, y)
	switch {
	case x < 0 || y < 0 || x >= puzzle.n || y >= puzzle.n:
		return false
	case puzzle.Get(x-1, y) == 0:
		puzzle.Set(x-1, y, val)
	case puzzle.Get(x+1, y) == 0:
		puzzle.Set(x+1, y, val)
	case puzzle.Get(x, y-1) == 0:
		puzzle.Set(x, y-1, val)
	case puzzle.Get(x, y+1) == 0:
		puzzle.Set(x, y+1, val)
	default:
		return false
	}
	puzzle.Set(x, y, 0)
	return true
}
