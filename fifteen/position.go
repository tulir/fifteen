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
	"fmt"
)

// Position is a simple X, Y tuple
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Valid checks if the given position is valid in a %n%-puzzle of the given size.
func (pos Position) Valid(size int) bool {
	return pos.X >= 1 && pos.Y >= 1 && pos.X <= size && pos.Y <= size
}

// AllMoves returns all four moves from this position. No checks are done to ensure the validity of the moves.
// Use the ValidMoves method to get valid moves from this position.
func (pos Position) AllMoves() (Position, Position, Position, Position) {
	return Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X, pos.Y - 1},
		Position{pos.X, pos.Y + 1}
}

// AllMovesArray returns the output of AllMoves() as a fixed-length array instead of four return values.
func (pos Position) AllMovesArray() [4]Position {
	a, b, c, d := pos.AllMoves()
	return [4]Position{a, b, c, d}
}

// ValidMoves returns the valid moves from this position on a puzzle of the given size.
func (pos Position) ValidMoves(n int) (validMoves []Position) {
	moveCount := 4
	if pos.X == 1 || pos.X == n {
		moveCount--
	}
	if pos.Y == 1 || pos.Y == n {
		moveCount--
	}
	validMoves = make([]Position, moveCount)
	for _, move := range pos.AllMovesArray() {
		if move.Valid(n) {
			moveCount--
			validMoves[moveCount] = move
		}
	}
	return
}

// String returns the coordinates of this position in a string.
func (pos Position) String() string {
	return fmt.Sprintf("%d, %d", pos.X, pos.Y)
}
