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

func TestPuzzle_HasAllNumbers_Solved(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.True(t, puzzle.HasAllNumbers())
}

func TestPuzzle_HasAllNumbers_AlmostValid(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Set(3, 3, 0)
	assert.False(t, puzzle.HasAllNumbers())
}

func TestPuzzle_HasAllNumbers_OutOfRange1(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Set(3, 3, 16)
	assert.False(t, puzzle.HasAllNumbers())
}

func TestPuzzle_HasAllNumbers_OutOfRange2(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Set(3, 3, -1)
	assert.False(t, puzzle.HasAllNumbers())
}

func TestPuzzle_HasAllNumbers_BlankFail(t *testing.T) {
	puzzle, _ := NewPuzzle(4)
	assert.False(t, puzzle.HasAllNumbers())
}

func TestPuzzle_Solvable_Solved(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.True(t, puzzle.Solvable())
}

func TestPuzzle_Solvable_SingleInversion(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Set(2, 4, 15)
	puzzle.Set(3, 4, 14)
	assert.False(t, puzzle.Solvable())
}

func TestPuzzle_Solvable_BlankFail(t *testing.T) {
	puzzle, _ := NewPuzzle(4)
	assert.False(t, puzzle.Solvable())
}

func TestPuzzle_Solvable_Solved_9p(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(3)
	assert.True(t, puzzle.Solvable())
}

func TestPuzzle_Solvable_SingleInversion_9p(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(3)
	puzzle.Set(1, 3, 8)
	puzzle.Set(2, 3, 7)
	assert.False(t, puzzle.Solvable())
}

func TestPuzzle_Solvable_BlankFail_9p(t *testing.T) {
	puzzle, _ := NewPuzzle(3)
	assert.False(t, puzzle.Solvable())
}

func TestPuzzle_IsSolved_Solved(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.True(t, puzzle.IsSolved())
}
func TestPuzzle_IsSolved_Blank(t *testing.T) {
	puzzle, _ := NewPuzzle(4)
	assert.False(t, puzzle.IsSolved())
}

func TestPuzzle_IsSolved_SingleMove(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Move(3, 4)
	assert.False(t, puzzle.IsSolved())
}

func TestPuzzle_IsSolved_NoZero(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Set(4, 4, 16)
	assert.False(t, puzzle.IsSolved())
}
