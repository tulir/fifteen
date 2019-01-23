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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePuzzle_Valid(t *testing.T) {
	puzzle, err := ParsePuzzle(
		"1 2 3 4\n" +
			"5 6 7 8\n" +
			"9 10 11 12\n" +
			"13 14 15 0")
	assert.NoError(t, err)
	assert.NotNil(t, puzzle)
	assert.Equal(t, puzzle.data, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0})
}

func TestParsePuzzle_SizeMismatch(t *testing.T) {
	puzzle, err := ParsePuzzle(
		"1 2 3 4\n" +
			"5 6 7 8 9\n" +
			"10 11 12\n" +
			"13 14 16 0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "size mismatch")
	assert.Nil(t, puzzle)
}

func TestParsePuzzle_IntegerParseError(t *testing.T) {
	puzzle, err := ParsePuzzle(
		"1 2 3 4\n" +
			"5 6 7 8\n" +
			"9 foo 11 12\n" +
			"13 14 15 0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parse error")
	assert.Nil(t, puzzle)
}

func TestParsePuzzle_ValueTooLarge(t *testing.T) {
	puzzle, err := ParsePuzzle(
		"1 2 3 4\n" +
			"5 6 99 8\n" +
			"9 10 11 12\n" +
			"13 14 16 0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value too large")
	assert.Nil(t, puzzle)
}

func TestParsePuzzle_ValueTooSmall(t *testing.T) {
	puzzle, err := ParsePuzzle(
		"1 -2 3 4\n" +
			"5 6 7 8\n" +
			"9 10 11 12\n" +
			"13 14 16 0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value too small")
	assert.Nil(t, puzzle)
}
