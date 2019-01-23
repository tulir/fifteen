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

func (pos Position) AllMoves() (Position, Position, Position, Position) {
	return Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X, pos.Y - 1},
		Position{pos.X, pos.Y + 1}
}

func (pos Position) AllMovesArray() [4]Position {
	a, b, c, d := pos.AllMoves()
	return [4]Position{a, b, c, d}
}

func (pos Position) ValidMoves(n int) (validMoves []Position) {
	for _, move := range pos.AllMovesArray() {
		if move.Valid(n) {
			validMoves = append(validMoves, move)
		}
	}
	return
}

func (pos Position) String() string {
	return fmt.Sprintf("%d, %d", pos.X, pos.Y)
}
