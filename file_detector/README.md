# Building

```bash
git submodule init --update --recursive
go generate ./..
go test -bench=. -benchtime=5s benchmark_test.go
```

# Benchmark

```bash
go generate ./... && go test -bench=. -benchtime=5s benchmark_test.go
Running Suite: Results
======================
Random Seed: 1550431761
Will run 1 of 1 specs

•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
goos: darwin
goarch: amd64
BenchmarkMatchTar-4    	50000000	       183 ns/op
BenchmarkMatchZip-4    	1000000000	         6.23 ns/op
BenchmarkMatchJpeg-4   	2000000000	         4.98 ns/op
BenchmarkMatchGif-4    	2000000000	         4.47 ns/op
BenchmarkMatchPng-4    	1000000000	         6.80 ns/op
PASS
ok  	command-line-arguments	43.604s
```
