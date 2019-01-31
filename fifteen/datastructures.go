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

// IntStack is a string array with additional methods to use it like a stack.
type IntStack []uint64

// Push pushes the given string to the top of the stack.
func (s *IntStack) Push(v uint64) {
	*s = append(*s, v)
}

// Remove removes the element at the top of the stack.
func (s *IntStack) Remove() {
	*s = (*s)[:len(*s)-1]
}

// Contains checks if the stack contains the given value.
func (s *IntStack) Contains(val uint64) bool {
	for _, i := range *s {
		if i == val {
			return true
		}
	}
	return false
}

// linkedMove is an item in a LinkedMoveStack.
type linkedMove struct {
	prev *linkedMove
	move Position
	next *linkedMove
}

// LinkedMoveStack is a linked list of Positions that acts like a stack.
type LinkedMoveStack struct {
	start *linkedMove
	end   *linkedMove
	size  int
}

// Push adds the given position to this list.
func (lms *LinkedMoveStack) Push(move Position) {
	if lms.start == nil {
		lms.start = &linkedMove{prev: nil, move: move, next: nil}
		lms.end = lms.start
	} else {
		lms.end = &linkedMove{prev: lms.end, move: move, next: nil}
		lms.end.prev.next = lms.end
	}
	lms.size++
}

// Pop pops the last position in the list.
func (lms *LinkedMoveStack) Pop() {
	lms.end = lms.end.prev
	if lms.end == nil {
		lms.start = nil
	}
	lms.size--
}

// Array converts this LinkedMoveStack into an array.
func (lms *LinkedMoveStack) Array() (arr []Position) {
	arr = make([]Position, lms.size)
	i := 0
	for move := lms.start; move != nil && i < lms.size; move = move.next {
		arr[i] = move.move
		i++
	}
	return
}
