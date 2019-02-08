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
	"github.com/gdamore/tcell"
	"maunium.net/go/fifteen/fifteen"
	"maunium.net/go/fifteen/fifteen/datastructures"
	"maunium.net/go/fifteen/fifteen/util"
	"strconv"
	"time"
)

const (
	CornerTopLeft     = '┌'
	CornerBottomLeft  = '└'
	CornerTopRight    = '┐'
	CornerBottomRight = '┘'
	LineHorizontal    = '─'
	LineVertical      = '│'
	THorizontalDown   = '┬'
	THorizontalUp     = '┴'
	TVerticalRight    = '├'
	TVerticalLeft     = '┤'
	Middle            = '┼'
)

func renderGrid(screen tcell.Screen, size int) {
	cellSize := util.Digits(size*size) + 1
	right := size * cellSize
	bottom := size * 2
	width, height := screen.Size()
	offsetX := (width - right) / 2
	offsetY := (height - bottom) / 2
	screen.SetContent(offsetX, offsetY, CornerTopLeft, nil, tcell.StyleDefault)
	screen.SetContent(offsetX+right, offsetY, CornerTopRight, nil, tcell.StyleDefault)
	screen.SetContent(offsetX, offsetY+bottom, CornerBottomLeft, nil, tcell.StyleDefault)
	screen.SetContent(offsetX+right, offsetY+bottom, CornerBottomRight, nil, tcell.StyleDefault)
	for x := 1; x < right; x++ {
		if x%cellSize != 0 {
			screen.SetContent(offsetX+x, offsetY, LineHorizontal, nil, tcell.StyleDefault)
			screen.SetContent(offsetX+x, offsetY+bottom, LineHorizontal, nil, tcell.StyleDefault)
		} else {
			screen.SetContent(offsetX+x, offsetY, THorizontalDown, nil, tcell.StyleDefault)
			screen.SetContent(offsetX+x, offsetY+bottom, THorizontalUp, nil, tcell.StyleDefault)
			for y := 1; y < bottom; y++ {
				if y%2 == 1 {
					screen.SetContent(offsetX+x, offsetY+y, LineVertical, nil, tcell.StyleDefault)
				} else {
					screen.SetContent(offsetX+x, offsetY+y, Middle, nil, tcell.StyleDefault)
				}
			}
		}
	}
	for y := 1; y < bottom; y++ {
		if y%2 == 1 {
			screen.SetContent(offsetX, offsetY+y, LineVertical, nil, tcell.StyleDefault)
			screen.SetContent(offsetX+right, offsetY+y, LineVertical, nil, tcell.StyleDefault)
		} else {
			screen.SetContent(offsetX, offsetY+y, TVerticalRight, nil, tcell.StyleDefault)
			screen.SetContent(offsetX+right, offsetY+y, TVerticalLeft, nil, tcell.StyleDefault)
			for x := 1; x < right; x++ {
				if x%cellSize != 0 {
					screen.SetContent(offsetX+x, offsetY+y, LineHorizontal, nil, tcell.StyleDefault)
				}
			}
		}
	}
}

func renderPuzzle(screen tcell.Screen, puzzle *fifteen.Puzzle) {
	cellSize := util.Digits(puzzle.Size()*puzzle.Size()) + 1
	right := puzzle.Size() * cellSize
	bottom := puzzle.Size() * 2
	width, height := screen.Size()
	offsetX := (width-right)/2 + 1
	offsetY := (height-bottom)/2 + 1
	for y, row := range puzzle.Data() {
		for x, cell := range row {
			str := strconv.Itoa(cell)
			if cell == 0 {
				str = ""
			}
			for i, char := range str {
				screen.SetContent(offsetX+x*cellSize+i, offsetY+(y*2), char, nil, tcell.StyleDefault)
			}
			for i := len(str); i < cellSize-1; i++ {
				screen.SetContent(offsetX+x*cellSize+i, offsetY+(y*2), ' ', nil, tcell.StyleDefault)
			}
		}
	}
	screen.Show()
}

func solveAnimated(puzzle *fifteen.Puzzle) []ds.Position {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = screen.Init()
	if err != nil {
		panic(err)
	}
	renderGrid(screen, puzzle.Size())
	renderPuzzle(screen, puzzle)
	fifteen.DrawIntermediate = func(step *fifteen.Puzzle) {
		renderPuzzle(screen, step)
	}
	time.Sleep(1 * time.Second)
	sol := puzzle.FindShortestSolution()
	fifteen.DrawIntermediate = nil
	time.Sleep(2 * time.Second)
	screen.Fini()
	return sol
}

func animateSolution(puzzle *fifteen.Puzzle, moves []ds.Position, delay time.Duration) {
	screen, _ := tcell.NewScreen()
	_ = screen.Init()
	renderGrid(screen, puzzle.Size())
	renderPuzzle(screen, puzzle)
	time.Sleep(1 * time.Second)
	for _, move := range moves {
		time.Sleep(delay)
		puzzle.MovePos(move)
		renderPuzzle(screen, puzzle)
	}
	time.Sleep(2 * time.Second)
	screen.Fini()
}
