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

const fnvOffsetBasis = 14695981039346656037
const fnvPrime = 1099511628211

// Hash calculates the FNV-1 hash of puzzle data.
//
// See https://en.wikipedia.org/wiki/FNV_hash_function
func (puzzle *Puzzle) Hash() (hash uint64) {
	hash = fnvOffsetBasis
	for _, val := range puzzle.data {
		hash *= fnvPrime
		hash ^= uint64(val)
	}
	return
}
