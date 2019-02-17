# Building

```bash
git submodule init --update --recursive
go generate ./..
go test -bench=. -benchtime=5s benchmark_test.go
```