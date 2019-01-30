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

// ManhattanDistance returns the sum of the manhattan distances between the cells and their final positions.
// This is used as the heuristic for the (ID)A* algorithm in the solver.
//
// Time complexity:  O(n)
// Space complexity: O(1)
func (puzzle *Puzzle) ManhattanDistance() int {
	var sum int
	for i, val := range puzzle.data {
		sum += puzzle.manhattanDistance(i, val)
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (puzzle *Puzzle) manhattanDistance(position, value int) int {
	if value == 0 {
		// Calculate empty slot as last value
		value = puzzle.n * puzzle.n
	}
	pX, pY := puzzle.Coordinates(position)
	tX, tY := puzzle.Coordinates(value-1)
	return abs(pX - tX) + abs(pY - tY)
}
