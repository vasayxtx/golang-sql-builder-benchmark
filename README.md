golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders.

# Builders

1. dbr: https://github.com/gocraft/dbr
2. goqu: https://github.com/doug-martin/goqu
3. squirrel: https://github.com/lann/squirrel

# Benchmarks

Start MySQL with employees databases in docker:
```
$ ./run_mysql_in_docker.sh
```

Go version:
```
go version                                                                                                                                                                                          [1]
go version go1.14 darwin/amd64
```

Was run on MacBook Pro (Mid 2015, 2.5 GHz Quad-Core Intel Core i7, 16 GB RAM) 

```
go test -bench=. -benchtime=5s .
goos: darwin
goarch: amd64
pkg: github.com/elgris/golang-sql-builder-benchmark
BenchmarkMySQLRaw-8                         2835           2399601 ns/op
BenchmarkMySQLRawWithoutPreparing-8         4732           1301761 ns/op
BenchmarkMySQLDbr-8                         4978           1201836 ns/op
BenchmarkMySQLGoquWithoutPreparing-8        4582           1174317 ns/op
BenchmarkMySQLGoquWithPreparing-8           2476           2172083 ns/op
PASS
```

```
go test -bench=. ./sqlbuild/...
goos: darwin
goarch: amd64
pkg: github.com/elgris/golang-sql-builder-benchmark/sqlbuild

BenchmarkDbrSelectVerySimple-8                    608104              1966 ns/op
BenchmarkDbrSelectVerySimpleRawExp-8             1724547               701 ns/op
BenchmarkDbrSelectSimple-8                        383678              2729 ns/op
BenchmarkDbrSelectSimpleRawExp-8                  830325              1344 ns/op
BenchmarkDbrSelectComplex-8                       237954              4919 ns/op

BenchmarkGoquSelectVerySimple-8                   196347              5980 ns/op
BenchmarkGoquSelectVerySimpleRawExp-8             284560              4078 ns/op
BenchmarkGoquSelectSimple-8                       108627             10884 ns/op
BenchmarkGoquSelectSimpleRawExp-8                 140072              8483 ns/op
BenchmarkGoquSelectComplex-8                       57002             20961 ns/op

BenchmarkSquirrelSelectVerySimple-8               151534              7777 ns/op
BenchmarkSquirrelSelectVerySimpleRawExp-8         257482              4478 ns/op
BenchmarkSquirrelSelectSimple-8                    82788             14166 ns/op
BenchmarkSquirrelSelectSimpleRawExp-8             116158             10175 ns/op
BenchmarkSquirrelSelectComplex-8                   47881             24908 ns/op
PASS
```
