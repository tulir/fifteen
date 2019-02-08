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
	"maunium.net/go/fifteen/fifteen/util"
	"strconv"
)

// String converts the puzzle into a string.
func (puzzle *Puzzle) String() string {
	maxLen := util.Digits(len(puzzle.data) - 1)
	outLen := len(puzzle.data)*(maxLen+1) - 1
	var data = make([]byte, outLen)
	var dataPtr int

	for i, val := range puzzle.data {
		if i != 0 && i%puzzle.n == 0 {
			// Newline after every n cells
			data[dataPtr] = '\n'
			dataPtr++
		}
		spaces := dataPtr + maxLen - util.Digits(val)
		// Left-pad with spaces
		for ; dataPtr < spaces; dataPtr++ {
			data[dataPtr] = ' '
		}
		if val != 0 {
			// Write number to output
			dataPtr += copy(data[dataPtr:], strconv.Itoa(val))
		} else {
			// Dash at the empty spot instead of space so we'd have some chance
			// of whitespace-independent parsing.
			data[dataPtr-1] = '-'
		}
		// Add space after number except if end of line (newline is added at beginning of loop)
		if (i+1)%puzzle.n != 0 {
			data[dataPtr] = ' '
			dataPtr++
		}
	}
	return string(data)
}
