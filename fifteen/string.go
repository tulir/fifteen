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
	"strconv"
)

func Digits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func (puzzle *Puzzle) String() string {
	maxLen := Digits(len(puzzle.data) - 1)
	outLen := len(puzzle.data)*(maxLen+1) - 1
	var data = make([]byte, outLen)
	var dataPtr int

	for i, val := range puzzle.data {
		if i != 0 && i%puzzle.n == 0 {
			data[dataPtr] = '\n'
			dataPtr++
		}
		spaces := dataPtr + maxLen - Digits(val)
		for ; dataPtr < spaces; dataPtr++ {
			data[dataPtr] = ' '
		}
		if val != 0 {
			dataPtr += copy(data[dataPtr:], strconv.Itoa(val))
		}
		if (i+1)%puzzle.n != 0 {
			data[dataPtr] = ' '
			dataPtr++
		}
	}
	return string(data)
}

func (puzzle *Puzzle) Bytes() string {
	var data = make([]byte, 8+len(puzzle.data)*2)
	copy(data[0:6], "MAU15P")
	data[7] = 0
	data[8] = 0
	const prefixLen = 8
	for i, val := range puzzle.data {
		if val == 0 {
			data[prefixLen+i*2] = 255
		} else {
			data[prefixLen+i*2] = byte(val)
		}
		data[prefixLen+i*2+1] = 0
	}
	return string(data)
}
