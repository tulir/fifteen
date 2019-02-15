# Testing document
Project contains unit tests for most functions and performance tests for
functions critical to the algorithm. Tests are implemented with Go's standard
testing framework.



## Running performance tests
Standard Go performance tests

Simple run:
```bash
go test ./... -bench=.
```

The benchmark will run for a minimum of 1 second by default. All the benchmarked
functions are fast enough that the benchmarks are accurate with that default.
If you still want to change the time, you can use `-benchtime=Xs` to change the
time to run the benchmarks for (replace X with number of seconds).

## Running unit tests
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
