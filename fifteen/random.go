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
)

// NewRandomPuzzle creates a puzzle with completely random values.
// It is not guaranteed to be solvable.
func NewRandomPuzzle() {
	// TODO
}

// Shuffle makes the given amount of random moves on the puzzle.
func (puzzle *Puzzle) Shuffle(moves int) {
	pos := puzzle.blank
	var validMoves []Position
	var allMoves [4]Position
	for i := 0; i < moves; i++ {
		allMoves[0], allMoves[1], allMoves[2], allMoves[3] = pos.AllMoves()

		validMoves = make([]Position, 0, 4)
		for _, move := range allMoves {
			if move.Valid(puzzle.n) {
				validMoves = append(validMoves, move)
			}
		}
		pos = validMoves[rand.Intn(len(validMoves))]
		puzzle.Move(pos.X, pos.Y)
	}
}
