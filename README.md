golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders. Provides feature list and benchmarks

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. sqrl: https://github.com/elgris/sqrl
4. gocu: https://github.com/doug-martin/goqu - just for SELECT query
5. sqlf: https://github.com/leporo/sqlf


# Feature list

| feature                    | dbr | squirrel | sqrl | goqu | sqlf |
|----------------------------|-----|----------|------|------|------|
| SelectBuilder              | +   | +        | +    | +    | +    |
| InsertBuilder              | +   | +        | +    | +    | +    |
| UpdateBuilder              | +   | +        | +    | +    | +    |
| DeleteBuilder              | +   | +        | +    | +    | +    |
| PostgreSQL support         | +   | +        | +    | +    | +    |
| Custom placeholders        | +   | +        | +    | +    | +    |
| JOINs support              | +   | +        | +    | +    |      |
| Subquery in query builder  |     | +        | +    | +    | +    |
| Aliases for columns        |     | +        | +    | +    |      |
| CASE expression            |     | +        | +    | +    |      |

Some explanations here:
- `Custom placeholders` - ability to use not only `?` placeholders, Useful for PostgreSQL
- `JOINs support` - ability to build JOINs in SELECT queries like `Select("*").From("a").Join("b")`
- `Subquery in query builder` - when you prepare a subquery with one builder and then pass it to another. Something like:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(subQ).From("a")
```
- `Aliases for columns` - easy way to alias a column, especially if column is specified by subquery:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(Alias(subQ, "alias")).From("a")
```
- `CASE expression` - syntactic sugar for [CASE expressions](http://dev.mysql.com/doc/refman/5.7/en/case.html)

# Benchmarks

`go test -bench=. -benchmem | column -t` on 2.6 GHz i5 Macbook Pro:

```
BenchmarkDbrSelectSimple-4                  2000000        782      ns/op  728    B/op  12   allocs/op
BenchmarkDbrSelectConditional-4             1000000        1156     ns/op  976    B/op  17   allocs/op
BenchmarkDbrSelectComplex-4                 500000         3157     ns/op  2360   B/op  38   allocs/op
BenchmarkDbrSelectSubquery-4                500000         2545     ns/op  2160   B/op  29   allocs/op
BenchmarkDbrInsert-4                        500000         2187     ns/op  1136   B/op  26   allocs/op
BenchmarkDbrUpdateSetColumns-4              1000000        2306     ns/op  1297   B/op  28   allocs/op
BenchmarkDbrUpdateSetMap-4                  1000000        2464     ns/op  1296   B/op  28   allocs/op
BenchmarkDbrDelete-4                        2000000        801      ns/op  496    B/op  12   allocs/op

BenchmarkGoquSelectSimple-4                 300000         5323     ns/op  3360   B/op  38   allocs/op
BenchmarkGoquSelectConditional-4            200000         5848     ns/op  3804   B/op  49   allocs/op
BenchmarkGoquSelectComplex-4                100000         18142    ns/op  9464   B/op  169  allocs/op

BenchmarkSqlfSelectSimple-4                 5000000        275      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectSimplePostgreSQL-4       3000000        401      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectConditional-4            3000000        411      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectConditionalPostgreSQL-4  2000000        595      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectComplex-4                2000000        719      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectComplexPostgreSQL-4      2000000        962      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectSubquery-4               2000000        839      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfSelectSubqueryPostgreSQL-4     1000000        1182     ns/op  0      B/op  0    allocs/op
BenchmarkSqlfInsert-4                       2000000        1028     ns/op  0      B/op  0    allocs/op
BenchmarkSqlfInsertPostgreSQL-4             1000000        1114     ns/op  0      B/op  0    allocs/op
BenchmarkSqlfUpdateSetColumns-4             3000000        520      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfUpdateSetColumnsPostgreSQL-4   2000000        674      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfDelete-4                       5000000        286      ns/op  0      B/op  0    allocs/op
BenchmarkSqlfDeletePostgreSQL-4             5000000        346      ns/op  0      B/op  0    allocs/op

BenchmarkSqrlSelectSimple-4                 1000000        1124     ns/op  704    B/op  15   allocs/op
BenchmarkSqrlSelectConditional-4            1000000        1426     ns/op  848    B/op  18   allocs/op
BenchmarkSqrlSelectComplex-4                200000         8461     ns/op  4352   B/op  87   allocs/op
BenchmarkSqrlSelectSubquery-4               300000         6439     ns/op  3352   B/op  65   allocs/op
BenchmarkSqrlSelectMoreComplex-4            200000         11943    ns/op  6961   B/op  138  allocs/op
BenchmarkSqrlInsert-4                       1000000        1706     ns/op  992    B/op  17   allocs/op
BenchmarkSqrlUpdateSetColumns-4             1000000        2133     ns/op  1056   B/op  25   allocs/op
BenchmarkSqrlUpdateSetMap-4                 500000         2568     ns/op  1130   B/op  27   allocs/op
BenchmarkSqrlDelete-4                       3000000        551      ns/op  272    B/op  8    allocs/op

BenchmarkSquirrelSelectSimple-4             200000         5731     ns/op  2512   B/op  49   allocs/op
BenchmarkSquirrelSelectConditional-4        200000         9630     ns/op  3756   B/op  79   allocs/op
BenchmarkSquirrelSelectComplex-4            50000          25307    ns/op  10162  B/op  224  allocs/op
BenchmarkSquirrelSelectSubquery-4           100000         22823    ns/op  8642   B/op  182  allocs/op
BenchmarkSquirrelSelectMoreComplex-4        50000          39758    ns/op  17155  B/op  384  allocs/op
BenchmarkSquirrelInsert-4                   200000         6682     ns/op  3040   B/op  67   allocs/op
BenchmarkSquirrelUpdateSetColumns-4         200000         10834    ns/op  4465   B/op  103  allocs/op
BenchmarkSquirrelUpdateSetMap-4             200000         10996    ns/op  4545   B/op  105  allocs/op
BenchmarkSquirrelDelete-4                   200000         6547     ns/op  2592   B/op  63   allocs/op
```

# Conclusion

If your queries are very simple, pick `sqlf` or `dbr`, the fastest ones.

If really need immutability of query builder and you're ready to sacrifice extra memory, use `squirrel`, the slowest but most reliable one.

If you like those sweet helpers that `squirrel` provides to ease your query building or if you plan to use the same builder for `PostgreSQL`, take `sqrl` as it's balanced between performance and features.

`goqu` has LOTS of features and ways to build queries. Although it requires stubbing sql connection if you need just to build a query. It can be done with [sqlmock](http://github.com/DATA-DOG/go-sqlmock). Disadvantage: the builder is slow and has TOO MANY features, so building a query may become a nightmare. But if you need total control on everything - this is your choice.