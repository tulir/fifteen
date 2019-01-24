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
	"math/rand"
	"testing"
)

func TestNewRandomPuzzle(t *testing.T) {
	// Seed chosen so that the initial randomization won't provide a solvable puzzle.
	rand.Seed(1236)
	puzzle := NewRandomPuzzle(4)
	assert.Equal(t, [][]int{
		{2, 12, 5, 15},
		{9, 14, 1, 13},
		{11, 3, 8, 0},
		{6, 4, 7, 10},
	}, puzzle.Data())
	assert.Equal(t, Position{4, 3}, puzzle.blank)
}

func TestPuzzle_Shuffle(t *testing.T) {
	puzzle := NewSolvedPuzzle(4)
	assert.True(t, puzzle.Solvable())
	assert.True(t, puzzle.IsSolved())
	puzzle.Shuffle(200)
	assert.True(t, puzzle.Solvable())
	assert.False(t, puzzle.IsSolved())
}
