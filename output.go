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

package main

import (
	"bytes"
	"fmt"
	"maunium.net/go/fifteen/fifteen"
	"strconv"
)

type JSONOutput struct {
	Puzzle   [][]int            `json:"puzzle"`
	Clicks   []fifteen.Position `json:"clicks"`
	Duration int64              `json:"duration,omitempty"`
}

func (op *JSONOutput) PrettyJSON(buf *bytes.Buffer) {
	buf.WriteString("{\n")
	buf.WriteString("  \"puzzle\": [\n")
	charSize := fifteen.Digits(len(op.Puzzle) * len(op.Puzzle))
	format := "%" + strconv.Itoa(charSize) + "d"
	for y, row := range op.Puzzle {
		buf.WriteString("    [")
		for x, col := range row {
			_, _ = fmt.Fprintf(buf, format, col)
			if x < len(row)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteRune(']')
		if y != len(op.Puzzle)-1 {
			buf.WriteRune(',')
		}
		buf.WriteRune('\n')
	}
	buf.WriteString("  ],\n")
	buf.WriteString("  \"clicks\": [\n")
	for i, click := range op.Clicks {
		_, _ = fmt.Fprintf(buf, `    {"x": %d, "y": %d}`, click.X, click.Y)
		if i != len(op.Clicks)-1 {
			buf.WriteRune(',')
		}
		buf.WriteRune('\n')
	}
	buf.WriteString("  ]")
	if op.Duration > 0 {
		buf.WriteString(",\n")
		_, _ = fmt.Fprintf(buf, "  \"duration\": %d", op.Duration)
	}
	buf.WriteRune('\n')
	buf.WriteString("}\n")
}
