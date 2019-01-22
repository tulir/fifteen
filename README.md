# 15-puzzle solver
[![Build Status](https://travis-ci.org/tulir/hy-tiralabra-15p.svg?branch=master)](https://travis-ci.org/tulir/hy-tiralabra-15p)
[![Maintainability](https://api.codeclimate.com/v1/badges/bb4f556a0d3fa2d3f500/maintainability)](https://codeclimate.com/github/tulir/hy-tiralabra-15p/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/bb4f556a0d3fa2d3f500/test_coverage)](https://codeclimate.com/github/tulir/hy-tiralabra-15p/test_coverage)

Data structures and algorithms project, 15-puzzle solver. Written in Go.

* [Project definition](docs/PROJECT_DEFINITION.md)
* Week reports: [1](docs/WEEK_1.md), [2](docs/WEEK_2.md)

## Instructions
### Installation
0. Install [Go](https://golang.org/) 1.11 or higher
1. Clone the repository (`git clone git@github.com:tulir/hy-tiralabra-15p`)
2. Fetch dependencies and compile with `go get`

The compiled binary should now exist at `$GOPATH/bin/fifteen`. `$GOPATH`
defaults to `$HOME/go`. You can now `go install` to recompile the binary at
`$GOPATH/bin/fifteen` or `go build` to create a binary in your working
directory.

You should run `go get` again after pulling changes from Git to make sure
dependencies are up to date.

### Usage
TODO 

### Running tests
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
