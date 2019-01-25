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
	"fmt"
)

type StringStack []string

func (s *StringStack) Push(v string) {
	*s = append(*s, v)
}

func (s *StringStack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *StringStack) Contains(val string) bool {
	for _, str := range *s {
		if str == val {
			return true
		}
	}
	return false
}

type path struct {
	nodes StringStack
	start *node
	cur   *node
	bound int
}

var DrawIntermediate func(puzzle *Puzzle)

// FindShortestSolution uses the iterative deepening A* along with the manhattan distance as a heuristic
// to find the least number of moves required to move the puzzle into the final position.
func (puzzle *Puzzle) FindShortestSolution() []Position {
	start := &node{
		puzzle: puzzle.Copy(),
		prev:   nil,
		cost:   0,
	}
	path := &path{
		start: start,
		cur:   start,
		bound: start.estimatedCostToGoal(),
		nodes: StringStack{start.puzzle.Bytes()},
	}

	found := path.idaStar()
	if !found {
		return nil
	}

	solution := make([]Position, len(path.nodes)-1)
	node := path.cur
	i := len(solution)
	for node != nil && node != path.start {
		i--
		solution[i] = node.move
		node = node.prev
	}
	return solution
}

func (p *path) idaStar() bool {
	for {
		t := p.search()
		if t == Found {
			return true
		} else if t == NotFound {
			return false
		}
		p.bound = t
	}
}

const Found = -1
const NotFound = 2 << 30

func (p *path) search() int {
	estimatedCost := p.cur.cost + p.cur.estimatedCostToGoal()
	if estimatedCost > p.bound {
		return estimatedCost
	} else if p.cur.puzzle.IsSolved() {
		return Found
	}
	min := NotFound
	prevCur := p.cur
	for _, succ := range p.cur.successors() {
		if p.nodes.Contains(succ.puzzle.Bytes()) {
			continue
		}
		p.cur = succ
		if DrawIntermediate != nil {
			DrawIntermediate(p.cur.puzzle)
		}
		p.nodes.Push(p.cur.puzzle.Bytes())
		t := p.search()
		if t == Found {
			return Found
		} else if t < min {
			min = t
		}
		p.cur = prevCur
		p.nodes.Pop()
		if DrawIntermediate != nil {
			DrawIntermediate(p.cur.puzzle)
		}
	}
	return min
}

type node struct {
	puzzle *Puzzle
	prev   *node
	move   Position
	cost   int
}

func (n *node) successors() (nodes []*node) {
	for _, move := range n.puzzle.GetValidMoves() {
		newPuzzle := n.puzzle.Copy()
		if !newPuzzle.Move(move.X, move.Y) {
			fmt.Println("Move", move, "failed!")
		}
		nodes = append(nodes, &node{
			puzzle: newPuzzle,
			prev:   n,
			move:   move,
			cost:   n.cost + 1,
		})
	}
	return
}

func (n *node) estimatedCostToGoal() int {
	return n.puzzle.ManhattanDistance()
}
