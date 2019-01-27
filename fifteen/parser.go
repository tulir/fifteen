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

// ParsePuzzle parses a integer grid from a string into a puzzle.
func ParsePuzzle(input string) (*Puzzle, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	puzzle, err := NewPuzzle(len(rows))
	if err != nil {
		return nil, err
	}
	for y, row := range rows {
		cells := strings.Fields(row)
		if len(cells) != len(rows) {
			return nil, fmt.Errorf("size mismatch: row %d has a different amount columns than the input has rows", y)
		}
		for x, cell := range cells {
			val, err := strconv.Atoi(cell)
			if err != nil {
				return nil, fmt.Errorf("parse error at row %d column %d: %v", y, x, err)
			} else if val >= puzzle.n*puzzle.n {
				return nil, fmt.Errorf("value too large at row %d column %d", y, x)
			} else if val < 0 {
				return nil, fmt.Errorf("value too small at row %d column %d", y, x)
			}
			puzzle.Set(x+1, y+1, val)
		}
	}
	return puzzle, nil
}
