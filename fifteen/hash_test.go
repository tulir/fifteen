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

func TestPuzzle_Hash_4(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.Equal(t, uint64(0x7c84dc9477851775), puzzle.Hash())
}

func TestPuzzle_Hash_11(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(11)
	assert.Equal(t, uint64(0xbae28509ef48216f), puzzle.Hash())
}

func BenchmarkPuzzle_Hash_4(b *testing.B) {
	puzzle, _ := NewSolvedPuzzle(4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		puzzle.Hash()
	}
}
