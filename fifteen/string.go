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
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func Digits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func (puzzle *Puzzle) String() string {
	var builder strings.Builder
	maxLen := Digits(len(puzzle.data) - 1)
	format := "%" + strconv.Itoa(maxLen) + "d"
	for i, val := range puzzle.data {
		if i != 0 && i%puzzle.n == 0 {
			builder.WriteRune('\n')
		}
		if val == 0 {
			_, _ = fmt.Fprint(&builder, strings.Repeat(" ", maxLen))
		} else {
			_, _ = fmt.Fprintf(&builder, format, val)
		}
		if (i+1)%puzzle.n != 0 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func (puzzle *Puzzle) Bytes() string {
	var builder bytes.Buffer
	if len(puzzle.data) > 254 {
		separator := []byte{0, 0, 0, 0}
		builder.WriteByte(0)
		builder.WriteByte(1)
		for _, val := range puzzle.data {
			if val == 0 {
				builder.Write([]byte{255, 255, 255, 255})
			} else {
				_ = binary.Write(&builder, binary.LittleEndian, val)
			}
			builder.Write(separator)
		}
	} else {
		builder.WriteByte(0)
		builder.WriteByte(0)
		for _, val := range puzzle.data {
			if val == 0 {
				builder.WriteByte(255)
			} else {
				builder.WriteByte(byte(val))
			}
			builder.WriteByte(0)
		}
	}
	return builder.String()
}
