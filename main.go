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
	"maunium.net/go/fifteen/fifteen/datastructures"
	"os"
	"runtime/pprof"
	"time"

	flag "maunium.net/go/mauflag"

	"maunium.net/go/fifteen/fifteen"
)

var input = flag.MakeFull("i", "input", "Path to read input from. Generates random puzzle by default", "").String()
var output = flag.MakeFull("o", "output", "Path to output result to. Defaults to stdout", "-").String()
var randomize = flag.MakeFull("r", "randomize", "Randomization mode. One of: shuffle, random", "shuffle").String()
var shuffle = flag.MakeFull("n", "shuffle", "Number of moves to shuffle the puzzle with before solving.", "0").Int()
var outputFormat = flag.MakeFull("f", "format", "Output format. One of: text, json", "text").String()
var animateFormat = flag.MakeFull("a", "animate", "Animation format. One of: solution, steps", "").String()
var animateTime = flag.MakeFull("d", "duration", "Animation duration (only applicable for solution animation", "10").Int()
var randomSeed = flag.MakeFull("s", "seed", "Seed for randomization. Defaults to current time", "-1").Int64()
var size = flag.MakeFull("w", "size", "Size of game. Only applicable when generating puzzle.", "4").Int()
var profile = flag.MakeFull("p", "profile", "File to write CPU profile to", "").String()
var wantHelp, _ = flag.MakeHelpFlag()

type JSONInput [][]int

func stderr(msg ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, msg...)
}

func readFlags() {
	flag.SetHelpTitles("fifteen - 15-puzzle solver",
		"fifteen [-i inputpath] [-o outputpath] [-s shufflecount] [-f text|json] [-a solution|steps]")
	err := flag.Parse()
	if err != nil {
		stderr(err)
		flag.PrintHelp()
		os.Exit(1)
	} else if *wantHelp {
		flag.PrintHelp()
		os.Exit(0)
	}

	if *randomize != "random" && *shuffle == 0 && len(*input) == 0 {
		stderr("No shuffling or input given")
		flag.PrintHelp()
		os.Exit(2)
	}

	if *size < 3 || *size > 15 {
		stderr("Puzzle size must be within 3 and 15.")
		os.Exit(3)
	}

	if *randomSeed == -1 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}
}

func readInput() (puzzle *fifteen.Puzzle) {
	if *randomize == "random" {
		puzzle, _ = fifteen.NewRandomPuzzle(*size)
	} else if len(*input) == 0 {
		puzzle, _ = fifteen.NewSolvedPuzzle(*size)
	} else {
		data, err := ioutil.ReadFile(*input)
		if err != nil {
			stderr("Failed to read input file:", err)
			os.Exit(10)
		}
		puzzle, err = fifteen.ParsePuzzle(string(data))
		if err != nil {
			var inputData JSONInput
			jsonErr := json.Unmarshal(data, &inputData)
			if jsonErr != nil {
				stderr("Input was not JSON or plaintext puzzle.")
				stderr("Plaintext parse error:", err)
				stderr("JSON parse error:", jsonErr)
				os.Exit(11)
			}
			puzzle, _ = fifteen.NewPuzzle(len(inputData))
			err = puzzle.SetData(inputData)
			if err != nil {
				stderr("Invalid array dimensions in input JSON")
				os.Exit(12)
			}
		}
	}

	solvable := puzzle.Solvable()
	if !solvable {
		stderr("Input puzzle is not solvable!")
		os.Exit(15)
	}
	if *randomize == "shuffle" {
		puzzle.Shuffle(*shuffle)
	}

	if puzzle.IsSolved() {
		stderr("Puzzle is already solved")
		os.Exit(16)
	}
	return
}

func solve(puzzle *fifteen.Puzzle) (solution []ds.Position, duration int64) {
	if *animateFormat == "steps" {
		solution = solveAnimated(puzzle)
	} else {
		if len(*profile) > 0 {
			file, err := os.OpenFile(*profile, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				stderr("Failed to open profile file:", err)
				os.Exit(20)
			}
			err = pprof.StartCPUProfile(file)
			if err != nil {
				stderr("Failed to start profiling:", err)
				os.Exit(21)
			}
		}
		duration, solution = solveBenchmark(puzzle)
		pprof.StopCPUProfile()
	}

	if *animateFormat == "solution" {
		duration := time.Duration(*animateTime) * time.Second
		animateSolution(puzzle.Copy(), solution, duration/time.Duration(len(solution)))
	}
	return
}

func printOutput(puzzle *fifteen.Puzzle, solution []ds.Position, duration int64) {
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
		_, _ = fmt.Fprintf(&buf, "\n\nSolution (%d moves):\n", len(solution))
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
		err := ioutil.WriteFile(*output, buf.Bytes(), 0644)
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
	}
	return fmt.Sprintf("%3.3f µs", float64(duration)/1000.0)
}

func solveBenchmark(puzzle *fifteen.Puzzle) (int64, []ds.Position) {
	start := time.Now().UnixNano()
	sol := puzzle.FindShortestSolution()
	end := time.Now().UnixNano()
	return end - start, sol
}

func main() {
	readFlags()
	puzzle := readInput()
	solution, duration := solve(puzzle)
	printOutput(puzzle, solution, duration)
}
