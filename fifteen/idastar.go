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

// IntStack is a string array with additional methods to use it like a stack.
type IntStack []int

// Push pushes the given string to the top of the stack.
func (s *IntStack) Push(v int) {
	*s = append(*s, v)
}

// Pop removes and returns the element at the top of the stack.
func (s *IntStack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// Remove removes the element at the top of the stack.
func (s *IntStack) Remove() {
	*s = (*s)[:len(*s)-1]
}

// Contains checks if the stack contains the given value.
func (s *IntStack) Contains(val int) bool {
	for _, str := range *s {
		if str == val {
			return true
		}
	}
	return false
}

type path struct {
	nodes IntStack
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
		nodes: IntStack{start.puzzle.Hash()},
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
		if t == idaFound {
			return true
		} else if t == idaNotFound {
			return false
		}
		p.bound = t
	}
}

const idaFound = -1
const idaNotFound = 2 << 30

func (p *path) search() int {
	estimatedCost := p.cur.cost + p.cur.estimatedCostToGoal()
	if estimatedCost > p.bound {
		return estimatedCost
	} else if p.cur.puzzle.IsSolved() {
		return idaFound
	}
	min := idaNotFound
	prevCur := p.cur
	for _, succ := range p.cur.successors() {
		hash := succ.puzzle.Hash()
		if p.nodes.Contains(hash) {
			continue
		}
		p.cur = succ
		if DrawIntermediate != nil {
			DrawIntermediate(p.cur.puzzle)
		}
		p.nodes.Push(hash)
		t := p.search()
		if t == idaFound {
			return idaFound
		} else if t < min {
			min = t
		}
		p.cur = prevCur
		p.nodes.Remove()
		if DrawIntermediate != nil {
			DrawIntermediate(p.cur.puzzle)
		}
	}
	return min
}

// node is a node in the IDA* search stack.
//
// It contains the puzzle state, cost to reach the state, the move made from
// the previous state and a pointer to the previous node.
type node struct {
	puzzle *Puzzle
	prev   *node
	move   Position
	cost   int
}

// successors returns the list of states that can follow the state in the node.
func (n *node) successors() (nodes []*node) {
	moves := n.puzzle.GetValidMoves()
	nodes = make([]*node, len(moves))
	for i, move := range moves {
		newPuzzle := n.puzzle.Copy()
		if !newPuzzle.Move(move.X, move.Y) {
			fmt.Println("Move", move, "failed!")
		}
		nodes[i] = &node{
			puzzle: newPuzzle,
			prev:   n,
			move:   move,
			cost:   n.cost + 1,
		}
	}
	return
}

// estimatedCostToGoal is the heuristic used by the (ID)A* algorithm.
func (n *node) estimatedCostToGoal() int {
	return n.puzzle.ManhattanDistance()
}
