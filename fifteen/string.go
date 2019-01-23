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
	"strconv"
	"strings"
)

func digits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func (puzzle *Puzzle) String() string {
	var builder strings.Builder
	maxLen := digits(len(puzzle.data)-1)
	format := "%" + strconv.Itoa(maxLen) + "d "
	for i, val := range puzzle.data {
		if i != 0 && i%puzzle.n == 0 {
			builder.WriteRune('\n')
		}
		if val == 0 {
			_, _ = fmt.Fprint(&builder, strings.Repeat(" ", maxLen+1))
		} else {
			_, _ = fmt.Fprintf(&builder, format, val)
		}
	}
	return builder.String()
}
