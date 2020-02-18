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

BenchmarkDbrSelectVerySimple-4             329149                                          3537     ns/op  1256   B/op  29   allocs/op
BenchmarkDbrSelectVerySimpleRawExp-4       1000000                                         1875     ns/op  728    B/op  12   allocs/op
BenchmarkDbrSelectSimple-4                 251250                                          4126     ns/op  1776   B/op  40   allocs/op
BenchmarkDbrSelectSimpleRawExp-4           599227                                          2198     ns/op  1208   B/op  22   allocs/op
BenchmarkDbrSelectComplex-4                166544                                          11034    ns/op  2840   B/op  66   allocs/op
BenchmarkDbrRealMySQL-4                    4566                                            297918   ns/op  2058   B/op  41   allocs/op
BenchmarkDbrRealMySQLRawSQL-4              4954                                            257248   ns/op  1738   B/op  32   allocs/op

BenchmarkGoquSelectVerySimple-4            125643                                          10877    ns/op  3148   B/op  96   allocs/op
BenchmarkGoquSelectVerySimpleRawExp-4      181848                                          7178     ns/op  2284   B/op  88   allocs/op
BenchmarkGoquSelectSimple-4                59034                                           18614    ns/op  7064   B/op  167  allocs/op
BenchmarkGoquSelectSimpleRawExp-4          90690                                           13836    ns/op  5992   B/op  156  allocs/op
BenchmarkGoquSelectComplex-4               34989                                           38992    ns/op  13344  B/op  331  allocs/op
BenchmarkGoquRealMySQL-4                   3106                                            334462   ns/op  4379   B/op  94   allocs/op

BenchmarkSquirrelSelectVerySimple-4        77854                                           13255    ns/op  4384   B/op  93   allocs/op
BenchmarkSquirrelSelectVerySimpleRawExp-4  136837                                          12902    ns/op  2512   B/op  49   allocs/op
BenchmarkSquirrelSelectSimple-4            52220                                           24539    ns/op  7369   B/op  162  allocs/op
BenchmarkSquirrelSelectSimpleRawExp-4      66244                                           18420    ns/op  5048   B/op  110  allocs/op
BenchmarkSquirrelSelectComplex-4           22764                                           47824    ns/op  13010  B/op  276  allocs/op
BenchmarkSquirrelRealMySQL-4               2235                                            477741   ns/op  4494   B/op  98   allocs/op

PASS
ok
```
