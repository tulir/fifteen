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
	flag "maunium.net/go/mauflag"
)

var input = flag.MakeFull("i", "input", "Path to read input from.", "-").String()
var output = flag.MakeFull("o", "output", "Path to output result to.", "-").String()
var generate = flag.MakeFull("g", "generate", "Generate a random solvable 15-puzzle", "false").Bool()
var wantHelp, _ = flag.MakeHelpFlag()

func main() {
	err := flag.Parse()
	if err != nil {
		fmt.Println(err)
		flag.PrintHelp()
		return
	} else if *wantHelp {
		flag.PrintHelp()
		return
	}
	fmt.Println("This program does not do anything yet.")
}
