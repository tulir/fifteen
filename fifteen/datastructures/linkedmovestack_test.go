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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedMoveStack_All(t *testing.T) {
	lms := LinkedMoveStack{}
	lms.Push(Position{1, 2})
	lms.Push(Position{2, 2})
	lms.Push(Position{2, 3})
	lms.Push(Position{3, 3})
	assert.Equal(t, []Position{{1, 2}, {2, 2}, {2, 3}, {3, 3}}, lms.Array())
	lms.Pop()
	assert.Equal(t, []Position{{1, 2}, {2, 2}, {2, 3}}, lms.Array())
	lms.Push(Position{2, 4})
	assert.Equal(t, []Position{{1, 2}, {2, 2}, {2, 3}, {2, 4}}, lms.Array())
	lms.Pop()
	lms.Pop()
	lms.Pop()
	lms.Pop()
	assert.Empty(t, lms.Array())
}
