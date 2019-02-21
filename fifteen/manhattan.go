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
	"maunium.net/go/fifteen/fifteen/datastructures"
	"maunium.net/go/fifteen/fifteen/util"
)

// ManhattanDistance returns the sum of the manhattan distances between the cells and their final positions.
// This is used as the heuristic for the (ID)A* algorithm in the solver.
//
// Time complexity:  O(n)
// Space complexity: O(1)
func (puzzle *Puzzle) ManhattanDistance() (sum int) {
	pX, pY := 1, 1
	var tX, tY int
	for _, val := range puzzle.data {
		if val == 0 {
			val = puzzle.n * puzzle.n
		}
		tX, tY = puzzle.Coordinates(val - 1)
		// https://en.wikipedia.org/wiki/Taxicab_geometry
		// Manhattan distance between (pX, pY) and (tX, tY) is |pX - tX| + |pY - tY|
		sum += util.Abs(pX-tX) + util.Abs(pY-tY)
		// We're iterating over the puzzle data in order (left to right, top to bottom).
		// Increment pX and pY appropriately.
		pX++
		if pX > puzzle.n {
			pX = 1
			pY++
		}
	}
	return
}

func (puzzle *Puzzle) ManhattanDiff(value int, from, to ds.Position) (sum int) {
	if value == 0 {
		value = puzzle.n * puzzle.n
	}
	tX, tY := puzzle.Coordinates(value-1)
	origDist := util.Abs(from.X - tX) + util.Abs(from.Y - tY)
	newDist := util.Abs(to.X - tX) + util.Abs(to.Y - tY)
	return newDist - origDist
}
