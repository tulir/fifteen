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

func TestPuzzle_ManhattanDistance_Solved(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.Zero(t, puzzle.ManhattanDistance())
}

func TestPuzzle_ManhattanDistance_ThreeMoves(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Move(4, 3)
	puzzle.Move(3, 3)
	puzzle.Move(3, 4)
	// Three moves misplace four tiles when the empty spot is included
	assert.Equal(t, 4, puzzle.ManhattanDistance())
}

func BenchmarkPuzzle_ManhattanDistance_Solved(b *testing.B) {
	puzzle, _ := NewSolvedPuzzle(4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		puzzle.ManhattanDistance()
	}
}

func BenchmarkPuzzle_ManhattanDistance_ThreeMoves(b *testing.B) {
	puzzle, _ := NewSolvedPuzzle(4)
	puzzle.Move(4, 3)
	puzzle.Move(3, 3)
	puzzle.Move(3, 4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		puzzle.ManhattanDistance()
	}
}

func BenchmarkPuzzle_ManhattanDistance_Random(b *testing.B) {
	puzzle, _ := NewRandomPuzzle(4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		puzzle.ManhattanDistance()
	}
}
