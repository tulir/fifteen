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

func TestNewPuzzle(t *testing.T) {
	puzzle := NewPuzzle(4)
	assert.NotNil(t, puzzle)
	assert.Len(t, puzzle.data, 16)
	for _, val := range puzzle.data {
		assert.Zero(t, val)
	}
}

func TestNewSolvedPuzzle(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.NotNil(t, puzzle)
	assert.Len(t, puzzle.data, 16)
	for i, val := range puzzle.data {
		if i == 15 {
			assert.Zero(t, val)
		} else {
			assert.Equal(t, i+1, val)
		}
	}
}

func TestPuzzle_Copy(t *testing.T) {
	puzzle1 := NewPuzzle(4)
	puzzle2 := puzzle1.Copy()
	assert.Zero(t, puzzle1.data[10])
	assert.Zero(t, puzzle2.data[10])
	puzzle2.data[10] = 5
	assert.Zero(t, puzzle1.data[10])
	assert.Equal(t, 5, puzzle2.data[10])
}

func TestPuzzle_Get(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.Equal(t, 5, puzzle.Get(1, 2))
	assert.Equal(t, 0, puzzle.Get(4, 4))
	assert.Equal(t, 1, puzzle.Get(1, 1))
	assert.Equal(t, -1, puzzle.Get(0, 1))
	assert.Equal(t, -1, puzzle.Get(1, -1))
	assert.Equal(t, -1, puzzle.Get(5, 1))
	assert.Equal(t, -1, puzzle.Get(1, 5))
}

func TestPuzzle_Set(t *testing.T) {
	puzzle := NewPuzzle(4)
	for _, val := range puzzle.data {
		assert.Zero(t, val)
	}
	puzzle.Set(5, 5, 5)
	for _, val := range puzzle.data {
		assert.Zero(t, val)
	}
	puzzle.Set(3, 3, 5)
	assert.Equal(t, 5, puzzle.data[10])
}

func TestPuzzle_Coordinates(t *testing.T) {
	puzzle := NewPuzzle(4)
	x, y := puzzle.Coordinates(9)
	assert.Equal(t, 1, x)
	assert.Equal(t, 3, y)
}

func TestPuzzle_Move(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.False(t, puzzle.Move(-1, -1))
	assert.True(t, puzzle.Move(3, 4))
	assert.Zero(t, puzzle.Get(3, 4))
	assert.Equal(t, 15, puzzle.Get(4, 4))
	assert.False(t, puzzle.Move(3, 4))
	assert.True(t, puzzle.Move(3, 3))
	assert.Zero(t, puzzle.Get(3, 3))
	assert.Equal(t, 11, puzzle.Get(3, 4))
	assert.True(t, puzzle.Move(4, 3))
	assert.Zero(t, puzzle.Get(4, 3))
	assert.Equal(t, 12, puzzle.Get(3, 3))
	assert.True(t, puzzle.Move(4, 4))
	assert.Zero(t, puzzle.Get(4, 4))
	assert.Equal(t, 15, puzzle.Get(4, 3))
}
