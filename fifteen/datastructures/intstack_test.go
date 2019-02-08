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

func TestIntStack_All(t *testing.T) {
	is := IntStack{}
	is.Push(1234)
	assert.Len(t, is, 1)
	is.Push(1234)
	assert.Len(t, is, 2)
	is.Push(1235)
	assert.Len(t, is, 3)
	assert.True(t, is.Contains(1235))
	is.Remove()
	assert.Len(t, is, 2)
	assert.False(t, is.Contains(1235))
	is.Push(1236)
	assert.Len(t, is, 3)
	assert.Equal(t, []uint64{1234, 1234, 1236}, []uint64(is))
}
