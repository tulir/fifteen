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

package ds

// IntStack is a string array with additional methods to use it like a stack.
//
// It's used for remembering the hashes of the positions the IDA* algorithm has
// visited in the current path to prevent loops.
// TODO array might not be the best underlying data structure as we push/pop as
//      much as in LinkedMoveStack (though we also iterate over it a lot).
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
