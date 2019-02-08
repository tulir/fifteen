# Testing document
Project contains unit tests for most functions and performance tests for
functions critical to the algorithm.

## Running performance tests
Standard Go performance tests

`go test ./... -bench=.`

// TODO full instructions

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
