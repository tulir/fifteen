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

type path struct {
	root  *Puzzle
	nodes IntStack
	moves LinkedMoveStack
}

var DrawIntermediate func(puzzle *Puzzle)

// FindShortestSolution uses the iterative deepening A* along with the manhattan distance as a heuristic
// to find the least number of moves required to move the puzzle into the final position.
func (puzzle *Puzzle) FindShortestSolution() []Position {
	path := &path{
		root: puzzle.Copy(),
	}

	found := path.idaStar()
	if !found {
		return nil
	}
	return path.moves.Array()
}

func (p *path) idaStar() bool {
	bound := p.root.ManhattanDistance()
	p.nodes = IntStack{p.root.Hash()}
	p.moves = LinkedMoveStack{}
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
		reverse := puzzle.MovePos(move)
		hash := puzzle.Hash()
		if p.nodes.Contains(hash) {
			puzzle.MovePos(reverse)
			continue
		}
		if DrawIntermediate != nil {
			DrawIntermediate(puzzle)
		}
		p.nodes.Push(hash)
		p.moves.Push(move)
		t := p.search(puzzle, cost+1, bound)
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
