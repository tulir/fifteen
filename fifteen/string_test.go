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

func TestPuzzle_String_3(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(3)
	assert.Equal(t, `1 2 3
4 5 6
7 8 -`, puzzle.String())
}

func TestPuzzle_String_4(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.Equal(t, ` 1  2  3  4
 5  6  7  8
 9 10 11 12
13 14 15  -`, puzzle.String())
}

func TestPuzzle_String_5(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(5)
	assert.Equal(t, ` 1  2  3  4  5
 6  7  8  9 10
11 12 13 14 15
16 17 18 19 20
21 22 23 24  -`, puzzle.String())
}

func TestPuzzle_String_11(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(11)
	assert.Equal(t, `  1   2   3   4   5   6   7   8   9  10  11
 12  13  14  15  16  17  18  19  20  21  22
 23  24  25  26  27  28  29  30  31  32  33
 34  35  36  37  38  39  40  41  42  43  44
 45  46  47  48  49  50  51  52  53  54  55
 56  57  58  59  60  61  62  63  64  65  66
 67  68  69  70  71  72  73  74  75  76  77
 78  79  80  81  82  83  84  85  86  87  88
 89  90  91  92  93  94  95  96  97  98  99
100 101 102 103 104 105 106 107 108 109 110
111 112 113 114 115 116 117 118 119 120   -`, puzzle.String())
}

func TestPuzzle_Hash_4(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(4)
	assert.Equal(t, uint64(0x7c84dc9477851775), puzzle.Hash())
}

func TestPuzzle_Hash_11(t *testing.T) {
	puzzle, _ := NewSolvedPuzzle(11)
	assert.Equal(t, uint64(0xbae28509ef48216f), puzzle.Hash())
}
