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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzle_String(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.Equal(t, ` 1  2  3  4
 5  6  7  8
 9 10 11 12
13 14 15   `, puzzle.String())
}

func TestPuzzle_Binary(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.Equal(t, string([]byte{0, 0,
		1, 0, 2, 0, 3, 0, 4, 0,
		5, 0, 6, 0, 7, 0, 8, 0,
		9, 0, 10, 0, 11, 0, 12, 0,
		13, 0, 14, 0, 15, 0, 255, 0}), puzzle.Bytes())
}
