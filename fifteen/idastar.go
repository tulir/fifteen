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
	"maunium.net/go/fifteen/fifteen/datastructures"
)

type path struct {
	root  *Puzzle
	nodes ds.IntStack
	moves ds.LinkedMoveStack
}

// DrawIntermediate is a function that the main program sets. It is then called
// by search() whenever it makes a move.
// Mostly useless as rendering each step in the algorithm doesn't even look fun.
var DrawIntermediate func(puzzle *Puzzle)

// FindShortestSolution uses the iterative deepening A* along with the manhattan distance as a heuristic
// to find the least number of moves required to move the puzzle into the final position.
func (puzzle *Puzzle) FindShortestSolution() []ds.Position {
	path := &path{
		root: puzzle.Copy(),
	}

	found := path.idaStar()
	if !found {
		return nil
	}
	return path.moves.Array()
}

// IDA* algorithm based on pseudocode from https://en.wikipedia.org/wiki/Iterative_deepening_A*
func (p *path) idaStar() bool {
	bound := p.root.ManhattanDistance()
	p.nodes = ds.IntStack{p.root.Hash()}
	p.moves = ds.LinkedMoveStack{}
	for {
		bound = p.search(p.root, 0, bound)
		if bound == idasFound {
			return true
		} else if bound == idasNotFound {
			return false
		}
	}
}

const idasFound = -1
const idasNotFound = 2 << 30

func (p *path) search(puzzle *Puzzle, cost, bound int) int {
	estimatedCost := cost + puzzle.ManhattanDistance()
	if estimatedCost > bound {
		return estimatedCost
	} else if puzzle.IsSolved() {
		return idasFound
	}
	min := idasNotFound
	for _, move := range puzzle.GetValidMoves() {
		// In order to save memory (and GC time), we mutate the puzzle instead of making copies.
		// MovePos returns the reverse move which we remember and apply after recursing search().
		reverse := puzzle.MovePos(move)
		hash := puzzle.Hash()
		if p.nodes.Contains(hash) {
			// Already visited this state, revert move and continue.
			puzzle.MovePos(reverse)
			continue
		}
		if DrawIntermediate != nil {
			DrawIntermediate(puzzle)
		}
		p.nodes.Push(hash)
		p.moves.Push(move)
		t := p.search(puzzle, cost+1, bound)
		// Revert move made at beginning of loop.
		puzzle.MovePos(reverse)
		if t == idasFound {
			return idasFound
		} else if t < min {
			min = t
		}
		p.moves.Pop()
		p.nodes.Remove()
		if DrawIntermediate != nil {
			DrawIntermediate(puzzle)
		}
	}
	return min
}
