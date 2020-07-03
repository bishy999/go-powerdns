# Contributing

If you submit a pull request, please keep the following guidelines in mind:

1. Code should be `go fmt` compliant.
2. Types, structs and funcs should be documented.
3. Tests pass.

## Getting set up

Assuming your `$GOPATH` is set up according to your desires, run:

```go
 go get github.com/bishy999/go-powerdns
```

Enables module-mode if any go.mod is found

```go
export GO111MODULE=on
```


## Branchingffore

Create a feature branch to make your contributions

```git

git branch feature-xxxx
git checkout feature-xxxx

```
## Running tests

When working on code in this repository, tests can be run via:

```go
go test -v ./pkg/...
```


Specific test can be targeted

```go
go test -run TestHostDoesNotExist ./pkg/powerdns/
go test -run TestInvalidCredentials/no_username -v ./pkg/powerdns/
```

## Running tests with coverage

When working on code in this repository, tests can be run via:

```go
go test -cover ./pkg/...
```

## Running benchmark tests

```go
go test -bench=BenchmarkConnectionInfo_CreateHost ./pkg/powerdns/

```

## Running benchmark tests and analyzing in Flamegraphs

```go

# Run specific benchmark
# cpu
go test -bench=BenchmarkConnectionInfo_CreateHost ./pkg/powerdns/ -cpuprofile cpuprof_createhost.out
# memory
go test -bench=BenchmarkConnectionInfo_CreateHost ./pkg/powerdns/ -memprofile memprof_createhost.out


# Analyze results on the cli
go tool pprof api.test cpuprof_createhost.out
go tool pprof --alloc_space api.test memprof_createhost.out

# Analyze results on ui
go tool pprof -http=: cpuprof_createhost.out
go tool pprof --alloc_space -http=: memprof_createhost.out

```
# Run go static analysis
```go 
go install github.com/golangci/golangci-lint/cmd/golangci-lint

golangci-lint run
```


Once feature is ready,tests are passing and commits to branch have been made create a pull request


## Merged to master
Once succesfully merged to master and approved the reviewer will create a tag based on the last commit. Tagging is based on [Semantic Versioning](https://semver.org/)
```git
git tag -a vx.y.z -m "add feature x"
git tag -a -f vx.y.z 1d258d20da4ba97fcd19a7c7c5f0af6a3638eec1
git push --tags
```