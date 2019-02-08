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

import (
	"math/rand"
	"maunium.net/go/fifteen/fifteen/datastructures"
)

// NewRandomPuzzle creates a puzzle with completely random values.
func NewRandomPuzzle(n int) (*Puzzle, error) {
	puzzle, err := NewPuzzle(n)
	if err != nil {
		return nil, err
	}
	puzzle.data = rand.Perm(len(puzzle.data))
	ptr := len(puzzle.data) - 1
	for !puzzle.Solvable() {
		puzzle.data[ptr], puzzle.data[ptr-1] = puzzle.data[ptr-1], puzzle.data[ptr]
		ptr--
	}
	for i, val := range puzzle.data {
		if val == 0 {
			puzzle.blank.X, puzzle.blank.Y = puzzle.Coordinates(i)
			break
		}
	}
	return puzzle, nil
}

// Shuffle makes the given amount of random moves on the puzzle.
func (puzzle *Puzzle) Shuffle(moves int) {
	pos := puzzle.blank
	vmPtr := 0
	var validMoves [4]ds.Position
	var allMoves [4]ds.Position
	for i := 0; i < moves; i++ {
		vmPtr = 0
		allMoves[0], allMoves[1], allMoves[2], allMoves[3] = pos.AllMoves()

		for _, move := range allMoves {
			if move.Valid(puzzle.n) {
				validMoves[vmPtr] = move
				vmPtr++
			}
		}
		pos = validMoves[rand.Intn(vmPtr)]
		puzzle.MovePos(pos)
	}
}
