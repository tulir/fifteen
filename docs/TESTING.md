# Testing document
Project contains unit tests for most functions and performance tests for
functions critical to the algorithm. Tests are implemented with Go's standard
testing framework.

## Unit tests
There are unit tests for most of the code, except for the IO/rendering code and
the main IDA* algorithm.

The unit tests were mostly created after the code based on what the code did or
was supposed to do.

### Running unit tests
Run tests and output coverage results:
```bash
go test ./... -coverprofile=coverage.out
```

Read coverage info in terminal:
```bash
go tool cover -func=coverage.out
```

Generate coverage report and open in browser:
```bash
go tool cover -html=coverage.out
```

Generate coverage report and save to file:
```bash
go tool cover -html=coverage.out -o file.html
```

## Performance tests
There are performance tests for functions that are critical to the algorithm,
such as the manhattan distance calculation function, the puzzle move function
and the puzzle solved check function.

### Running benchmarks
Simple benchmark:
```bash
go test ./... -bench=.
```

The benchmark will run for a minimum of 1 second by default. All the benchmarked
functions are fast enough that the benchmarks are accurate with that default.
If you still want to change the time, you can use `-benchtime=Xs` to change the
time to run the benchmarks for (replace X with number of seconds).

