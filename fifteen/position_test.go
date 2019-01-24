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

func TestPosition_Valid(t *testing.T) {
	assert.True(t, Position{1, 1}.Valid(4))
	assert.True(t, Position{4, 4}.Valid(4))
	assert.True(t, Position{2, 3}.Valid(4))
	assert.False(t, Position{0, 3}.Valid(4))
	assert.False(t, Position{2, 0}.Valid(4))
	assert.False(t, Position{5, 1}.Valid(4))
	assert.False(t, Position{4, 5}.Valid(4))
	assert.True(t, Position{4, 5}.Valid(5))
}

func TestPosition_AllMovesArray(t *testing.T) {
	moves := Position{1, 1}.AllMovesArray()
	assert.Contains(t, moves, Position{0, 1})
	assert.Contains(t, moves, Position{2, 1})
	assert.Contains(t, moves, Position{1, 0})
	assert.Contains(t, moves, Position{1, 2})
}

func TestPosition_ValidMoves(t *testing.T) {
	moves := Position{1, 1}.ValidMoves(6)
	assert.NotContains(t, moves, Position{0, 1})
	assert.Contains(t, moves, Position{2, 1})
	assert.NotContains(t, moves, Position{1, 0})
	assert.Contains(t, moves, Position{1, 2})
	moves = Position{4, 3}.ValidMoves(4)
	assert.Contains(t, moves, Position{3, 3})
	assert.NotContains(t, moves, Position{5, 3})
	assert.Contains(t, moves, Position{4, 2})
	assert.Contains(t, moves, Position{4, 4})
}

func TestPosition_String(t *testing.T) {
	assert.Equal(t, "122, 124", Position{122, 124}.String())
}