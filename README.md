golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders.

# Builders

1. dbr: https://github.com/gocraft/dbr
2. goqu: https://github.com/doug-martin/goqu - just for SELECT query
3. squirrel: https://github.com/lann/squirrel

# Benchmarks

Start MySQL with employees databases in docker:
```
$ ./run_mysql_in_docker.sh
```

Was run on Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz (VirtualBox, 4 CPU).

```
$ go version
go version go1.13.8 linux/amd64

$ go test -bench=. -benchmem | column -t
goos:                                    linux
goarch:                                  amd64
pkg:                                     github.com/elgris/golang-sql-builder-benchmark

BenchmarkDbrSelectVerySimple-4             399624                                          3161     ns/op  1256   B/op  29   allocs/op
BenchmarkDbrSelectVerySimpleRawExp-4       1105027                                         1025     ns/op  728    B/op  12   allocs/op
BenchmarkDbrSelectSimple-4                 277364                                          4213     ns/op  1776   B/op  40   allocs/op
BenchmarkDbrSelectSimpleRawExp-4           314782                                          3899     ns/op  1208   B/op  22   allocs/op
BenchmarkDbrSelectComplex-4                162261                                          7671     ns/op  2840   B/op  66   allocs/op

BenchmarkGoquSelectVerySimple-4            125707                                          10861    ns/op  3148   B/op  96   allocs/op
BenchmarkGoquSelectVerySimpleRawExp-4      177907                                          7129     ns/op  2284   B/op  88   allocs/op
BenchmarkGoquSelectSimple-4                70857                                           18724    ns/op  7064   B/op  167  allocs/op
BenchmarkGoquSelectSimpleRawExp-4          74078                                           14912    ns/op  5992   B/op  156  allocs/op
BenchmarkGoquSelectComplex-4               37394                                           45673    ns/op  13344  B/op  331  allocs/op

BenchmarkSquirrelSelectVerySimple-4        99931                                           12673    ns/op  4384   B/op  93   allocs/op
BenchmarkSquirrelSelectVerySimpleRawExp-4  148560                                          8007     ns/op  2512   B/op  49   allocs/op
BenchmarkSquirrelSelectSimple-4            49158                                           26098    ns/op  7369   B/op  162  allocs/op
BenchmarkSquirrelSelectSimpleRawExp-4      69764                                           18661    ns/op  5048   B/op  110  allocs/op
BenchmarkSquirrelSelectComplex-4           30265                                           41124    ns/op  13010  B/op  276  allocs/op

PASS
ok
```
