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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	flag "maunium.net/go/mauflag"

	"maunium.net/go/fifteen/fifteen"
)

var input = flag.MakeFull("i", "input", "Path to read input from. Generates random puzzle by default", "").String()
var output = flag.MakeFull("o", "output", "Path to output result to. Defaults to stdout", "-").String()
var shuffle = flag.MakeFull("s", "shuffle", "Number of moves to shuffle the puzzle with before solving.", "0").Int()
var outputFormat = flag.MakeFull("f", "format", "Output format. One of: text, json", "text").String()
var animateFormat = flag.MakeFull("a", "animate", "Animation format. One of: solution, steps", "").String()
var animateTime = flag.MakeFull("d", "duration", "Animation duration (only applicable for solution animation", "10").Int()
var randomSeed = flag.MakeFull("r", "seed", "Seed for randomization. Defaults to current time", "-1").Int64()
var wantHelp, _ = flag.MakeHelpFlag()

type JSONInput [][]int

func stderr(msg ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, msg...)
}

func main() {
	flag.SetHelpTitles("fifteen - 15-puzzle solver",
		"fifteen [-i inputpath] [-o outputpath] [-s shufflecount] [-f text|json] [-a solution|steps]")
	err := flag.Parse()
	if err != nil {
		stderr(err)
		flag.PrintHelp()
		return
	} else if *wantHelp {
		flag.PrintHelp()
		return
	}

	if *randomSeed == -1 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	var puzzle *fifteen.Puzzle
	if len(*input) == 0 {
		puzzle = fifteen.NewSolvedPuzzle(4)
	} else {
		data, err := ioutil.ReadFile(*input)
		if err != nil {
			stderr("Failed to read input file:", err)
			return
		}
		puzzle, err = fifteen.ParsePuzzle(string(data))
		if err != nil {
			var inputData JSONInput
			jsonErr := json.Unmarshal(data, &inputData)
			if jsonErr != nil {
				stderr("Input was not JSON or plaintext puzzle.")
				stderr("Plaintext parse error:", err)
				stderr("JSON parse error:", jsonErr)
				return
			}
			err = puzzle.SetData(inputData)
			if err != nil {
				stderr("Invalid array dimensions in input JSON")
				return
			}
		}
	}

	solvable := puzzle.Solvable()
	if !solvable {
		stderr("Input puzzle is not solvable!")
	}
	puzzle.Shuffle(*shuffle)

	var solution []fifteen.Position
	var duration int64
	if *animateFormat == "steps" {
		solution = solveAnimated(puzzle)
	} else {
		duration, solution = solveBenchmark(puzzle)
	}

	if *animateFormat == "solution" {
		duration := time.Duration(*animateTime) * time.Second
		animateSolution(puzzle.Copy(), solution, duration/time.Duration(len(solution)))
	}

	var buf bytes.Buffer
	if *outputFormat == "json" {
		data := JSONOutput{
			Puzzle:   puzzle.Data(),
			Clicks:   solution,
			Duration: duration,
		}
		data.PrettyJSON(&buf)
	} else {
		buf.WriteString("Puzzle:\n")
		buf.WriteString(puzzle.String())
		buf.WriteString("\n\nSolution (click coordinates):\n")
		for _, click := range solution {
			buf.WriteString(click.String())
			buf.WriteRune('\n')
		}
		if duration > 0 {
			_, _ = fmt.Fprintf(&buf, "\nDuration to solve: %s\n", formatDuration(duration))
		}
	}
	if *output == "-" {
		fmt.Print(buf.String())
	} else {
		err = ioutil.WriteFile(*output, buf.Bytes(), 0644)
		if err != nil {
			stderr("Failed to write output:", err)
		}
	}
}

func formatDuration(duration int64) string {
	if duration > 1000*1000*10 {
		ms := duration / (1000 * 1000)
		if ms > 1000 {
			seconds, ms := ms/1000, ms%1000
			if seconds > 60 {
				minutes, seconds := seconds/60, seconds%60
				return fmt.Sprintf("%d minutes and %d.%d seconds", minutes, seconds, ms)
			}
			return fmt.Sprintf("%d.%d seconds", seconds, ms)
		}
		return fmt.Sprintf("%d ms", ms)
	} else if duration > 1000 {
		return fmt.Sprintf("%3.3f µs", float64(duration)/1000.0)
	}
	return fmt.Sprintf("%d ns", duration)
}

func solveBenchmark(puzzle *fifteen.Puzzle) (int64, []fifteen.Position) {
	start := time.Now().UnixNano()
	sol := puzzle.FindShortestSolution()
	end := time.Now().UnixNano()
	return end - start, sol
}
