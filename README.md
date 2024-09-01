# Journal
This repo is my Go learning journal.

## Check Version Used
Go - 
```bash
go version
```

Go Lint -
```bash
golangci-lint --version
```

## Installation
Go - Download and install go from [here](https://go.dev/doc/install)
Godoc - 
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
Launch godoc by -
```bash
godoc -http :8000
```
then visit [localhost:8000/pkg/](http://localhost:8000/pkg/) to access Go docs.

Errcheck -
```bash
go install github.com/kisielk/errcheck@latest
```

## Instruction 1: Go project initialization

```bash
go mod init github.com/himalayanegi10/go-with-tests
```


## Instruction 2: Running a go file

```bash
go run <filename>.go
```

## Instruction 3: Testing in Go
```bash
go test
```

## Instruction 4: Run Example in Go

```bash
go test -v
```

## Instruction 5: Benchmarking in Go

In Linux -
```bash
go test -bench=.
```

In windows -
```bash
go test -bench="."
```

Sample Output-
```bash
goos: windows
goarch: amd64
pkg: github.com/himalayanegi10/go-with-tests/src/iteration
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkRepeat-16      13069944                88.61 ns/op
PASS
ok      github.com/himalayanegi10/go-with-tests/src/iteration   1.681s
```

## Instruction 6: Coverage in Go

Coverage can be tested by - 
```bash
go test -cover
```

## Instruction 7: Error Check in Go

errcheck is used to check errors - 
``` bash
errcheck .
```

## Instruction 8: Race test in Go

Race can be tested by -
```bash
go test -race
```