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

type position struct {
	x int
	y int
}

func (pos position) valid(size int) bool {
	return pos.x >= 1 && pos.y >= 1 && pos.x <= size && pos.y <= size
}

// Shuffle makes the given amount of random moves on the puzzle.
func (puzzle *Puzzle) Shuffle(moves int) {
	var pos position
	for i := len(puzzle.data) - 1; i >= 0; i-- {
		if puzzle.data[i] == 0 {
			pos.x, pos.y = puzzle.Coordinates(i)
			break
		}
	}
	var validMoves []position
	var allMoves [4]position
	for i := 0; i < moves; i++ {
		allMoves = [4]position{
			{pos.x - 1, pos.y},
			{pos.x + 1, pos.y},
			{pos.x, pos.y - 1},
			{pos.x, pos.y + 1},
		}
		validMoves = make([]position, 0, 4)
		for _, move := range allMoves {
			if move.valid(puzzle.n) {
				validMoves = append(validMoves, move)
			}
		}
		pos = validMoves[rand.Intn(len(validMoves))]
		puzzle.Move(pos.x, pos.y)
	}
}
